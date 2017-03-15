import React, { Component, PropTypes } from 'react'
import cx from 'classnames'

import { Popup } from '../popup'

import './dropdown.view.styl'

export class Dropdown extends Component {

  static propTypes = {
    ...Popup.propTypes,
    onClick: PropTypes.func
  }

  static defaultProps = {
    prefixCls: 'dropdown',
    transitionCls: '',
    offset: [0, 8]
  }

  constructor (props) {
    super(props)

    this.refPopup = Object.create(null)
    this.content = this.getContent(this.props.content)
  }

  componentWillReceiveProps (nextProps) {
    if (nextProps.content !== this.props.content) {
      this.content = this.getContent(nextProps.content)
    }
  }

  savePopup = (el) => {
    this.refPopup = el
  }

  handleClick = (e) => {
    const { content } = this.props

    if (this.refPopup) {
      this.refPopup.close()
    }

    if (content.onClick) {
      content.onClick(e)
    }
  }

  getRootClassNames () {
    return cx(
      'popup',
      this.props.className
    )
  }

  getContentClassnames (content) {
    return cx(
      'dropdownContent',
      content.props.className
    )
  }

  getContent (content) {
    if (!content || !content.props) {
      return (
        <noscript />
      )
    }

    return React.cloneElement(content, {
      className: this.getContentClassnames(content),
      onClick: this.handleClick
    })
  }

  render () {
    const props = {
      ...this.props,
      content: this.content
    }

    return (
      <Popup
        {...props}
        className={this.getRootClassNames()}
        ref={this.savePopup}
      />
    )
  }

}
