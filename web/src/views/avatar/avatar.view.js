import React, { Component, PropTypes } from 'react'
import cx from 'classnames'

import { capitalize } from 'utils'

import './avatar.view.styl'

const DEFAULT_AVATAR_URL = require('./assets/default-avator@2x.png')

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
    const avatarUrl = this.props.url || DEFAULT_AVATAR_URL

    return {
      backgroundImage: `url(${avatarUrl})`
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
