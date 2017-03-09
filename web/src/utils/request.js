import 'isomorphic-fetch'

import 'rxjs/add/observable/dom/ajax'
import 'rxjs/add/observable/of'
import 'rxjs/add/observable/throw'
import 'rxjs/add/operator/catch'
import 'rxjs/add/operator/delay'
import 'rxjs/add/operator/map'
import 'rxjs/add/operator/mergeMap'
import 'rxjs/add/operator/concatMap'
import 'rxjs/add/operator/switchMap'
import 'rxjs/add/operator/take'
import 'rxjs/add/operator/takeUntil'

import { Observable } from 'rxjs/Observable'
import { Subject } from 'rxjs/Subject'

import config from 'config'
import { NOOP } from './constants'
import { cookie } from './cookie'
import { serialize } from './data'

// @REF: https://github.com/github/fetch/blob/v1.0.0/fetch.js#L313
export const parseHeaders = (rawHeader) => {
  const head = Object.create(null)
  const pairs = rawHeader.trim().split('\n')
  pairs.forEach(function (header) {
    const split = header.trim().split(':')
    const key = split.shift().trim()
    const value = split.join(':').trim()
    head[key] = value
  })
  return head
}

export const httpError$ = new Subject()

export const subscribeHTTPError = (handler = NOOP) => {
  return httpError$.subscribe((res) => {
    handler(res)
  })
}

export class RequestClient {

  static getDefaultOptions () {
    return {
      APIHost: config.API_HOST,
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      credentials: 'include'
    }
  }

  constructor (options) {
    this.options = Object.assign(RequestClient.getDefaultOptions(), options)
    const token = cookie('access_token')
    if (token) this.setToken(token)
  }

  _buildQuery (url, query) {
    if (!query || typeof query !== 'object') {
      return url
    }
    const serializedQuery = serialize(query)
    const combinedUrl = `${url}?${serializedQuery}`
    return combinedUrl
  }

  _makeRequester (method) {
    return (url, body) => {
      const {
        APIHost,
        headers,
        credentials,
        responseType
      } = this.options

      return Observable.ajax({
        url: `${APIHost}/${url}`,
        body,
        method,
        headers: headers,
        responseType: responseType || 'json',
        withCredentials: credentials === 'include'
      })
        .map(value => {
          const resp = value.response
          try {
            return JSON.parse(resp)
          } catch (e) {
            return resp
          }
        })
        .catch((e) => {
          const headers = e.xhr.getAllResponseHeaders()
          const errorMessage = {
            error: new Response(new Blob([JSON.stringify(e.xhr.response)]), {
              status: e.xhr.status,
              statusText: e.xhr.statusText,
              headers: headers.length ? new Headers(parseHeaders(headers)) : new Headers()
            }),
            method,
            url,
            body
          }
          setTimeout(() => {
            httpError$.next(errorMessage)
          }, 10)
          return Observable.throw(errorMessage)
        })
    }
  }

  // Export APIs
  setToken (token) {
    delete this.options.credentials
    this.options.headers.Authorization = `Bearer ${token}`
    cookie('access_token', token)
  }

  get (url, query) {
    const combinedUrl = this._buildQuery(url, query)
    return this._makeRequester('GET')(combinedUrl)
  }

  post (url, body) {
    return this._makeRequester('POST')(url, body)
  }

  put (url, body) {
    return this._makeRequester('PUT')(url, body)
  }

  delete (url, query) {
    const combinedUrl = this._buildQuery(url, query)
    return this._makeRequester('DELETE')(combinedUrl)
  }

}

export const request = new RequestClient()
