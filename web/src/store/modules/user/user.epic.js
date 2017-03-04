import { I18n } from 'react-redux-i18n'
import { combineEpics } from 'redux-observable'
import { push } from 'react-router-redux'
import { normalize } from 'normalizr'
import { Observable } from 'rxjs/Observable'

import { toast } from 'uis'
import { request, sha256, cookie } from 'utils'
import { userSchema } from './user.schema'
import {
  signUpUserAction,
  signUpUserSuccessAction,
  signUpUserFailureAction,
  signUpUserAbortAction,

  signInUserAction,
  signInUserSuccessAction,
  signInUserFailureAction,
  signInUserAbortAction,

  signOutUserAction,
  signOutUserSuccessAction,
  signOutUserFailureAction,

  readUserAction,
  readUserSuccessAction,
  readUserFailureAction,

  updateUserAction,
  updateUserSuccessAction,
  updateUserFailureAction,

  setUserMeIdAction,
  setUserEntitiesAction
} from './user.reducer'

const signUpUserEpic = (action$) => {
  return action$
    .ofType(`${signUpUserAction}`)
    .switchMap((action) => {
      const { username, password, formPromise } = action.payload

      const body = {
        id: username,
        pass: sha256(password)
      }

      return request
        .post('join', body)
        .takeUntil(action$.ofType(
          `${signUpUserAbortAction}`
        ))
        .concatMap((response) => {
          formPromise.resolve()

          return Observable.of(
            signUpUserSuccessAction(),
            signInUserAction({
              username,
              password: sha256(password)
            })
          )
        })
        .catch((errorMessage) => {
          formPromise.reject(errorMessage)

          switch (errorMessage.error.status) {
            case 409:
              toast.error({
                message: I18n.t('account.signUpUserExisted')
              })
              break
            default:
              toast.error({
                message: I18n.t('account.signUpFailed')
              })
          }

          return Observable.of(
            signUpUserFailureAction(errorMessage)
          )
        })
    })
}

const signInUserEpic = (action$) => {
  return action$
    .ofType(`${signInUserAction}`)
    .switchMap((action) => {
      const { username, password, formPromise } = action.payload

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
          formPromise.resolve()

          const token = response.access_token
          request.setToken(token)
          cookie('kp_username', username)

          return Observable.of(
            signInUserSuccessAction(),
            updateUserAction({
              body: response.user
            }),
            setUserMeIdAction({
              userMeId: response.user.id
            }),
            push('/')
          )
        })
        .catch((errorMessage) => {
          formPromise.reject(errorMessage)

          toast.error({
            message: I18n.t('account.signInFailed')
          })

          return Observable.of(
            signInUserFailureAction(errorMessage)
          )
        })
    })
}

const signOutUserEpic = (action$) => {
  return action$
    .ofType(`${signOutUserAction}`)
    .switchMap(() => {
      cookie('kp_username', null)
      return Observable.of(
        signOutUserSuccessAction(),
        setUserMeIdAction({
          userMeId: 'Who\'s Your Daddy?'
        }),
        push('/')
      )
    })
    .catch((errorMessage) => {
      return Observable.of(
        signOutUserFailureAction(errorMessage)
      )
    })
}

const readUserEpic = (action$) => {
  return action$
    .ofType(`${readUserAction}`)
    .switchMap((action) => {
      const { username } = action.payload

      return request
        .get(`user/${username}`)
        .concatMap((response) => {
          return Observable.of(
            readUserSuccessAction(),
            updateUserAction({
              body: response
            })
          )
        })
        .catch((errorMessage) => {
          return Observable.of(
            readUserFailureAction(errorMessage)
          )
        })
    })
}

const updateUserEpic = (action$) => {
  return action$
    .ofType(`${updateUserAction}`)
    .switchMap((action) => {
      const { body } = action.payload

      const normalizedBody = normalize(body, userSchema)

      return Observable.of(
        updateUserSuccessAction(),
        setUserEntitiesAction({
          entities: normalizedBody.entities.users
        })
      )
    })
    .catch((errorMessage) => {
      return Observable.of(
        updateUserFailureAction(errorMessage)
      )
    })
}

export const userEpic = combineEpics(
  signUpUserEpic,
  signInUserEpic,
  signOutUserEpic,
  readUserEpic,
  updateUserEpic
)
