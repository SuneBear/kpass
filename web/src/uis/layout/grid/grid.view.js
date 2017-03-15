import React, { Component } from 'react'
import cx from 'classnames'

// @REF: https://github.com/loicmahieu/react-styled-flexboxgrid
import { Grid as FlexboxgridGrid } from 'react-styled-flexboxgrid'

export class Grid extends Component {

  static propTypes = {
    ...FlexboxgridGrid.propTypes
  }

  getRootClassNames () {
    return cx(
      'grid',
      this.props.className
    )
  }

  render () {
    return (
      <FlexboxgridGrid
        {...this.props}
        className={this.getRootClassNames()}
      />
    )
  }

}
