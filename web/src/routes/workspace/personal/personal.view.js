import React, { Component, PropTypes } from 'react'

import './personal.view.styl'

export class Personal extends Component {

  static propTypes = {
    children: PropTypes.element,
    privateTeamId: PropTypes.string,
    actions: PropTypes.object
  }

  componentWillUpdate (nextProps) {
    const { privateTeamId } = nextProps

    this.props.actions.setCurrentTeam({
      teamId: privateTeamId
    })
  }

  componentDidMount () {
    const { privateTeamId } = this.props

    this.props.actions.setCurrentTeam({
      teamId: privateTeamId
    })
  }

  render () {
    return (
      <div className={'personalView'}>
        {this.props.children}
      </div>
    )
  }

}
