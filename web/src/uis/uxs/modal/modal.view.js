import React, { Component, PropTypes } from 'react'
import Dialog from 'rc-dialog'
import cx from 'classnames'

import { Icon } from 'uis'
import { NOOP } from 'utils/constants'

import './modal.view.styl'

export class Modal extends Component {

  static propTypes = {
    children: PropTypes.element,
    style: PropTypes.object,
    size: PropTypes.oneOf(['small', 'normal', 'middle', 'large']),
    zIndex: PropTypes.number,
    title: PropTypes.string,
    className: PropTypes.string,
    prefixCls: PropTypes.string,
    transitionCls: PropTypes.string,
    opened: PropTypes.bool,
    centered: PropTypes.bool,
    onOpen: PropTypes.func,
    onClose: PropTypes.func,
    mask: PropTypes.object,
    footer: PropTypes.oneOfType([
      PropTypes.element,
      PropTypes.func
    ]),
    close: PropTypes.oneOfType([
      PropTypes.element,
      PropTypes.func
    ])
  }

  static defaultProps = {
    prefixCls: 'modal',
    transitionCls: 'slide',
    size: 'normal',
    opened: false,
    centered: false,
    onOpen: NOOP,
    onClose: NOOP,
    mask: {
      transitionCls: 'fade',
      visible: true,
      closable: true,
      style: null
    },
    close: <Icon className={'modal-close'} name={'remove'} />
  }

  constructor (props) {
    super(props)

    this.state = {
      visible: null
    }
  }

  // @TODO: Distinguish `props.opened` and `state.visible`
  componentWillReceiveProps (nextProps, nextState) {
    if (nextProps.opened === this.props.opened) {
      return
    }

    if (nextProps.opened) {
      this.open()
    } else {
      this.close()
    }
  }

  componentDidMount () {
    if (this.props.opened) {
      this.open()
    }
  }

  getRootClassNames () {
    const {
      className,
      size
    } = this.props

    return cx(
      className,
      [`size-${size}`]
    )
  }

  open = () => {
    this.props.onOpen()

    this.state.visible = true
    this.setState(this.state)
  }

  close = () => {
    this.props.onClose()

    this.state.visible = false
    this.setState(this.state)
  }

  optionalProps () {
    const {
      prefixCls,
      transitionCls,
      centered,
      mask
    } = this.props

    const wrapClassName =
      centered
      ? 'center-align'
      : null

    const transitionName =
      transitionCls
      ? `${prefixCls}-${transitionCls}`
      : null

    const maskConfig =
      Modal.defaultProps.mask === mask
      ? mask
      : { ...Modal.defaultProps.mask, mask }

    let maskTrans
    if (!!maskConfig.transitionCls && !!transitionCls) {
      maskTrans = `${prefixCls}-mask-${maskConfig.transitionCls}`
    }

    return {
      prefixCls,
      wrapClassName,
      transitionName,
      maskClosable: maskConfig.closable,
      mask: maskConfig.visible,
      maskStyle: maskConfig.style,
      maskTransitionName: maskTrans
    }
  }

  renderClose = () => {
    if (!this.props.close) {
      return null
    }

    const closeEl = this.props.close

    if (typeof closeEl === 'function') {
      // Stateless Component
      return React.createElement(closeEl, {
        extraHandler: this.close
      })
    } else {
      const clickBinding = closeEl && closeEl.props.onClick
      return React.cloneElement(closeEl, {
        onClick: (e) => {
          if (typeof clickBinding === 'function') {
            clickBinding(e)
          }
          this.close()
        }
      })
    }
  }

  render () {
    const {
      style,
      title,
      zIndex,
      footer
    } = this.props

    if (!this.state.visible) {
      return null
    }

    return (
      <Dialog
        visible={this.state.visible}
        closable={false}
        style={style}
        className={this.getRootClassNames()}
        title={title}
        zIndex={zIndex}
        footer={footer}
        onClose={this.close}
        {...this.optionalProps()}
      >
        {this.renderClose()}
        <div className={'scrollable-viewport'}>
          {this.props.children}
        </div>
      </Dialog>
    )
  }

}
