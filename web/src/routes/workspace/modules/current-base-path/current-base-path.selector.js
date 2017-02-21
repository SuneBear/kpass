import { createSelector } from 'reselect'

import { currentTeamSelector } from 'modules'
import { getWorkspaceBashPath } from '../../index'

export const currentBasePathSelector = createSelector(
  (state) => currentTeamSelector(state),
  (currentTeam) => {
    if (!currentTeam) {
      return null
    }

    return getWorkspaceBashPath(currentTeam)
  }
)
