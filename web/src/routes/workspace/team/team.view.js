import React, { Component, PropTypes } from 'react'

import './team.view.styl'

export class Team extends Component {

  static propTypes = {
    children: PropTypes.element
  }

  render () {
    return (
      <div className={'teamView'}>
        {this.props.children}
      </div>
    )
  }

}
