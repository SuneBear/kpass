import { isObject } from './validators'

export const serialize = (obj, prefix) => {
  const str = []

  for (let prop in obj) {
    if (obj.hasOwnProperty(prop)) {
      const key = prefix ? prefix + '[' + prop + ']' : prop
      const value = obj[prop]
      str.push((value !== null && typeof value === 'object')
        ? serialize(value, key)
        : encodeURIComponent(key) + '=' + encodeURIComponent(value))
    }
  }

  return str
    .filter(query => query.length)
    .join('&')
}

export const unserialize = (serializedQuery) => {
  const obj = {}

  serializedQuery
    .replace(/^\?/g, '')
    .split('&')
    .map((item) => {
      const [ key, value ] = decodeURIComponent(item).split('=')
      obj[key] = value
    })

  return obj
}

export const mixinDeep = (target, ...objects) => {
  for (let object of objects) {
    if (!isObject(object)) continue
    for (let key in object) {
      const obj = target[key]
      const val = object[key]
      if (isObject(val) && isObject(obj)) {
        target[key] = mixinDeep(Object.assign({}, obj), Object.assign({}, val))
      } else {
        target[key] = val
      }
    }
  }

  return target
}
