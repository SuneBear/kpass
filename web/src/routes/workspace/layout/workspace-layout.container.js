import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'

import { WorkspaceLayout as WorkspaceLayoutView } from './workspace-layout.view'

const mapStateToProps = (state) => ({
})

const mapDispatchToProps = (dispatch) => ({
})

const makeContainer = (component) => {
  return connect(
    mapStateToProps,
    mapDispatchToProps
  )(component)
}

export const WorkspaceLayout = makeContainer(WorkspaceLayoutView)
