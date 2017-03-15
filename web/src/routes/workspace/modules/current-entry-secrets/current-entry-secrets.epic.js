import { combineEpics } from 'redux-observable'
import { I18n } from 'react-redux-i18n'
import { normalize } from 'normalizr'
import { Observable } from 'rxjs/Observable'

import { request } from 'utils'
import { toast } from 'uis'
import {
  secretSchema,
  setSecretEntitiesAction,
  setEntryEntitiesAction
} from 'modules'
import {
  createCurrentEntrySecretAction,
  createCurrentEntrySecretSuccessAction,
  createCurrentEntrySecretFailureAction,

  updateCurrentEntrySecretAction,
  updateCurrentEntrySecretSuccessAction,
  updateCurrentEntrySecretFailureAction,

  deleteCurrentEntrySecretAction,
  deleteCurrentEntrySecretSuccessAction,
  deleteCurrentEntrySecretFailureAction
} from './current-entry-secrets.action'

const createCurrentEntrySecretEpic = (action$) => {
  return action$
    .ofType(`${createCurrentEntrySecretAction}`)
    .switchMap((action) => {
      const { entry, entryId, body, formPromise } = action.payload

      return request
        .post(`entries/${entryId}/secrets`, body)
        .concatMap((response) => {
          formPromise.resolve()

          toast.success({
            message: I18n.t('secret.createSucceed')
          })

          const normalizedResponse = normalize(response, secretSchema)

          // @Hack: Mock response
          const parsedEntryEntities = {
            [`${entry.id}`]: {
              ...entry,
              secrets: [
                ...entry.secrets,
                normalizedResponse.result
              ]
            }
          }

          return Observable.of(
            setSecretEntitiesAction({
              entities: normalizedResponse.entities.secrets
            }),
            setEntryEntitiesAction({
              entities: parsedEntryEntities
            }),
            createCurrentEntrySecretSuccessAction()
          )
        })
        .catch((error) => {
          formPromise.reject(error)

          return Observable.of(
            createCurrentEntrySecretFailureAction(error)
          )
        })
    })
}

const updateCurrentEntrySecretEpic = (action$) => {
  return action$
    .ofType(`${updateCurrentEntrySecretAction}`)
    .switchMap((action) => {
      const { entryId, secretId, body, formPromise } = action.payload

      return request
        .put(`entries/${entryId}/secrets/${secretId}`, body)
        .concatMap((response) => {
          formPromise.resolve()

          toast.success({
            message: I18n.t('secret.editSucceed')
          })

          const normalizedResponse = normalize(response, secretSchema)

          return Observable.of(
            setSecretEntitiesAction({
              entities: normalizedResponse.entities.secrets
            }),
            updateCurrentEntrySecretSuccessAction()
          )
        })
        .catch((error) => {
          formPromise.reject(error)

          return Observable.of(
            updateCurrentEntrySecretFailureAction(error)
          )
        })
    })
}

const deleteCurrentEntrySecretEpic = (action$) => {
  return action$
    .ofType(`${deleteCurrentEntrySecretAction}`)
    .switchMap((action) => {
      const { entry, entryId, secretId } = action.payload

      return request
        .delete(`entries/${entryId}/secrets/${secretId}`)
        .concatMap(() => {
          toast.success({
            message: I18n.t('secret.deleteSucceed')
          })

          // @Hack: Mock response
          const parsedEntryEntities = {
            [`${entry.id}`]: {
              ...entry,
              secrets: [
                ...entry.secrets.filter(
                  currentSecretId => currentSecretId !== secretId
                )
              ]
            }
          }

          return Observable.of(
            setEntryEntitiesAction({
              entities: parsedEntryEntities
            }),
            deleteCurrentEntrySecretSuccessAction()
          )
        })
        .catch((error) => {
          return Observable.of(
            deleteCurrentEntrySecretFailureAction(error)
          )
        })
    })
}

export const currentEntrySecretsEpic = combineEpics(
  createCurrentEntrySecretEpic,
  updateCurrentEntrySecretEpic,
  deleteCurrentEntrySecretEpic
)
