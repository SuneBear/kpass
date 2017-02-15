import { AccountLayout } from './layout'
import { SignIn } from './sign-in'

export const ACCOUNT_BASE_PATH = '/account'
export const SIGN_IN_PATH = 'sign-in'

export const redirectToSignIn = (nextState, replace) => {
  return replace(`${ACCOUNT_BASE_PATH}/${SIGN_IN_PATH}`)
}

export default (store) => ({
  path : ACCOUNT_BASE_PATH,
  indexRoute : { onEnter: redirectToSignIn },
  component : AccountLayout,
  childRoutes: [
    {
      path : SIGN_IN_PATH,
      component: SignIn
    }
  ]
})
