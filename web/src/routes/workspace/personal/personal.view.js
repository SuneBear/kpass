import React, { Component, PropTypes } from 'react'

import './personal.view.styl'

export class Personal extends Component {

  static propTypes = {
    children: PropTypes.element,
    privateTeamId: PropTypes.string,
    actions: PropTypes.object
  }

  componentDidMount () {
    const { privateTeamId, actions } = this.props

    if (!privateTeamId) {
      return
    }

    actions.mountCurrentTeam({
      teamId: privateTeamId
    })
  }

  componentWillReceiveProps (nextProps) {
    const { privateTeamId, actions } = this.props

    if (
      !nextProps.privateTeamId ||
      nextProps.privateTeamId === privateTeamId
    ) {
      return
    }

    actions.mountCurrentTeam({
      teamId: nextProps.privateTeamId
    })
  }

  componentWillUnmount () {
    const { actions } = this.props

    actions.unmountCurrentTeam()
  }

  render () {
    return (
      <div className={'workspaceType personalView'}>
        {this.props.children}
      </div>
    )
  }

}
