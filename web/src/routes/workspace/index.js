import { requireAuth } from 'utils'
import { WorkspaceLayout } from './layout'
import { Personal } from './personal'
import { Team } from './team'
import { Entries } from './views'

export const WORKSPACE_BASE_PATH = '/workspace'
export const PERSONAL_PATH = 'personal'
export const TEAM_PATH = 'team'
export const ENTRIES_PATH = 'entries'

export const redirectToPersonal = (nextState, replace) => {
  return replace(`${WORKSPACE_BASE_PATH}/${PERSONAL_PATH}`)
}

export const redirectToPersonalEntries = (nextState, replace) => {
  return replace(`${WORKSPACE_BASE_PATH}/${PERSONAL_PATH}/${ENTRIES_PATH}`)
}

export default (store) => ({
  path : WORKSPACE_BASE_PATH,
  indexRoute : { onEnter: redirectToPersonal },
  component : requireAuth(WorkspaceLayout),
  childRoutes: [
    {
      path : PERSONAL_PATH,
      indexRoute : { onEnter: redirectToPersonalEntries },
      component: Personal,
      childRoutes: [
        {
          path : ENTRIES_PATH,
          component : Entries
        }
      ]
    },

    {
      path : `${TEAM_PATH}/:teamId`,
      component: Team,
      childRoutes: [
        {
          path : ENTRIES_PATH,
          component : Entries
        }
      ]
    }
  ]
})
