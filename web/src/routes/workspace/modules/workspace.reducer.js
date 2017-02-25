import { combineReducers } from 'redux'

import { currentTeamReducer as currentTeam } from './current-team/index'

export const workspaceReducer = combineReducers({
  currentTeam
})
