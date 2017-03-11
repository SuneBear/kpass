import { createSelector } from 'reselect'

import { currentTeamSelector } from '../current-team'
import { currentTeamEntriesFilterSelector } from '../current-team-entries'
import {
  getWorkspaceBashPath,
  ENTRIES_PATH
} from '../../index'

export const currentBasePathSelector = createSelector(
  (state) => currentTeamSelector(state),
  (currentTeam) => {
    if (!currentTeam) {
      return '/'
    }

    return getWorkspaceBashPath(currentTeam)
  }
)

export const currentEntriesBasePathSelector = createSelector(
  (state) => currentBasePathSelector(state),
  (state) => currentTeamEntriesFilterSelector(state),
  (basePath, filter) => {
    return `${basePath}/${ENTRIES_PATH}/${filter}`
  }
)
