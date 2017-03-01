import { combineReducers } from 'redux'

import { currentEntryReducer as currentEntry } from './current-entry/index'
import { currentTeamReducer as currentTeam } from './current-team/index'

export const workspaceReducer = combineReducers({
  currentEntry,
  currentTeam
})
