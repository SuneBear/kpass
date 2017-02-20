import { I18n } from 'react-redux-i18n'
import { combineEpics } from 'redux-observable'
import { normalize } from 'normalizr'
import { Observable } from 'rxjs/Observable'

import { request, sha256 } from 'utils'
import { toast } from 'uis'
import { readTeamsAction } from '../team'
import { userSchema } from './user.schema'
import {
  signUpUserAction,
  signUpUserSuccessAction,
  signUpUserFailureAction,
  signUpUserAbortAction,

  signInUserAction,
  signInUserSuccessAction,
  signInUserFailureAction,
  signInUserAbortAction,

  readUserMeAction,
  readUserMeSuccessAction,
  readUserMeFailureAction,

  setUserMeIdAction,
  setUserEntitiesAction
} from './user.reducer'

const signUpUserEpic = (action$) => {
  return action$
    .ofType(`${signUpUserAction}`)
    .switchMap((action) => {
      const { username, password } = action.payload

      const body = {
        id: username,
        pass: sha256(password)
      }

      return request
        .post('join', body)
        .takeUntil(action$.ofType(
          `${signUpUserAbortAction}`
        ))
        .concatMap((response) => {
          const token = response.access_token
          request.setToken(token)
          return Observable.of(
            signInUserSuccessAction(),
            readTeamsAction()
          )
        })
        .catch((error) => {
          toast.error({
            message: I18n.t('account.signUpFailed')
          })
          return Observable.of(
            signInUserFailureAction(error)
          )
        })
    })
}

const signInUserEpic = (action$) => {
  return action$
    .ofType(`${signInUserAction}`)
    .switchMap((action) => {
      const { username, password } = action.payload

      const body = {
        username,
        password: sha256(password),
        grant_type: 'password'
      }

      return request
        .post('login', body)
        .takeUntil(action$.ofType(
          `${signInUserAbortAction}`
        ))
        .concatMap((response) => {
          const token = response.access_token
          request.setToken(token)
          return Observable.of(
            signInUserSuccessAction(),
            readTeamsAction()
          )
        })
        .catch((error) => {
          toast.error({
            message: I18n.t('account.signInFailed')
          })
          return Observable.of(
            signInUserFailureAction(error)
          )
        })
    })
}

const readUserMeEpic = (action$) => {
  return action$
    .ofType(`${readUserMeAction}`)
    .switchMap(() => {
      return request
        .get('user')
        .concatMap((response) => {
          const normalizedResponse = normalize(response, userSchema)

          return Observable.of(
            readUserMeSuccessAction(),
            setUserEntitiesAction({
              entities: normalizedResponse.entities.users
            }),
            setUserMeIdAction({
              userMeId: normalizedResponse.result
            })
          )
        })
        .catch((error) => {
          return Observable.of(
            readUserMeFailureAction(error)
          )
        })
    })
}

export const userEpic = combineEpics(
  signUpUserEpic,
  signInUserEpic,
  readUserMeEpic
)
