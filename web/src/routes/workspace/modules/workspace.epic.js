import { combineEpics } from 'redux-observable'

import { currentTeamEpic } from './current-team'
import { currentTeamEntriesEpic } from './current-team-entries'

export const workspaceEpic = combineEpics(
  currentTeamEpic,
  currentTeamEntriesEpic
)
