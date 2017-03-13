import { createAction, handleActions } from 'redux-actions'
import { keys as _keys } from 'lodash'

const initialState = {
  entities: {}
}

export const setSecretEntitiesAction = createAction('SET_SECRET_ENTITIES')

export const secretReducer = handleActions({

  [`${setSecretEntitiesAction}`]: (state, action) => {
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
