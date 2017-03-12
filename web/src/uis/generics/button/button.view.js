import React, { Component, PropTypes } from 'react'
import cx from 'classnames'
import { Icon } from '../icon'
import { capitalize } from 'utils/string'

import './button.view.styl'

export class Button extends Component {

  static propTypes = {
    className: PropTypes.string,
    children: PropTypes.oneOfType([
      PropTypes.element,
      PropTypes.string
    ]),
    onClick: PropTypes.func,
    icon: PropTypes.string,
    type: PropTypes.oneOf(['primary', 'text', 'danger', 'normal']),
    htmlType: PropTypes.oneOf(['submit', 'button', 'reset']),
    size: PropTypes.oneOf(['small', 'normal', 'large']),
    block: PropTypes.bool,
    loading: PropTypes.bool,
    disabled: PropTypes.bool,
    ghost: PropTypes.bool
  }

  static defaultProps = {
    type: 'primary',
    size: 'normal',
    onClick: () => {}
  }

  getRootClassNames () {
    const {
      className,
      type,
      size,
      block,
      loading,
      disabled,
      ghost
    } = this.props

    return cx(
      'button',
      `button${capitalize(type)}`,
      `size${capitalize(size)}`,
      { isBlock: block },
      { isGhost: ghost },
      { isLoading: loading },
      { isDisabled: disabled },
      className
    )
  }

  handleClick = (e) => {
    const {
      loading,
      disabled,
      onClick
    } = this.props

    if (loading || disabled) {
      e.preventDefault()
      return null
    }

    if (onClick) {
      onClick(e)
    }
  }

  render () {
    const {
      icon,
      htmlType,
      // loading,
      children
    } = this.props

    const iconName = icon // @OLD: loading ? 'loading' : icon
    const iconNode = iconName ? <Icon name={iconName} /> : null

    return (
      <button
        className={this.getRootClassNames()}
        onClick={this.handleClick}
        type={htmlType || 'button'}
      >
        {iconNode}{children}
      </button>
    )
  }

}

