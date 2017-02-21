import React, { Component, PropTypes } from 'react'

import './personal.view.styl'

export class Personal extends Component {

  static propTypes = {
    children: PropTypes.element
  }

  render () {
    return (
      <div className={'personalView'}>
        {this.props.children}
      </div>
    )
  }

}
