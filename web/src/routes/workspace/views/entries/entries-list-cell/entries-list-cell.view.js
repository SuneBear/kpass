import React, { Component, PropTypes } from 'react'
import cx from 'classnames'

import { Card } from 'views'
import { EntryItem } from '../../entry-item'

import './entries-list-cell.view.styl'

export class EntriesListCell extends Component {

  static propTypes = {
    className: PropTypes.string,
    entry: PropTypes.object,
    creatorPermissions: PropTypes.object
  }

  getRootClassnames () {
    return cx(
      'entriesListCellView',
      this.props.className
    )
  }

  render () {
    const {
      entry,
      creatorPermissions
    } = this.props

    return (
      <Card className={this.getRootClassnames()}>
        <EntryItem
          entry={entry}
          creatorPermissions={creatorPermissions}
        />
      </Card>
    )
  }

}
