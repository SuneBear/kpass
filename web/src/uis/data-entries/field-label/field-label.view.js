import React, { Component, PropTypes } from 'react'
import cx from 'classnames'

import './field-label.view.styl'

export class FieldLabel extends Component {

  static PropTypes = {
    className: PropTypes.string,
    htmlFor: PropTypes.string,
    text: PropTypes.string
  }

  getRootClassNames () {
    return cx(
      'fieldLabel',
      this.props.className
    )
  }

  render () {
    const {
      htmlFor,
      text
    } = this.props

    return (
      <label className={this.getRootClassNames()} htmlFor={htmlFor}>
        {text}
      </label>
    )
  }

}
