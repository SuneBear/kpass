import { createSelector } from 'reselect'

import { currentTeamSelector } from '../current-team'

export const currentTeamMembersSelector = createSelector(
  (state) => state.member.entities,
  (state) => currentTeamSelector(state),
  (entities, currentTeam) => {
    const memberIds = currentTeam.members

    if (!memberIds) {
      return null
    }

    if (memberIds.length === 0) {
      return []
    }

    return memberIds
      .map((memberId) => {
        return entities[memberId]
      })
  }
)
