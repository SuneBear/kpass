import React, { Component, PropTypes } from 'react'
import cx from 'classnames'

import { Modal } from 'uis'
import { Card } from 'views'
import { EntryItem } from '../../entry-item'
import { EntryDetail } from '../../entry-detail'

import './entries-list-cell.view.styl'

export class EntriesListCell extends Component {

  static propTypes = {
    className: PropTypes.string,
    entry: PropTypes.object,
    creatorPermissions: PropTypes.object,
    onClick: PropTypes.func
  }

  getRootClassnames () {
    return cx(
      'entriesListCellView',
      this.props.className
    )
  }

  saveEntryDetailModalRef = (ref) => {
    this.entryDetailModalRef = ref
  }

  handleCellClick = (e) => {
    const { onClick } = this.props

    if (onClick) {
      onClick(e)
    }

    this.entryDetailModalRef.open()
  }

  renderEntryDetailModal () {
    return (
      <Modal
        ref={this.saveEntryDetailModalRef}
        className={'entryDetailModal'}
      >
        {this.renderEntryDetail()}
      </Modal>
    )
  }

  renderEntryItem () {
    const {
      entry,
      creatorPermissions
    } = this.props

    return (
      <EntryItem
        entry={entry}
        creatorPermissions={creatorPermissions}
      />
    )
  }

  renderEntryDetail () {
    const {
      entry,
      creatorPermissions
    } = this.props

    return (
      <EntryDetail
        entry={entry}
        creatorPermissions={creatorPermissions}
      />
    )
  }

  render () {
    return (
      <Card
        className={this.getRootClassnames()}
        onClick={this.handleCellClick}
      >
        {this.renderEntryItem()}

        {this.renderEntryDetailModal()}
      </Card>
    )
  }

}
