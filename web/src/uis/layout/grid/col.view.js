import React, { Component } from 'react'
import cx from 'classnames'

// @REF: https://github.com/loicmahieu/react-styled-flexboxgrid
import { Col as FlexboxgridCol } from 'react-styled-flexboxgrid'

export class Col extends Component {

  static propTypes = {
    ...FlexboxgridCol.propTypes
  }

  getRootClassNames () {
    return cx(
      'gridCol',
      this.props.className
    )
  }

  render () {
    return (
      <FlexboxgridCol
        {...this.props}
        className={this.getRootClassNames()}
      />
    )
  }

}
