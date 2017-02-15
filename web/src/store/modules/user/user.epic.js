import { I18n } from 'react-redux-i18n'
import { combineEpics } from 'redux-observable'
import { normalize } from 'normalizr'
import { Observable } from 'rxjs/Observable'

import { request, sha256 } from 'utils'
import { toast } from 'uis'
import { userSchema } from './user.schema'
import {
  signInUserAction,
  signInUserSuccessAction,
  signInUserFailureAction,
  signInUserAbortAction,
  readUserMeAction,
  readUserMeSuccessAction,
  readUserMeFailureAction,
  setUserMeIdAction,
  setUserEntitiesAction
} from './user.reducer'

const signInUserEpic = (action$) => {
  return action$
    .ofType(`${signInUserAction}`)
    .switchMap((action) => {
      const { username, password } = action.payload

      const body = {
        username,
        password: sha256(password),
        grant_type: 'password'
      }

      return request
        .post('login', body)
        .takeUntil(action$.ofType(
          `${signInUserAbortAction}`
        ))
        .concatMap((response) => {
          return Observable.of(
            signInUserSuccessAction()
          )
        })
        .catch((error) => {
          toast.error({
            message: I18n.t('account.signInFailed')
          })
          return Observable.of(
            signInUserFailureAction(error)
          )
        })
    })
}

const readUserMeEpic = (action$) => {
  return action$
    .ofType(`${readUserMeAction}`)
    .switchMap(() => {
      return request
        .get('user')
        .concatMap((response) => {
          const normalizedResponse = normalize(response, userSchema)

          return Observable.of(
            readUserMeSuccessAction(),
            setUserEntitiesAction({
              entities: normalizedResponse.entities.users
            }),
            setUserMeIdAction({
              userMeId: normalizedResponse.result
            })
          )
        })
        .catch((error) => {
          return Observable.of(
            readUserMeFailureAction(error)
          )
        })
    })
}

export const userEpic = combineEpics(
  signInUserEpic,
  readUserMeEpic
)
