import React, { Component, PropTypes } from 'react'
import { findDOMNode } from 'react-dom'
import cx from 'classnames'

import './loading.view.styl'

export class Loading extends Component {

  static propTypes = {
    className: PropTypes.string,
    isEnded: PropTypes.bool,
    hideAnimated: PropTypes.bool
  }

  componentWillLeave (callback) {
    const indicator = findDOMNode(this).firstElementChild

    if (this.props.hideAnimated) {
      const duration = 418
      indicator.classList.add('hideAnimated')
      window.setTimeout(() => callback(), duration)
    } else {
      callback()
    }
  }

  getRootClassNames () {
    return cx(
      'loadingWrap',
      this.props.className
    )
  }

  renderLoadingIndicator () {
    if (this.props.isEnded) {
      return null
    }

    return (
      <div className='loadingIndicator'>
        <div className='loaderDot' />
        <div className='loaderDot' />
        <div className='loaderDot' />
      </div>
    )
  }

  renderEndPoint () {
    if (!this.props.isEnded) {
      return null
    }

    return (
      <div className='endPointWrapper'>
        <span className='endPoint' />
      </div>
    )
  }

  render () {
    return (
      <div className={this.getRootClassNames()}>
        {this.renderLoadingIndicator()}
        {this.renderEndPoint()}
      </div>
    )
  }

}
