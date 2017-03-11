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
    willOpenEntry: PropTypes.object,
    entry: PropTypes.object,
    creatorPermissions: PropTypes.object,
    onClick: PropTypes.func,
    onModalClose: PropTypes.func
  }

  getRootClassnames () {
    return cx(
      'entriesListCellView',
      this.props.className
    )
  }

  shouldEntryDetailOpenModal () {
    const { willOpenEntry, entry } = this.props
    return willOpenEntry.id === entry.id
  }

  saveEntryDetailModalRef = (ref) => {
    this.entryDetailModalRef = ref
  }

  handleCellClick = () => {
    const { entry, onClick } = this.props

    if (onClick) {
      onClick(entry)
    }
  }

  renderEntryDetailModal () {
    const { onModalClose } = this.props

    return (
      <Modal
        ref={this.saveEntryDetailModalRef}
        opened={this.shouldEntryDetailOpenModal()}
        className={'entryDetailModal'}
        onClose={onModalClose}
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
