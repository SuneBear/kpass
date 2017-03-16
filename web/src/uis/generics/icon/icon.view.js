import React, { Component, PropTypes } from 'react'

import './icon.view.styl'

export class Icon extends Component {

  static propTypes = {
    name: PropTypes.string.isRequired,
    className: PropTypes.string,
    Component: PropTypes.oneOfType([PropTypes.string, PropTypes.func]),
    onClick: PropTypes.func
  }

  static defaultProps = {
    Component: 'span'
  }

  render () {
    const {
      Component,
      name, className, ...props
    } = this.props

    let classNames = `icon icon-${name}`

    if (className) {
      classNames = `${classNames} ${className}`
    }

    return <Component {...props} className={classNames} />
  }

}
