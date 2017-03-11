import React, { Component, PropTypes } from 'react'
import cx from 'classnames'

import { getCreatorPermissions } from 'utils'
import { Loading } from 'uis'
import { EntriesListCell } from '../entries-list-cell'

import './entries-list.view.styl'

export class EntriesList extends Component {

  static propTypes = {
    className: PropTypes.string,
    userMe: PropTypes.object,
    team: PropTypes.object,
    willOpenEntry: PropTypes.object,
    entries: PropTypes.array,
    onCellClick: PropTypes.func,
    onCellModalClose: PropTypes.func
  }

  getRootClassnames () {
    return cx(
      'entriesListView',
      this.props.className
    )
  }

  handleCellClick = (entry) => {
    const { onCellClick } = this.props

    if (onCellClick) {
      onCellClick(entry)
    }
  }

  handleCellModalClose = () => {
    const { onCellModalClose } = this.props

    if (onCellModalClose) {
      onCellModalClose()
    }
  }

  renderCell (entry) {
    const {
      userMe,
      team,
      willOpenEntry
    } = this.props

    const creatorPermissions = getCreatorPermissions(
      entry.userID,
      userMe,
      team
    )

    return (
      <EntriesListCell
        key={entry.id}
        willOpenEntry={willOpenEntry}
        entry={entry}
        creatorPermissions={creatorPermissions}
        onClick={this.handleCellClick}
        onModalClose={this.handleCellModalClose}
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
