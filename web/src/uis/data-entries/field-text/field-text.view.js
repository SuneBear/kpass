import React, { Component, PropTypes } from 'react'
import cx from 'classnames'

import { Input } from '../input'
import { FieldLabel } from '../field-label'

import './field-text.view.styl'

export class FieldText extends Component {

  static PropTypes = {
    className: PropTypes.string,
    input: {
      value: PropTypes.string,
      onBlur: PropTypes.func,
      onChange: PropTypes.func,
      onFocus: PropTypes.func
    },
    label: PropTypes.string,
    meta: {
      error: PropTypes.bool,
      touched: PropTypes.bool
    },
    name: PropTypes.string.isRequired,
    placeholder: PropTypes.string,
    autoFocus: PropTypes.bool,
    type: PropTypes.string
  }

  getRootClassNames () {
    const {
      className,
      meta
    } = this.props

    return cx(
      'fieldText',
      meta.touched && meta.error ? 'isError' : null,
      className
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

  renderInput () {
    const {
      value,
      name,
      placeholder,
      autoFocus,
      type
    } = this.props

    return (
      <div className='inputWrap'>
        <Input
          defaultValue={value}
          name={name}
          placeholder={placeholder}
          autoFocus={autoFocus}
          type={type}
          onFocus={this.handleFocus}
          onBlur={this.handleBlur}
          onChange={this.handleChange}
        />
      </div>
    )
  }

  render () {
    return (
      <div className={this.getRootClassNames()}>
        {this.renderLabel()}
        {this.renderInput()}
      </div>
    )
  }

}
