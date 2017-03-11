import { I18n } from 'react-redux-i18n'
import { push } from 'react-router-redux'
import { combineEpics } from 'redux-observable'
import { normalize } from 'normalizr'
import { Observable } from 'rxjs/Observable'

import { request } from 'utils'
import { toast } from 'uis'
import { getWorkspaceBashPath } from 'routes'
import { setMemberEntitiesAction } from '../member'
import { teamSchema, teamsSchema } from './team.schema'
import {
  createTeamAction,
  createTeamSuccessAction,
  createTeamFailureAction,

  readTeamsAction,
  readTeamsSuccessAction,
  readTeamsFailureAction,

  setTeamEntitiesAction
} from './team.reducer'

const createTeamEpic = (action$) => {
  return action$
    .ofType(`${createTeamAction}`)
    .switchMap((action) => {
      const { body, formPromise } = action.payload

      return request
        .post('teams', body)
        .concatMap((response) => {
          formPromise.resolve()

          toast.success({
            message: I18n.t('team.createSucceed')
          })

          const normalizedResponse = normalize(response, teamSchema)

          return Observable.of(
            createTeamSuccessAction(),
            setMemberEntitiesAction({
              entities: normalizedResponse.entities.members
            }),
            setTeamEntitiesAction({
              entities: normalizedResponse.entities.teams
            }),
            push(getWorkspaceBashPath({
              id: normalizedResponse.result
            }))
          )
        })
        .catch((error) => {
          formPromise.reject(error)

          return Observable.of(
            createTeamFailureAction(error)
          )
        })
    })
}

const readTeamsEpic = (action$) => {
  return action$
    .ofType(`${readTeamsAction}`)
    .switchMap((action) => {
      return request
        .get('teams')
        .concatMap((response) => {
          const normalizedResponse = normalize(response, teamsSchema)

          return Observable.of(
            readTeamsSuccessAction(),
            setMemberEntitiesAction({
              entities: normalizedResponse.entities.members
            }),
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
  createTeamEpic,
  readTeamsEpic
)
