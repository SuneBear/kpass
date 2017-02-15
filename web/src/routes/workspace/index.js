import { requireAuth } from 'utils'
import { WorkspaceLayout } from './layout'
import { Personal } from './personal'

const WORKSPACE_BASE_PATH = '/workspace'
const PERSONAL_PATH = 'personal'

export const redirectToPersonal = (nextState, replace) => {
  return replace(`${WORKSPACE_BASE_PATH}/${PERSONAL_PATH}`)
}

export default (store) => ({
  path : WORKSPACE_BASE_PATH,
  indexRoute : { onEnter: redirectToPersonal },
  component : requireAuth(WorkspaceLayout),
  childRoutes: [
    {
      path : PERSONAL_PATH,
      component: Personal
    }
  ]
})
