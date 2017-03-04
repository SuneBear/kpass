import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'

import { userMeSelector, updateUserAction } from 'modules'
import { AccountSettings as AccountSettingsView } from './account-settings.view'

const mapStateToProps = (state) => ({
  userMe: userMeSelector(state)
})

const mapDispatchToProps = (dispatch) => ({
  actions: bindActionCreators({
    updateUser: updateUserAction
  }, dispatch)
})

const makeContainer = (component) => {
  return connect(
    mapStateToProps,
    mapDispatchToProps
  )(component)
}

export const AccountSettings = makeContainer(AccountSettingsView)
