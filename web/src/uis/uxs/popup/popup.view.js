import React, { Component, PropTypes } from 'react'
import { findDOMNode } from 'react-dom'
import Trigger from 'rc-trigger'

import { Position } from '../position'

import './popup.view.styl'

export class Popup extends Component {

  static propTypes = {
    style: PropTypes.object,
    className: PropTypes.string,
    children: PropTypes.children,
    prefixCls: PropTypes.string,
    transitionCls: PropTypes.string,
    action: PropTypes.oneOf(['click', 'hover', 'focus']),
    placement: PropTypes.string,
    opened: PropTypes.bool,
    mask: PropTypes.bool,
    maskCloseable: PropTypes.bool,
    mountIn: PropTypes.func,
    content: PropTypes.element,
    nested: PropTypes.bool,
    offset: PropTypes.array,
    onOpen: PropTypes.func,
    onHide: PropTypes.func,
    adjust: PropTypes.bool
  }

  static defaultProps = {
    action: 'click',
    prefixCls: 'popup',
    opened: false,
    mask: false,
    maskCloseable: true,
    nested: false,
    transitionCls: ''
  }

  constructor (props) {
    super(props)
    this.state = {
      visible: this.props.opened
    }
  }

  componentDidMount () {
    if (this.props.nested) {
      this.parentEl = findDOMNode(this).parentElement
    }
    document.addEventListener('keyup', this.escapeHandler)
  }

  componentWillUnmount () {
    document.removeEventListener('keyup', this.escapeHandler)
  }

  getAction () {
    const { action } = this.props
    return typeof action === 'string' ? [action] : action
  }

  findNestedParent = () => {
    return () => this.parentEl
  }

  visibleChanged = (status) => {
    this.state.visible = status
    this.setState(this.state)

    const { onHide, onOpen } = this.props
    if (status) {
      return onOpen && onOpen()
    } else {
      return onHide && onHide()
    }
  }

  escapeHandler = (e) => {
    if (e.which === 27 && this.state.visible) {
      this.close()
    }
  }

  open () {
    this.state.visible = true
    this.setState(this.state)
  }

  close () {
    this.state.visible = false
    this.setState(this.state)
  }

  render () {
    const {
      mask,
      maskCloseable,
      mountIn,
      content,
      prefixCls,
      placement,
      nested,
      transitionCls,
      adjust,
      offset,
      className,
      style
    } = this.props

    const transtion = transitionCls ? `${prefixCls}-${transitionCls}` : ''

    return (
      <Position placement={placement} offset={offset} adjust={adjust}>
        <Trigger
          action={this.getAction()}
          mask={mask}
          maskCloseable={maskCloseable}
          popupVisible={this.state.visible}
          getPopupContainer={nested ? this.findNestedParent() : mountIn}
          popup={content}
          prefixCls={prefixCls}
          popupTransitionName={transtion}
          onPopupVisibleChange={this.visibleChanged}
          popupClassName={className}
          style={style}
        >
          {this.props.children}
        </Trigger>
      </Position>
    )
  }

}
