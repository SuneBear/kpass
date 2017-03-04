import { combineEpics } from 'redux-observable'
import { Observable } from 'rxjs/Observable'

import {
  mountCurrentTeamAction,
  setCurrentTeamAction
} from './current-team.reducer'

const mountCurrentTeamEpic = (action$) => {
  return action$
    .ofType(`${mountCurrentTeamAction}`)
    .concatMap((action) => {
      const { teamId } = action.payload

      return Observable.of(
        setCurrentTeamAction({
          teamId
        })
      )
    })
}

export const currentTeamEpic = combineEpics(
  mountCurrentTeamEpic
)
