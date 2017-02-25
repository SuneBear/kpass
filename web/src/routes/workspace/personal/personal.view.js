import React, { Component, PropTypes } from 'react'

import './personal.view.styl'

export class Personal extends Component {

  static propTypes = {
    actions: PropTypes.object,
    children: PropTypes.element
  }

  componentDidMount () {
    this.props.actions.setCurrentTeam({
      teamId: null
    })
  }

  componentWillReceiveProps (nextProps) {
    this.props.actions.setCurrentTeam({
      teamId: null
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
