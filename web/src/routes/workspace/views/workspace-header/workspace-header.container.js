import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'

import { userMeSelector, currentTeamSelector } from 'modules'
import { WorkspaceHeader as WorkspaceHeaderView } from './workspace-header.view'

const mapStateToProps = (state) => ({
  userMe: userMeSelector(state),
  currentTeam: currentTeamSelector(state)
})

const mapDispatchToProps = (dispatch) => ({
  actions: bindActionCreators({
  }, dispatch)
})

const makeContainer = (component) => {
  return connect(
    mapStateToProps,
    mapDispatchToProps
  )(component)
}

export const WorkspaceHeader = makeContainer(WorkspaceHeaderView)
