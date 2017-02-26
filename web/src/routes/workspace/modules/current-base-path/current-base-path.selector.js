import { createSelector } from 'reselect'

import { currentTeamSelector } from '../current-team'
import { getWorkspaceBashPath } from '../../index'

export const currentBasePathSelector = createSelector(
  (state) => currentTeamSelector(state),
  (currentTeam) => {
    if (!currentTeam) {
      return {}
    }

    return getWorkspaceBashPath(currentTeam)
  }
)
