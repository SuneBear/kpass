import { createSelector } from 'reselect'

export const currentTeamSelector = createSelector(
  (state) => state.team.entities,
  (state) => state.team.currentTeamId,
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

export const teamsSelector = createSelector(
  (state) => state.team.entities,
  (entities) => {
    const entitiesKeys = Object.keys(entities)

    return entitiesKeys.map(key => entities[key])
  }
)
