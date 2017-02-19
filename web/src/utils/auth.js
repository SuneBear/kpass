import { UserAuthWrapper } from 'redux-auth-wrapper'
import { push } from 'react-router-redux'

import { userMeSelector } from 'modules'
import { ACCOUNT_BASE_PATH } from 'routes/account'

export const isAuthedUser = (user) => {
  return user && !user.isBlocked
}

export const isAuthedUserMe = (store) => {
  const state = store.getState()
  const userMe = userMeSelector(state)
  return isAuthedUser(userMe)
}

export const requireAuth = UserAuthWrapper({
  authSelector: state => userMeSelector(state),
  predicate: isAuthedUser,
  redirectAction: push,
  failureRedirectPath: ACCOUNT_BASE_PATH,
  wrapperDisplayName: 'UserIsAuthenticated'
})
