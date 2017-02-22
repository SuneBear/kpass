import React, { Component, PropTypes } from 'react'
import cx from 'classnames'

import { Select } from '../select'
import { FieldLabel } from '../field-label'

import './field-select.view.styl'

export class FieldSelect extends Component {

  static propTypes = {
    className: PropTypes.string,
    label: PropTypes.string,
    name: PropTypes.string,
    placeholder: PropTypes.string,
    options: PropTypes.array,

    // Redux from
    input: PropTypes.object,
    meta: PropTypes.object
  }

  getRootClassNames () {
    return cx(
      'fieldSelect',
      this.props.className
    )
  }

  renderLabel () {
    const {
      label,
      name
    } = this.props

    if (!label) {
      return null
    }

    return (
      <FieldLabel
        htmlFor={name}
        text={label}
      />
    )
  }

  handleFocus = (e) => {
    this.props.input.onFocus(e)
  }

  handleBlur = (e) => {
    this.props.input.onBlur(e)
  }

  handleChange = (e) => {
    this.props.input.onChange(e)
  }

  renderSelect () {
    const {
      input,
      name,
      placeholder,
      options
    } = this.props

    return (
      <div className='selectWrap'>
        <Select
          defaultValue={input.value}
          name={name}
          placeholder={placeholder}
          options={options}
          onChange={this.handleChange}
        />
      </div>
    )
  }

  render () {
    return (
      <div className={this.getRootClassNames()}>
        {this.renderLabel()}
        {this.renderSelect()}
      </div>
    )
  }

}
