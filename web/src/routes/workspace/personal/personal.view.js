import React, { Component, PropTypes } from 'react'

import './personal.view.styl'

export class Personal extends Component {

  static PropTypes = {
    userMe: PropTypes.object
  }

  render () {
    return (
      <div className='personalView'>
        <div>Personal View</div>
        { this.props.children }
      </div>
    )
  }

}
