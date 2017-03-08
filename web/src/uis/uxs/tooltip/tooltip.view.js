import React, { Component } from 'react'

// @REF: https://ant.design/components/tooltip/
import { Tooltip as AntTooltip } from 'antd'

import './tooltip.view.styl'

export class Tooltip extends Component {

  static propTypes = {
    ...AntTooltip.propTypes
  }

  static defaultProps = {
    ...AntTooltip.defaultProps,
    prefixCls: 'tooltip',
    transitionName: 'fade-transition'
  }

  render () {
    const {
      ...props
    } = this.props

    return (
      <AntTooltip
        ref={'input'}
        {...props}
      />
    )
  }
}
