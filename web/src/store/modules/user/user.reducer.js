import { createAction, handleActions } from 'redux-actions'
import { keys as _keys } from 'lodash'

const initialState = {
  entities: {},
  userMeId: null
}

export const signInUserAction = createAction('SIGN_IN_USER')
export const signInUserSuccessAction = createAction('SIGN_IN_USER_SUCCESS')
export const signInUserFailureAction = createAction('SIGN_IN_USER_FAILURE')
export const signInUserAbortAction = createAction('SIGN_IN_USER_ABORT')
export const readUserMeAction = createAction('READ_USER_ME')
export const readUserMeSuccessAction = createAction('READ_USER_ME_SUCCESS')
export const readUserMeFailureAction = createAction('READ_USER_ME_FAILURE')
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
