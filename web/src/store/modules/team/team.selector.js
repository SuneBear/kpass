import { createSelector } from 'reselect'

export const teamsSelector = createSelector(
  (state) => state.team.entities,
  (entities) => {
    const entitiesKeys = Object.keys(entities)

    return entitiesKeys.map(key => entities[key])
  }
)
