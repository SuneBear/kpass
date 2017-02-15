import React, { Component, PropTypes } from 'react'

import './workspace-layout.view.styl'

export class WorkspaceLayout extends Component {

  render () {
    return (
      <div className='workspaceLayout'>
        { this.props.children }
      </div>
    )
  }

}
