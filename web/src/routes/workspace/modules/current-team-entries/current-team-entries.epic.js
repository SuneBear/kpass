import { combineEpics } from 'redux-observable'
import { normalize } from 'normalizr'
import { Observable } from 'rxjs/Observable'

import { request } from 'utils'
import { entrySchema, entriesSchema, setEntryEntitiesAction } from 'modules'
import { unmountCurrentTeamAction } from '../index'
import {
  createCurrentTeamEntryAction,
  createCurrentTeamEntrySuccessAction,
  createCurrentTeamEntryFailureAction,

  readCurrentTeamEntriesAction,
  readCurrentTeamEntriesSuccessAction,
  readCurrentTeamEntriesFailureAction,

  deleteCurrentTeamEntryAction,
  deleteCurrentTeamEntrySuccessAction,
  deleteCurrentTeamEntryFailureAction
} from './current-team-entries.reducer'

const createCurrentTeamEntryEpic = (action$) => {
  return action$
    .ofType(`${createCurrentTeamEntryAction}`)
    .switchMap((action) => {
      const { teamId, body, formPromise } = action.payload

      return request
        .post(`teams/${teamId}/entries`, body)
        .takeUntil(action$.ofType(
          `${unmountCurrentTeamAction}`
        ))
        .concatMap((response) => {
          formPromise.resolve()

          const normalizedResponse = normalize(response, entrySchema)

          return Observable.of(
            setEntryEntitiesAction({
              entities: normalizedResponse.entities.entries
            }),
            createCurrentTeamEntrySuccessAction({
              entryId: normalizedResponse.result
            })
          )
        })
        .catch((error) => {
          formPromise.reject(error)

          return Observable.of(
            createCurrentTeamEntryFailureAction(error)
          )
        })
    })
}

const readCurrentTeamEntriesEpic = (action$) => {
  return action$
    .ofType(`${readCurrentTeamEntriesAction}`)
    .switchMap((action) => {
      const { teamId } = action.payload

      return request
        .get(`teams/${teamId}/entries`)
        .takeUntil(action$.ofType(
          `${unmountCurrentTeamAction}`
        ))
        .concatMap((response) => {
          const normalizedResponse = normalize(response, entriesSchema)

          return Observable.of(
            setEntryEntitiesAction({
              entities: normalizedResponse.entities.entries
            }),
            readCurrentTeamEntriesSuccessAction({
              entryIds: normalizedResponse.result
            })
          )
        })
        .catch((error) => {
          return Observable.of(
            readCurrentTeamEntriesFailureAction(error)
          )
        })
    })
}

const deleteCurrentTeamEntryEpic = (action$) => {
  return action$
    .ofType(`${deleteCurrentTeamEntryAction}`)
    .switchMap((action) => {
      const { entryId } = action.payload

      return request
        .delete(`entries/${entryId}`)
        .takeUntil(action$.ofType(
          `${unmountCurrentTeamAction}`
        ))
        .concatMap(() => {
          return Observable.of(
            deleteCurrentTeamEntrySuccessAction({
              entryId
            })
          )
        })
        .catch((error) => {
          return Observable.of(
            deleteCurrentTeamEntryFailureAction(error)
          )
        })
    })
}

export const currentTeamEntriesEpic = combineEpics(
  createCurrentTeamEntryEpic,
  readCurrentTeamEntriesEpic,
  deleteCurrentTeamEntryEpic
)
