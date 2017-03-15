import { createAction, handleActions } from 'redux-actions'

const initialState = {
  entryId: null
}

export const readCurrentEntryAction = createAction('READ_CURRENT_ENTRY')
export const readCurrentEntrySuccessAction = createAction('READ_CURRENT_ENTRY_SUCCESS')
export const readCurrentEntryFailureAction = createAction('READ_CURRENT_ENTRY_FAILURE')

export const setCurrentEntryAction = createAction('SET_CURRENT_ENTRY')

export const currentEntryReducer = handleActions({

  [`${setCurrentEntryAction}`]: (state, action) => {
    if (!action.payload) {
      return state
    }

    const { entryId } = action.payload

    return {
      ...state,
      entryId
    }
  }

}, initialState)
