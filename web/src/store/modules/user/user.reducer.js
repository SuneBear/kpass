import { createAction, handleActions } from 'redux-actions'
import { keys as _keys } from 'lodash'

// @Dev Hack: Mock data
// import { mockUserState } from './__mock__'

const initialState = {
  entities: {},
  userMeId: null
}

export const signUpUserAction = createAction('SIGN_UP_USER')
export const signUpUserSuccessAction = createAction('SIGN_UP_USER_SUCCESS')
export const signUpUserFailureAction = createAction('SIGN_UP_USER_FAILURE')
export const signUpUserAbortAction = createAction('SIGN_UP_USER_ABORT')

export const signInUserAction = createAction('SIGN_IN_USER')
export const signInUserSuccessAction = createAction('SIGN_IN_USER_SUCCESS')
export const signInUserFailureAction = createAction('SIGN_IN_USER_FAILURE')
export const signInUserAbortAction = createAction('SIGN_IN_USER_ABORT')

export const signOutUserAction = createAction('SIGN_OUT_USER')
export const signOutUserSuccessAction = createAction('SIGN_OUT_USER_SUCCESS')
export const signOutUserFailureAction = createAction('SIGN_OUT_USER_FAILURE')

export const readUserAction = createAction('READ_USER')
export const readUserSuccessAction = createAction('READ_USER_SUCCESS')
export const readUserFailureAction = createAction('READ_USER_FAILURE')

export const updateUserAction = createAction('UPDATE_USER')
export const updateUserSuccessAction = createAction('UPDATE_USER_SUCCESS')
export const updateUserFailureAction = createAction('UPDATE_USER_FAILURE')

export const setUserMeIdAction = createAction('SET_USER_ME_ID')
export const setUserEntitiesAction = createAction('SET_USER_ENTITIES')

export const userReducer = handleActions({

  [`${setUserMeIdAction}`]: (state, action) => {
    if (!action.payload) {
      return state
    }

    const { userMeId } = action.payload

    if (!userMeId) {
      return state
    }

    return {
      ...state,
      userMeId
    }
  },

  [`${setUserEntitiesAction}`]: (state, action) => {
    if (!action.payload) {
      return state
    }

    const { entities } = action.payload

    if (!entities) {
      return state
    }

    const entitiesKeys = _keys(entities)

    const nextStateEntities = entitiesKeys
      .reduce((stateEntities, entityKey) => {
        return {
          ...stateEntities,
          [entityKey]: {
            ...stateEntities[entityKey],
            ...entities[entityKey]
          }
        }
      }, state.entities)

    return {
      ...state,
      entities: nextStateEntities
    }
  }

}, initialState)
