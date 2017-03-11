import { combineEpics } from 'redux-observable'
import { normalize } from 'normalizr'
import { Observable } from 'rxjs/Observable'
import { I18n } from 'react-redux-i18n'

import { request } from 'utils'
import { toast } from 'uis'
import { teamSchema, setTeamEntitiesAction } from 'modules'
import { readCurrentTeamEntriesAction } from '../../modules'
import {
  updateCurrentTeamAction,
  updateCurrentTeamSuccessAction,
  updateCurrentTeamFailureAction,

  mountCurrentTeamAction,
  unmountCurrentTeamAction,
  setCurrentTeamAction
} from './current-team.reducer'

const updateCurrentTeamEpic = (action$) => {
  return action$
    .ofType(`${updateCurrentTeamAction}`)
    .switchMap((action) => {
      const { teamId, body, formPromise } = action.payload

      return request
        .put(`teams/${teamId}`, body)
        .takeUntil(action$.ofType(
          `${unmountCurrentTeamAction}`
        ))
        .concatMap((response) => {
          formPromise && formPromise.resolve()

          toast.success({
            message: I18n.t('teamSettings.updateSucceed')
          })

          const normalizedResponse = normalize(response, teamSchema)

          return Observable.of(
            setTeamEntitiesAction({
              entities: normalizedResponse.entities.teams
            }),
            updateCurrentTeamSuccessAction()
          )
        })
        .catch((error) => {
          formPromise && formPromise.reject(error)

          return Observable.of(
            updateCurrentTeamFailureAction(error)
          )
        })
    })
}

const mountCurrentTeamEpic = (action$) => {
  return action$
    .ofType(`${mountCurrentTeamAction}`)
    .concatMap((action) => {
      const { teamId } = action.payload

      return Observable.of(
        setCurrentTeamAction({
          teamId
        }),

        readCurrentTeamEntriesAction({
          teamId
        })
      )
    })
}

export const currentTeamEpic = combineEpics(
  updateCurrentTeamEpic,
  mountCurrentTeamEpic
)
