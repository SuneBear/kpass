import { createSelector } from 'reselect'

import { isPublicTeam } from 'utils'

export const teamsSelector = createSelector(
  (state) => state.team.entities,
  (entities) => {
    const entitiesKeys = Object.keys(entities)

    return entitiesKeys.map(key => entities[key])
  }
)

export const sortedTeamsSelector = createSelector(
  (state) => teamsSelector(state),
  (teams) => {
    if (!teams) {
      return null
    }

    if (teams.length === 0) {
      return []
    }

    return teams
      .sort((prevTeam, nextTeam) => {
        const prevCreatedTime = new Date(prevTeam.created).getTime()
        const nextCreatedTime = new Date(nextTeam.created).getTime()
        if (prevCreatedTime > nextCreatedTime) {
          return -1
        } else {
          return 1
        }
      })
      .sort((prevTeam, nextTeam) => {
        if (!isPublicTeam(prevTeam)) {
          return 1
        }
      })
  }
)
