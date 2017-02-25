import React, { Component, PropTypes } from 'react'
import { findDOMNode } from 'react-dom'
import cx from 'classnames'

// @REF: https://ant.design/components/input/
import { Select as AntSelect } from 'antd'

import { Icon } from '../../generics/icon'

import './select.view.styl'

export class Select extends Component {

  static propTypes = {
    ...AntSelect.propTypes,
    options: PropTypes.arrayOf(
      PropTypes.shape({
        title: PropTypes.string,
        value: PropTypes.string.isRequired
      })
    )
  }

  static defaultProps = {
    prefixCls: 'select',
    transitionName: '',
    choiceTransitionName: ''
  }

  static Option = AntSelect.Option

  // @Hack: Replace className
  componentDidMount () {
    const $el = findDOMNode(this)
    const $arrow = $el.querySelector('.select-arrow')
    $arrow.classList.add('icon')
    $arrow.classList.add('icon-chevron-down')
  }

  getOptionItems () {
    const { Option } = Select
    const { children, options } = this.props

    if (children) {
      return children
    }

    if (!options) {
      return null
    }

    return options.map((option) => {
      const value = option.value
      const text = option.title || value
      return (
        <Option value={value}>
          {text}<Icon className={'selectedLabel'} name={'tick'} />
        </Option>
      )
    })
  }

  getDropdownClassNames () {
    return cx(
      'popup',
      this.props.dropdownClassName
    )
  }

  render () {
    return (
      <AntSelect
        {...this.props}
        children={this.getOptionItems()}
        dropdownClassName={this.getDropdownClassNames()}
      />
    )
  }

}
