import { createSelector } from 'reselect'

export const currentTeamSelector = createSelector(
  (state) => state.team.entities,
  (state) => state.workspace.currentTeam.teamId,
  (entities, currentTeamId) => {
    if (!currentTeamId) {
      const entitiesKeys = Object.keys(entities)
      currentTeamId = entitiesKeys.filter(
        teamId => entities[teamId].visibility === 'private'
      )
    }

    return entities[currentTeamId]
  }
)
