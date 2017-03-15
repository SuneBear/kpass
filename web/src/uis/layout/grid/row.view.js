import React, { Component } from 'react'
import cx from 'classnames'

// @REF: https://github.com/loicmahieu/react-styled-flexboxgrid
import { Row as FlexboxgridRow } from 'react-styled-flexboxgrid'

export class Row extends Component {

  static propTypes = {
    ...FlexboxgridRow.propTypes
  }

  getRootClassNames () {
    return cx(
      'gridRow',
      this.props.className
    )
  }

  render () {
    return (
      <FlexboxgridRow
        {...this.props}
        className={this.getRootClassNames()}
      />
    )
  }

}
