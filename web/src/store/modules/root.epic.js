import { combineEpics } from 'redux-observable'
import { BehaviorSubject } from 'rxjs/BehaviorSubject'

import { teamEpic } from './team'
import { userEpic } from './user'

export const epic$ = new BehaviorSubject(
  combineEpics(
    teamEpic,
    userEpic
  )
)

export const makeRootEpic = () => {
  return (action$, store) => {
    return epic$.mergeMap(epic => epic(action$, store))
  }
}

export const injectEpic = (epic) => {
  epic$.next(epic)
}
