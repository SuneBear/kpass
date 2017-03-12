import { createAction, handleActions } from 'redux-actions'
import { keys as _keys } from 'lodash'

// @Dev Hack: Mock data
// import { mockTeamState } from './__mock__'

const initialState = {
  entities: {}
}

export const createTeamAction = createAction('CREATE_TEAM')
export const createTeamSuccessAction = createAction('CREATE_TEAM_SUCCESS')
export const createTeamFailureAction = createAction('CREATE_TEAM_FAILURE')

export const joinTeamAction = createAction('JOIN_TEAM')
export const joinTeamSuccessAction = createAction('JOIN_TEAM_SUCCESS')
export const joinTeamFailureAction = createAction('JOIN_TEAM_FAILURE')

export const readTeamsAction = createAction('READ_TEAMS')
export const readTeamsSuccessAction = createAction('READ_TEAMS_SUCCESS')
export const readTeamsFailureAction = createAction('READ_TEAMS_FAILURE')

export const deleteTeamAction = createAction('DELETE_TEAM')
export const deleteTeamSuccessAction = createAction('DELETE_TEAM_SUCCESS')
export const deleteTeamFailureAction = createAction('DELETE_TEAM_FAILURE')

export const setTeamEntitiesAction = createAction('SET_TEAM_ENTITIES')

export const teamReducer = handleActions({

  [`${deleteTeamSuccessAction}`]: (state, action) => {
    if (!action.payload) {
      return state
    }

    const { teamId } = action.payload

    if (!teamId) {
      return state
    }

    const nextEntitiesKeys = _keys(state.entities)
      .filter((id) => id !== teamId)

    const nextStateEntities = nextEntitiesKeys
      .reduce((stateEntities, entityKey) => {
        stateEntities[entityKey] = state.entities[entityKey]
        return stateEntities
      }, {})

    return {
      ...state,
      entities: nextStateEntities
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
