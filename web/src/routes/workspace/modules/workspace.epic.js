import { combineEpics } from 'redux-observable'

import { currentTeamEpic } from './current-team'
import { currentTeamEntriesEpic } from './current-team-entries'
import { currentTeamMembersEpic } from './current-team-members'

export const workspaceEpic = combineEpics(
  currentTeamEpic,
  currentTeamEntriesEpic,
  currentTeamMembersEpic
)
