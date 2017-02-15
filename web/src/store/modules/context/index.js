import { combineReducers } from 'redux'

import { routerReducer as routing } from './routing'
import { titleReducer as title } from './title'

export const contextReducer = combineReducers({
  routing,
  title
})
