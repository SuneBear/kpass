import React, { Component, PropTypes } from 'react'
import cx from 'classnames'

import { Popup } from '../popup'

import './dropdown.view.styl'

export class Dropdown extends Component {

  static propTypes = {
    className: PropTypes.string,
    prefixCls: PropTypes.string,
    transitionCls: PropTypes.string,
    action: PropTypes.oneOf(['click', 'hover', 'focus']),
    placement: PropTypes.string,
    opened: PropTypes.bool,
    mountIn: PropTypes.func,
    content: PropTypes.element,
    nested: PropTypes.bool,
    offset: PropTypes.array,
    onOpen: PropTypes.func,
    onHide: PropTypes.func
  }

  static defaultProps = {
    prefixCls: 'dropdown',
    transitionCls: '',
    offset: [0, 8]
  }

  constructor (props) {
    super(props)
    this.ref = Object.create(null)
    this.content = this.getContent(this.props.content)
  }

  savePopup = (el) => {
    this.ref = el
  }

  onClick = (e) => {
    const { content } = this.props

    if (this.ref) {
      this.ref.close()
    }

    if (content.onClick) {
      content.onClick(e)
    }
  }

  getContentClassnames (content) {
    return cx(
      'dropdownMenu',
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
      onClick: this.onClick,
      className: this.getContentClassnames(content)
    })
  }

  componentWillReceiveProps(nextProps) {
    if (nextProps.content !== this.props.content) {
      this.content = this.getContent(nextProps.content)
    }
  }

  render() {
    const props = {
      ...this.props,
      content: this.content
    }

    return <Popup {...props}  ref={this.savePopup} />
  }

}
