import { combineEpics } from 'redux-observable'

import { userEpic } from './user'

export const makeRootEpic = () => {
  return combineEpics(
    userEpic
  )
}
