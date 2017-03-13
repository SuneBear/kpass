import { createSelector } from 'reselect'

import { ENTRIES_FILTER_DEFAULT_PATH } from '../../index'

export const currentTeamEntriesSelector = createSelector(
  (state) => state.entry.entities,
  (state) => state.workspace.currentTeamEntries.entryIds,
  (entities, entryIds) => {
    if (!entities || !entryIds) {
      return null
    }

    if (entryIds.length === 0) {
      return []
    }

    return entryIds
      .map((entryId) => {
        return entities[entryId]
      })
  }
)

export const currentTeamSortedEntriesSelector = createSelector(
  (state) => currentTeamEntriesSelector(state),
  (entries) => {
    if (!entries) {
      return null
    }

    if (entries.length === 0) {
      return []
    }

    return entries
      .sort((prevEntry, nextEntry) => {
        if (new Date(prevEntry.updated).getTime() > new Date(nextEntry.updated).getTime()) {
          return -1
        } else {
          return 1
        }
      })
  }
)

export const currentTeamEntriesFilterSelector = createSelector(
  (state) => state.workspace.currentTeamEntries.filter,
  (filter) => {
    return filter || ENTRIES_FILTER_DEFAULT_PATH
  }
)
