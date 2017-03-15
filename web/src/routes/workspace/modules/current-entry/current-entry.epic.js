import { combineEpics } from 'redux-observable'
import { normalize } from 'normalizr'
import { Observable } from 'rxjs/Observable'

import { request } from 'utils'
import {
  entrySchema,
  setEntryEntitiesAction,
  setSecretEntitiesAction
} from 'modules'
import { unmountCurrentTeamAction } from '../index'
import {
  readCurrentEntryAction,
  readCurrentEntrySuccessAction,
  readCurrentEntryFailureAction
} from './current-entry.reducer'

const readCurrentEntryEpic = (action$) => {
  return action$
    .ofType(`${readCurrentEntryAction}`)
    .switchMap((action) => {
      const { entryId } = action.payload

      return request
        .get(`entries/${entryId}`)
        .takeUntil(action$.ofType(
          `${unmountCurrentTeamAction}`
        ))
        .concatMap((response) => {
          const normalizedResponse = normalize(response, entrySchema)

          return Observable.of(
            setSecretEntitiesAction({
              entities: normalizedResponse.entities.secrets
            }),
            setEntryEntitiesAction({
              entities: normalizedResponse.entities.entries
            }),
            readCurrentEntrySuccessAction()
          )
        })
        .catch((error) => {
          return Observable.of(
            readCurrentEntryFailureAction(error)
          )
        })
    })
}

export const currentEntryEpic = combineEpics(
  readCurrentEntryEpic
)
