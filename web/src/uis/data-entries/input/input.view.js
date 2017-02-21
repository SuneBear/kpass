import React, { Component } from 'react'
import cx from 'classnames'

// @REF: https://ant.design/components/input/
import { Input as AntInput } from 'antd'

import './input.view.styl'

export class Input extends Component {

  static propTypes = AntInput.propTypes

  static defaultProps = {
    prefixCls: 'input'
  }

  getRootClassNames () {
    return cx(
      this.props.className
    )
  }

  render () {
    const {
      ...props
    } = this.props

    return (
      <AntInput
        className={this.getRootClassNames()}
        {...props}
      />
    )
  }
}
