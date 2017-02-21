import React, { Component, PropTypes } from 'react'
import cx from 'classnames'

import './logo.view.styl'

export class Logo extends Component {

  static propTypes = {
    className: PropTypes.string,
    url: PropTypes.string,
    height: PropTypes.number
  }

  static defaultProps = {
    url: require('assets/logo.svg'),
    height: 28
  }

  getRootClassnames () {
    return cx(
      'logoView',
      this.props.className
    )
  }

  getLogoSize () {
    return {
      height: this.props.height
    }
  }

  render () {
    return (
      <img
        className={this.getRootClassnames()}
        style={this.getLogoSize()}
        src={this.props.url}
      />
    )
  }

}
