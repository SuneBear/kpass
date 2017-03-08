import React, { Component, PropTypes } from 'react'
import cx from 'classnames'

import './badge.view.styl'

export class Badge extends Component {

  static propTypes = {
    className: PropTypes.string,
    size: PropTypes.number,
    count: PropTypes.number,
    max: PropTypes.number,
    bgColor: PropTypes.string
  }

  static defaultProps = {
    count: null,
    max: 99,
    size: 18
  }

  getRootClassnames () {
    return cx(
      'badgeView',
      this.props.className
    )
  }

  getRootStyles () {
    let {
      count,
      size,
      bgColor
    } = this.props

    if (count === null) {
      size = 10
    }

    const styles = {
      minWidth: size,
      height: size,
      borderRadius: size / 2,
      backgroundColor: bgColor
    }

    return styles
  }

  renderCount () {
    const { max, count } = this.props

    if (count === null) {
      return null
    }

    const limitedCount = count > max ? `${max}+` : `${count}`

    return (
      <span className={'badgeCount'}>
        {limitedCount}
      </span>
    )
  }

  render () {
    return (
      <div
        className={this.getRootClassnames()}
        style={this.getRootStyles()}
      >
        {this.renderCount()}
      </div>
    )
  }

}
