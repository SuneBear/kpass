import { combineEpics } from 'redux-observable'

import { currentEntryEpic } from './current-entry'
import { currentEntrySecretsEpic } from './current-entry-secrets'
import { currentTeamEpic } from './current-team'
import { currentTeamEntriesEpic } from './current-team-entries'
import { currentTeamMembersEpic } from './current-team-members'

export const workspaceEpic = combineEpics(
  currentEntryEpic,
  currentEntrySecretsEpic,
  currentTeamEpic,
  currentTeamEntriesEpic,
  currentTeamMembersEpic
)
