import React, { Component, PropTypes } from 'react'
import cx from 'classnames'

import './card.view.styl'

export class Card extends Component {

  static propTypes = {
    className: PropTypes.string,
    children: PropTypes.element,
    title: PropTypes.string
  }

  getRootClassnames () {
    return cx(
      'cardView',
      { withTitle: !!this.props.title },
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

  render () {
    return (
      <div className={this.getRootClassnames()}>
        {this.renderCardTitle()}
        <div className={'cardBody'}>
          {this.props.children}
        </div>
      </div>
    )
  }

}
