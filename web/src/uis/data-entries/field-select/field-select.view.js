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
    defaultValue: PropTypes.string,
    placeholder: PropTypes.string,
    options: PropTypes.array,

    // Redux from
    input: PropTypes.object
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

  handleChange = (e) => {
    this.props.input.onChange(e)
  }

  renderSelect () {
    const {
      input,
      name,
      defaultValue,
      placeholder,
      options
    } = this.props

    if (!input.value) {
      input.value = defaultValue
    }

    return (
      <div className='selectWrap'>
        <Select
          name={name}
          defaultValue={input.value}
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
