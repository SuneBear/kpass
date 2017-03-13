import React, { Component, PropTypes } from 'react'
import cx from 'classnames'

import { Input } from '../input'
import { FieldLabel } from '../field-label'

import './field-text.view.styl'

export class FieldText extends Component {

  static propTypes = {
    className: PropTypes.string,
    label: PropTypes.string,
    name: PropTypes.string,
    defaultValue: PropTypes.string,
    placeholder: PropTypes.string,
    autosize: PropTypes.object,
    autoFocus: PropTypes.bool,
    prefix: PropTypes.element,
    type: PropTypes.string,
    onFocus: PropTypes.func,
    onBlur: PropTypes.func,
    onChange: PropTypes.func,

    // Redux from
    input: PropTypes.object,
    meta: PropTypes.object
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

  saveInputRef = (ref) => {
    this.inputRef = ref
  }

  focus () {
    window.setTimeout(() =>
      this.inputRef.focus()
    , 10)
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
    if (this.props.onFocus) {
      this.props.onFocus(e)
    }
    this.props.input.onFocus(e)
  }

  handleBlur = (e) => {
    if (this.props.onBlur) {
      this.props.onBlur(e)
    }
    this.props.input.onBlur(e)
  }

  handleChange = (e) => {
    if (this.props.onChange) {
      this.props.onChange(e)
    }
    this.props.input.onChange(e)
  }

  renderInput () {
    const {
      input,
      name,
      placeholder,
      defaultValue,
      autosize,
      autoFocus,
      prefix,
      type
    } = this.props

    if (!input.value) {
      input.value = defaultValue
    }

    return (
      <div className='inputWrap'>
        <Input
          ref={this.saveInputRef}
          defaultValue={input.value}
          name={name}
          placeholder={placeholder}
          autosize={autosize}
          autoFocus={autoFocus}
          prefix={prefix}
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
