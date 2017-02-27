import React, { Component } from 'react'

// @REF: https://ant.design/components/switch/
import { Switch as AntSwitch } from 'antd'

import './toggle.view.styl'

export class Toggle extends Component {

  static propTypes = {
    ...AntSwitch.propTypes
  }

  static defaultProps = {
    prefixCls: 'toggle'
  }

  render () {
    const {
      ...props
    } = this.props

    return (
      <AntSwitch
        {...props}
      />
    )
  }
}
