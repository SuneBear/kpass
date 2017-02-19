import React, { Component, PropTypes } from 'react'
import cx from 'classnames'

import { capitalize } from 'utils/string'
import { Icon } from '../../generics/icon'

import './toast.content.view.styl'

export class ToastContent extends Component {

  static propTypes = {
    type: PropTypes.string.isRequired,
    message: PropTypes.string.isRequired,
    description: PropTypes.string
  }

  getRootClassnames () {
    return cx(
      'toastContent',
      [`type${capitalize(this.props.type)}`]
    )
  }

  renderIcon () {
    let iconName = 'circle-info'
    switch (this.props.type) {
      case 'info':
        iconName = 'circle-info'
        break
      case 'warning':
        iconName = 'circle-warning'
        break
      case 'error':
        iconName = 'circle-remove'
        break
      case 'success':
        iconName = 'state-check'
        break
    }

    return (
      <Icon name={iconName} />
    )
  }

  renderDesctiption () {
    if (!this.props.description) {
      return null
    }

    return (
      <div className={'toastDescription'}>{this.props.description}</div>
    )
  }

  render () {
    return (
      <div className={this.getRootClassnames()}>
        {this.renderIcon()}
        <div className={'toastMessage'}>{this.props.message}</div>
        {this.renderDesctiption()}
      </div>
    )
  }

}
