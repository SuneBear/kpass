import { combineReducers } from 'redux'

import { contextReducer as context } from './context'
import { entryReducer as entry } from './entry'
import { formReducer as form } from './form'
import { i18nReducer as i18n } from './i18n'
import { memberReducer as member } from './member'
import { secretReducer as secret } from './secret'
import { teamReducer as team } from './team'
import { userReducer as user } from './user'

export const makeRootReducer = (asyncReducers) => {
  return combineReducers({
    context,
    entry,
    form,
    i18n,
    member,
    secret,
    team,
    user,
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
