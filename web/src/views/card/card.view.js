import React, { Component, PropTypes } from 'react'
import cx from 'classnames'

import './card.view.styl'

export class Card extends Component {

  static propTypes = {
    className: PropTypes.string,
    children: PropTypes.any,
    title: PropTypes.string,
    withoutPadding: PropTypes.bool,
    handler: PropTypes.element,
    onClick: PropTypes.func
  }

  getRootClassnames () {
    return cx(
      'cardView',
      { withTitle: !!this.props.title },
      { withoutPadding: this.props.withoutPadding },
      this.props.className
    )
  }

  renderCardTitle () {
    const {
      title
    } = this.props

    if (!title) {
      return null
    }

    return (
      <div className={'cardTitle'}>{title}</div>
    )
  }

  renderCardHandler () {
    const {
      handler
    } = this.props

    if (!handler) {
      return null
    }

    return (
      <div className={'cardHandler'}>{handler}</div>
    )
  }

  renderCardHeader () {
    const {
      title
    } = this.props

    if (!title) {
      return null
    }

    return (
      <div className={'cardHeader'}>
        {this.renderCardTitle()}
        {this.renderCardHandler()}
      </div>
    )
  }

  render () {
    const { onClick } = this.props

    return (
      <div
        className={this.getRootClassnames()}
        onClick={onClick}
      >
        {this.renderCardHeader()}
        <div className={'cardBody'}>
          {this.props.children}
        </div>
      </div>
    )
  }

}
