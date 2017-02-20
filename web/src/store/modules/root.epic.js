import { combineEpics } from 'redux-observable'

import { teamEpic } from './team'
import { userEpic } from './user'

export const makeRootEpic = () => {
  return combineEpics(
    teamEpic,
    userEpic
  )
}
