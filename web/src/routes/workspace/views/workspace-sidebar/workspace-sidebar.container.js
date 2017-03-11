import { connect } from 'react-redux'

import { routingSelector } from 'modules'
import { currentBasePathSelector, currentUserPermissionsSelector } from '../../modules'
import { WorkspaceSidebar as WorkspaceSidebarView } from './workspace-sidebar.view'

const mapStateToProps = (state) => ({
  basePath: currentBasePathSelector(state),
  routing: routingSelector(state),
  userPermissions: currentUserPermissionsSelector(state)
})

const mapDispatchToProps = (dispatch) => ({})

const makeContainer = (component) => {
  return connect(
    mapStateToProps,
    mapDispatchToProps
  )(component)
}

export const WorkspaceSidebar = makeContainer(WorkspaceSidebarView)
