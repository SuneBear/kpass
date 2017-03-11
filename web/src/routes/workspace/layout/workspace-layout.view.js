import React, { Component, PropTypes } from 'react'

import { Card } from 'views'
import { WorkspaceHeader, WorkspaceSidebar } from '../views'

import './workspace-layout.view.styl'

export class WorkspaceLayout extends Component {

  static propTypes = {
    children: PropTypes.element,
    actions: PropTypes.object
  }

  componentWillMount () {
    const { actions } = this.props

    actions.readTeams()
  }

  render () {
    return (
      <div className={'workspaceLayout'}>
        <WorkspaceHeader />
        <div className={'workspaceMain'}>
          <Card className={'workspaceSidebar'}>
            <WorkspaceSidebar />
          </Card>
          <div className={'workspaceContent'}>
            {this.props.children}
          </div>
        </div>
      </div>
    )
  }

}
