import { combineReducers } from 'redux'

import { currentEntryReducer as currentEntry } from './current-entry'
import { currentTeamReducer as currentTeam } from './current-team'
import { currentTeamEntriesReducer as currentTeamEntries } from './current-team-entries'

export const workspaceReducer = combineReducers({
  currentEntry,
  currentTeam,
  currentTeamEntries
})
