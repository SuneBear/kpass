import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'

import { readTeamsAction } from 'modules'
import { WorkspaceLayout as WorkspaceLayoutView } from './workspace-layout.view'

const mapStateToProps = (state) => ({
})

const mapDispatchToProps = (dispatch) => ({
  actions: bindActionCreators({
    readTeams: readTeamsAction
  }, dispatch)
})

const makeContainer = (component) => {
  return connect(
    mapStateToProps,
    mapDispatchToProps
  )(component)
}

export const WorkspaceLayout = makeContainer(WorkspaceLayoutView)
