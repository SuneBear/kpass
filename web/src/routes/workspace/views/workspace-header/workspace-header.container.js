import { connect } from 'react-redux'
import { push } from 'react-router-redux'
import { bindActionCreators } from 'redux'

import { userMeSelector, sortedTeamsSelector, signOutUserAction } from 'modules'
import { currentTeamSelector } from '../../modules'
import { WorkspaceHeader as WorkspaceHeaderView } from './workspace-header.view'

const mapStateToProps = (state) => ({
  userMe: userMeSelector(state),
  teams: sortedTeamsSelector(state),
  currentTeam: currentTeamSelector(state)
})

const mapDispatchToProps = (dispatch) => ({
  actions: bindActionCreators({
    push: push,
    signOutUser: signOutUserAction
  }, dispatch)
})

const makeContainer = (component) => {
  return connect(
    mapStateToProps,
    mapDispatchToProps
  )(component)
}

export const WorkspaceHeader = makeContainer(WorkspaceHeaderView)
