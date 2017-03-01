import { createSelector } from 'reselect'

import { userMeSelector } from 'modules'
import { getMemberPermissions } from 'utils'

import { currentTeamSelector } from '../current-team'

export const currentUserPermissionsSelector = createSelector(
  (state) => userMeSelector(state),
  (state) => currentTeamSelector(state),
  (userMe, currentTeam) => {
    const permissions = getMemberPermissions(currentTeam, userMe)
    // @TODO: Pick required permissions
    return permissions
  }
)
