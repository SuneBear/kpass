import { combineReducers } from 'redux'

import i18n from './i18n'
import context from './context'

export const makeRootReducer = (asyncReducers) => {
  return combineReducers({
    context,
    i18n,
    ...asyncReducers
  })
}

export const injectReducer = (store, { key, reducer }) => {
  if (Object.hasOwnProperty.call(store.asyncReducers, key)) return

  store.asyncReducers[key] = reducer
  store.replaceReducer(makeRootReducer(store.asyncReducers))
}

export const injectViewReducer = (store, { key, reducer }) => {
  if (!store.asyncReducers._view) store.asyncReducers._view = {}
  if (Object.hasOwnProperty.call(store.asyncReducers._view, key)) return

  store.asyncReducers._view[key] = reducer
  store.asyncReducers.view = combineReducers(store.asyncReducers._view)
  store.replaceReducer(makeRootReducer(store.asyncReducers))
}

export default makeRootReducer
