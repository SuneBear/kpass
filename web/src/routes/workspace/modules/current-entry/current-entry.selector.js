import { createSelector } from 'reselect'

export const currentEntrySelector = createSelector(
  (state) => state.entry.entities,
  (state) => state.workspace.currentEntry.entryId,
  (entities, currentEntryId) => {
    if (!currentEntryId) {
      return {}
    }

    return entities[currentEntryId] || {}
  }
)
