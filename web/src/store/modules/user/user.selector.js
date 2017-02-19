import { createSelector } from 'reselect'

export const userMeSelector = createSelector(
  (state) => state.user.entities,
  (state) => state.user.userMeId,
  (entities, userMeId) => {
    if (!userMeId) {
      return null
    }

    return entities[userMeId]
  }
)
