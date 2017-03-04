import { createSelector } from 'reselect'

import { userMeSelector } from 'modules'
import { getUserPermissions } from 'utils'

import { currentTeamSelector } from '../current-team'

export const currentUserPermissionsSelector = createSelector(
  (state) => userMeSelector(state),
  (state) => currentTeamSelector(state),
  (userMe, currentTeam) => {
    const userPermissions = getUserPermissions(userMe, currentTeam)
    // @TODO: Pick required permissions
    return userPermissions
  }
)
