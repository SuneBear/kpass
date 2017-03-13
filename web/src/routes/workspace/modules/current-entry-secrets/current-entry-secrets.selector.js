import { createSelector } from 'reselect'

import { currentEntrySelector } from '../current-entry'

export const currentEntrySecretsSelector = createSelector(
  (state) => state.secret.entities,
  (state) => currentEntrySelector(state),
  (entities, currentEntry) => {
    const secretIds = currentEntry.secrets

    if (!secretIds) {
      return null
    }

    if (secretIds.length === 0) {
      return []
    }

    return secretIds
      .map((secretId) => {
        return entities[secretId]
      })
  }
)

export const currentEntrySortedSecretsSelector = createSelector(
  (state) => currentEntrySecretsSelector(state),
  (secrets) => {
    if (!secrets) {
      return null
    }

    if (secrets.length === 0) {
      return []
    }

    return secrets
      .sort((prevSecret, nextSecret) => {
        const prevCreatedTime = new Date(prevSecret.created).getTime()
        const nextCreatedTime = new Date(nextSecret.created).getTime()
        if (prevCreatedTime > nextCreatedTime) {
          return 1
        } else {
          return -1
        }
      })
  }
)
