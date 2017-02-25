import React, { Component, PropTypes } from 'react'
import cx from 'classnames'

import './card.view.styl'

export class Card extends Component {

  static propTypes = {
    className: PropTypes.string,
    children: PropTypes.any,
    title: PropTypes.string,
    withoutPadding: PropTypes.bool,
    handler: PropTypes.element
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
    return (
      <div className={this.getRootClassnames()}>
        {this.renderCardHeader()}
        <div className={'cardBody'}>
          {this.props.children}
        </div>
      </div>
    )
  }

}
