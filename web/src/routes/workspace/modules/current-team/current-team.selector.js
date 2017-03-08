import { createSelector } from 'reselect'

import { isPublicTeam } from 'utils'

export const privateTeamIdSelector = createSelector(
  (state) => state.team.entities,
  (entities) => {
    const entitiesKeys = Object.keys(entities)
    const privateTeamId = entitiesKeys.find(
      teamId => !isPublicTeam(entities[teamId])
    )

    return privateTeamId
  }
)

export const currentTeamSelector = createSelector(
  (state) => state.team.entities,
  (state) => state.workspace.currentTeam.teamId,
  (state) => privateTeamIdSelector(state),
  (entities, currentTeamId, privateTeamId) => {
    if (!currentTeamId) {
      currentTeamId = privateTeamId
    }

    return entities[currentTeamId] || {}
  }
)
