import { combineEpics } from 'redux-observable'
import { normalize } from 'normalizr'
import { Observable } from 'rxjs/Observable'

import { request } from 'utils'
import { teamSchema } from './team.schema'
import {
  readTeamsAction,
  readTeamsSuccessAction,
  readTeamsFailureAction,

  setTeamEntitiesAction
} from './team.reducer'

const readTeamsEpic = (action$) => {
  return action$
    .ofType(`${readTeamsAction}`)
    .switchMap((action) => {
      return request
        .get('teams')
        .concatMap((response) => {
          const normalizedResponse = normalize(response, teamSchema)

          return Observable.of(
            readTeamsSuccessAction(),
            setTeamEntitiesAction({
              entities: normalizedResponse.entities.teams
            })
          )
        })
        .catch((error) => {
          return Observable.of(
            readTeamsFailureAction(error)
          )
        })
    })
}

export const teamEpic = combineEpics(
  readTeamsEpic
)
