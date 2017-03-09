import { createAction } from 'redux-actions'

export const createCurrentTeamMemberAction = createAction('CREATE_CURRENT_TEAM_MEMBER')
export const createCurrentTeamMemberSuccessAction = createAction('CREATE_CURRENT_TEAM_MEMBER_SUCCESS')
export const createCurrentTeamMemberFailureAction = createAction('CREATE_CURRENT_TEAM_MEMBER_FAILURE')

export const deleteCurrentTeamMemberAction = createAction('DELETE_CURRENT_TEAM_MEMBER')
export const deleteCurrentTeamMemberSuccessAction = createAction('DELETE_CURRENT_TEAM_MEMBER_SUCCESS')
export const deleteCurrentTeamMemberFailureAction = createAction('DELETE_CURRENT_TEAM_MEMBER_FAILURE')
