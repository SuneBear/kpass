import { createSelector } from 'reselect'

export { routerReducer } from 'react-router-redux'

export const routingSelector = createSelector(
  (state) => state.context.routing.locationBeforeTransitions,
  (routing) => {
    if (!routing) {
      return null
    }

    return routing
  }
)

