import { createAction, handleActions } from 'redux-actions'
import { keys as _keys } from 'lodash'

// @Dev Hack: Mock data
import { mockTeamState } from './__mock__'

const initialState = mockTeamState || {
  entities: {},
  currentTeamId: null
}

export const readTeamsAction = createAction('READ_TEAMS')
export const readTeamsSuccessAction = createAction('READ_TEAMS_SUCCESS')
export const readTeamsFailureAction = createAction('READ_TEAMS_FAILURE')

export const setCurrentTeamAction = createAction('SET_CURRENT_TEAM')
export const setTeamEntitiesAction = createAction('SET_TEAM_ENTITIES')

export const teamReducer = handleActions({

  [`${setCurrentTeamAction}`]: (state, action) => {
    if (!action.payload) {
      return state
    }

    const { currentTeamId } = action.payload

    if (!currentTeamId) {
      return state
    }

    return {
      ...state,
      currentTeamId
    }
  },

  [`${setTeamEntitiesAction}`]: (state, action) => {
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