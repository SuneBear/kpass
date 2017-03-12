import { requireAuth, isPublicTeam } from 'utils'
import { injectReducer, injectEpic } from 'modules'
import { WorkspaceLayout } from './layout'
import { workspaceReducer, workspaceEpic } from './modules'
import { Personal } from './personal'
import { Team } from './team'
import { Entries, Members, TeamSettings } from './views'

export const WORKSPACE_BASE_PATH = '/workspace'
export const PERSONAL_PATH = 'personal'
export const TEAM_PATH = 'team'
export const ENTRIES_PATH = 'entries'
export const ENTRIES_FILTER_DEFAULT_PATH = 'unfiltered'
export const ENTRY_PATH = 'entry'
export const MEMBERS_PATH = 'members'
export const SETTINGS_PATH = 'settings'

export const getWorkspaceBashPath = (team) => {
  if (isPublicTeam(team)) {
    return `${WORKSPACE_BASE_PATH}/${TEAM_PATH}/${team.id}`
  } else {
    return `${WORKSPACE_BASE_PATH}/${PERSONAL_PATH}`
  }
}

export const getEntryPathById = (basePath, entryId) => {
  return `${basePath}/${ENTRY_PATH}/${entryId}`
}

export const initWorkspaceLayout = (store) => {
  // Add reducer & epic async
  injectReducer(store, { key: 'workspace', reducer: workspaceReducer })
  injectEpic(workspaceEpic)
  return requireAuth(WorkspaceLayout)
}

export const redirectToPersonal = (_, replace) => {
  return replace(`${WORKSPACE_BASE_PATH}/${PERSONAL_PATH}`)
}

export const redirectToPersonalEntries = (_, replace) => {
  return replace(`${WORKSPACE_BASE_PATH}/${PERSONAL_PATH}/${ENTRIES_PATH}`)
}

export const redirectToTeamEntries = (nextState, replace) => {
  const pathname = nextState.location.pathname
  return replace(`${pathname}/${ENTRIES_PATH}`)
}

export const redirectToEntriesDefaultFilter = (nextState, replace) => {
  const pathname = nextState.location.pathname
  const filterName = nextState.params.filterName
  const entryId = nextState.params.entryId

  if (filterName || entryId) {
    return null
  }

  return replace(`${pathname}/${ENTRIES_FILTER_DEFAULT_PATH}`)
}

export const entryRoutes = (store) => ({
  path : `${ENTRIES_PATH}(/:filterName)`,
  indexRoute : { onEnter: redirectToEntriesDefaultFilter },
  component : Entries,
  childRoutes : [
    {
      path: `${ENTRY_PATH}/:entryId`
    }
  ]
})

export default (store) => ({
  path : WORKSPACE_BASE_PATH,
  indexRoute : { onEnter: redirectToPersonal },
  component : initWorkspaceLayout(store),
  childRoutes : [
    {
      path : PERSONAL_PATH,
      indexRoute : { onEnter: redirectToPersonalEntries },
      component : Personal,
      childRoutes : [
        entryRoutes(store)
      ]
    },

    {
      path : `${TEAM_PATH}/:teamId`,
      component : Team,
      indexRoute : { onEnter: redirectToTeamEntries },
      childRoutes : [
        entryRoutes(store),

        {
          path : MEMBERS_PATH,
          component : Members
        },

        {
          path : SETTINGS_PATH,
          component : TeamSettings
        }
      ]
    }
  ]
})
