import { isAuthedUserMe } from 'utils'
import { Layout } from './layout'
import accountRoutes, { redirectToSignIn } from './account'
import workspaceRoutes, { redirectToPersonal } from './workspace'

export * from './account'
export * from './workspace'

const redirectByAuth = (store) => {
  return isAuthedUserMe(store) ? redirectToPersonal : redirectToSignIn
}

const notFoundRoute = (store) => ({
  path: '*',
  indexRoute : { onEnter: redirectByAuth(store) }
})

export const createRoutes = (store) => ({
  path          : '/',
  indexRoute    : { onEnter: redirectByAuth(store) },
  component     : Layout,
  childRoutes   : [
    accountRoutes(store),
    workspaceRoutes(store),
    notFoundRoute(store)
  ]
})

export default createRoutes
