import React, { Component, PropTypes } from 'react'
import cx from 'classnames'

import { Loading } from 'uis'
import { getCreatorPermissions } from 'utils'
import { EntriesListCell } from '../entries-list-cell'

import './entries-list.view.styl'

export class EntriesList extends Component {

  static propTypes = {
    className: PropTypes.string,
    userMe: PropTypes.object,
    team: PropTypes.object,
    entries: PropTypes.array
  }

  getRootClassnames () {
    return cx(
      'entriesListView',
      this.props.className
    )
  }

  renderCell (entry) {
    const {
      userMe,
      team
    } = this.props

    const creatorPermissions = getCreatorPermissions(
      entry.userID,
      userMe,
      team
    )

    return (
      <EntriesListCell
        key={entry.id}
        entry={entry}
        creatorPermissions={creatorPermissions}
       />
    )
  }

  renderCells () {
    const {
      entries
    } = this.props

    return entries.map((entry) => {
      return this.renderCell(entry)
    })
  }

  renderEndPoint () {
    return (
      <Loading isEnded />
    )
  }

  render () {
    return (
      <div className={this.getRootClassnames()}>
        {this.renderCells()}
        {this.renderEndPoint()}
      </div>
    )
  }

}
