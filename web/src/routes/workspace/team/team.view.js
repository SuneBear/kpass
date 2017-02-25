import React, { Component, PropTypes } from 'react'

import './team.view.styl'

export class Team extends Component {

  static propTypes = {
    actions: PropTypes.object,
    params: PropTypes.object,
    children: PropTypes.element
  }

  componentDidMount () {
    const { teamId } = this.props.params

    this.props.actions.setCurrentTeam({
      teamId
    })
  }

  componentWillReceiveProps (nextProps) {
    const { teamId } = nextProps.params

    this.props.actions.setCurrentTeam({
      teamId
    })
  }

  render () {
    return (
      <div className={'teamView'}>
        {this.props.children}
      </div>
    )
  }

}
