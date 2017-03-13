import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'

import { userMeSelector } from 'modules'
import {
  currentTeamSelector,
  currentEntrySelector,
  currentEntrySortedSecretsSelector,

  deleteCurrentEntrySecretAction
} from '../../modules'

import { Secrets as SecretsView } from './secrets.view'

const mapStateToProps = (state) => ({
  userMe: userMeSelector(state),
  team: currentTeamSelector(state),
  entry: currentEntrySelector(state),
  secrets: currentEntrySortedSecretsSelector(state)
})

const mapDispatchToProps = (dispatch) => ({
  actions: bindActionCreators({
    deleteSecret: deleteCurrentEntrySecretAction
  }, dispatch)
})

const makeContainer = (component) => {
  return connect(
    mapStateToProps,
    mapDispatchToProps
  )(component)
}

export const Secrets = makeContainer(SecretsView)
