import { createSelector } from 'reselect'

import { currentTeamSelector } from 'modules'
import {
  WORKSPACE_BASE_PATH,
  PERSONAL_PATH,
  TEAM_PATH
} from '../../index'

export const currentBasePathSelector = createSelector(
  (state) => currentTeamSelector(state),
  (currentTeam) => {
    if (!currentTeam) {
      return null
    }

    if (currentTeam.visibility === 'private') {
      return `${WORKSPACE_BASE_PATH}/${PERSONAL_PATH}`
    } else {
      return `${WORKSPACE_BASE_PATH}/${TEAM_PATH}/${currentTeam.id}`
    }
  }
)
