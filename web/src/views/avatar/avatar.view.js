import React, { Component, PropTypes } from 'react'
import cx from 'classnames'

import { capitalize } from 'utils'

import './avatar.view.styl'

export class Avatar extends Component {

  static propTypes = {
    className: PropTypes.string,
    url: PropTypes.string.isRequired,
    size: PropTypes.string
  }

  static defaultProps = {
    size: 'normal'
  }

  getRootClassnames () {
    return cx(
      'avatarView',
      [`size${capitalize(this.props.size)}`],
      this.props.className
    )
  }

  getAvatarUrl () {
    return {
      backgroundImage: `url(${this.props.url})`
    }
  }

  render () {
    return (
      <div
        className={this.getRootClassnames()}
        style={this.getAvatarUrl()}
      />
    )
  }

}
