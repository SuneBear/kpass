import React, { Component, PropTypes } from 'react'
import cx from 'classnames'

import './entry-item.view.styl'

export class EntryItem extends Component {

  static propTypes = {
    className: PropTypes.string,
    entry: PropTypes.object,
    creatorPermissions: PropTypes.object,
    actions: PropTypes.object
  }

  getRootClassnames () {
    return cx(
      'entryItemView',
      this.props.className
    )
  }

  render () {
    const {
      entry
    } = this.props

    return (
      <div className={this.getRootClassnames()}>
        {entry.name}
      </div>
    )
  }

}
