import { createAction, handleActions } from 'redux-actions'

const initialState = {
  entryIds: null
}

export const createCurrentTeamEntryAction = createAction('CREATE_CURRENT_TEAM_ENTRY')
export const createCurrentTeamEntrySuccessAction = createAction('CREATE_CURRENT_TEAM_ENTRY_SUCCESS')
export const createCurrentTeamEntryFailureAction = createAction('CREATE_CURRENT_TEAM_ENTRY_FAILURE')

export const readCurrentTeamEntriesAction = createAction('READ_CURRENT_TEAM_ENTRIES')
export const readCurrentTeamEntriesSuccessAction = createAction('READ_CURRENT_TEAM_ENTRIES_SUCCESS')
export const readCurrentTeamEntriesFailureAction = createAction('READ_CURRENT_TEAM_ENTRIES_FAILURE')

export const deleteCurrentTeamEntryAction = createAction('DELETE_CURRENT_TEAM_ENTRY')
export const deleteCurrentTeamEntrySuccessAction = createAction('DELETE_CURRENT_TEAM_ENTRY_SUCCESS')
export const deleteCurrentTeamEntryFailureAction = createAction('DELETE_CURRENT_TEAM_ENTRY_FAILURE')

export const currentTeamEntriesReducer = handleActions({

  [`${createCurrentTeamEntrySuccessAction}`]: (state, action) => {
    if (!action.payload) {
      return state
    }

    const { entryId } = action.payload

    if (!entryId) {
      return state
    }

    return {
      ...state,
      entryIds: [ ...state.entryIds, entryId ]
    }
  },

  [`${readCurrentTeamEntriesAction}`]: (state) => {
    return {
      ...state,
      entryIds: null
    }
  },

  [`${readCurrentTeamEntriesSuccessAction}`]: (state, action) => {
    if (!action.payload) {
      return state
    }

    const { entryIds } = action.payload

    if (!entryIds) {
      return state
    }

    return {
      ...state,
      entryIds
    }
  },

  [`${readCurrentTeamEntriesFailureAction}`]: (state) => {
    return {
      ...state,
      entryIds: null
    }
  },

  [`${deleteCurrentTeamEntrySuccessAction}`]: (state, action) => {
    if (!action.payload) {
      return state
    }

    const { entryId } = action.payload

    if (!entryId) {
      return state
    }

    const entryIdIndex = state.indexOf(entryId)

    if (!entryIdIndex) {
      return state
    }

    return {
      ...state,
      entryIds: [
        ...state.entryIds.slice(0, entryIdIndex),
        ...state.entryIds.slice(entryIdIndex + 1)
      ]
    }
  }

}, initialState)
