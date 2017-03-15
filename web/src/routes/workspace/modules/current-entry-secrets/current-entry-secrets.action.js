import { createAction } from 'redux-actions'

export const createCurrentEntrySecretAction = createAction('CREATE_CURRENT_ENTRY_SECRET')
export const createCurrentEntrySecretSuccessAction = createAction('CREATE_CURRENT_ENTRY_SECRET_SUCCESS')
export const createCurrentEntrySecretFailureAction = createAction('CREATE_CURRENT_ENTRY_SECRET_FAILURE')

export const updateCurrentEntrySecretAction = createAction('UPDATE_CURRENT_ENTRY_SECRET')
export const updateCurrentEntrySecretSuccessAction = createAction('UPDATE_CURRENT_ENTRY_SECRET_SUCCESS')
export const updateCurrentEntrySecretFailureAction = createAction('UPDATE_CURRENT_ENTRY_SECRET_FAILURE')

export const deleteCurrentEntrySecretAction = createAction('DELETE_CURRENT_ENTRY_SECRET')
export const deleteCurrentEntrySecretSuccessAction = createAction('DELETE_CURRENT_ENTRY_SECRET_SUCCESS')
export const deleteCurrentEntrySecretFailureAction = createAction('DELETE_CURRENT_ENTRY_SECRET_FAILURE')
