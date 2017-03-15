import React, { Component, PropTypes } from 'react'
import { I18n, Translate } from 'react-redux-i18n'
import cx from 'classnames'

import { Button, Loading, Modal } from 'uis'
import { Card, Placeholder } from 'views'
import { EntriesList } from './entries-list'
import { EntryMake } from '../entry-make'
import { getEntryPathById } from '../../index'

import './entries.view.styl'

export class Entries extends Component {

  static propTypes = {
    params: PropTypes.object,
    className: PropTypes.string,
    userMe: PropTypes.object,
    team: PropTypes.object,
    currentEntry: PropTypes.object,
    entries: PropTypes.array,
    entriesFilter: PropTypes.string,
    entriesBashPath: PropTypes.string,
    userPermissions: PropTypes.object,
    actions: PropTypes.object
  }

  componentWillReceiveProps (nextProps) {
    const { params, currentEntry, entriesFilter, actions } = this.props

    if (
      !params.filterName ||
      entriesFilter !== nextProps.params.filterName
    ) {
      actions.setCurrentFilter({
        filter: nextProps.params.filterName
      })
    }

    if (
      nextProps.params.entryId === currentEntry.id
    ) {
      return
    }

    actions.setCurrentEntry({
      entryId: nextProps.params.entryId
    })
  }

  getRootClassNames () {
    return cx(
      'entriesView',
      this.props.className
    )
  }

  getCardTitle () {
    const count = this.getEntriesCount()

    if (!count) {
      return I18n.t('entries.title')
    }

    return (
      I18n.t('entries.title') +
      ' · ' +
      count
    )
  }

  getEntriesCount () {
    const {
      entries
    } = this.props

    if (!entries) {
      return null
    }

    return entries.length
  }

  saveNewEntryModalRef = (ref) => {
    this.newEntryModalRef = ref
  }

  handleNewEntryClick = () => {
    this.newEntryModalRef.open()
  }

  handleNewEntrySubmitSuccess = () => {
    this.newEntryModalRef.close()
  }

  handleEntryCellClick = (entry) => {
    const { entriesBashPath, actions } = this.props

    actions.push(getEntryPathById(
      entriesBashPath, entry.id
    ))
  }

  handleEntryModalClose = () => {
    const { entriesBashPath, actions } = this.props

    actions.push(entriesBashPath)
  }

  getNewEntryHandler (type, isGhost) {
    const { userPermissions } = this.props

    if (!userPermissions.createEntry) {
      return null
    }

    return (
      <Button
        type={type}
        ghost={isGhost}
        icon={'circle-plus'}
        onClick={this.handleNewEntryClick}
      >
        <Translate value={'entry.new'} />
      </Button>
    )
  }

  renderNewEntryModal () {
    const initialvalues = {
      category: 'Login'
    }

    return (
      <Modal
        ref={this.saveNewEntryModalRef}
        title={I18n.t('entry.new')}
        size={'small'}
      >
        <EntryMake
          initialValues={initialvalues}
          onSubmitSuccess={this.handleNewEntrySubmitSuccess}
        />
      </Modal>
    )
  }

  renderLoading () {
    const { entries } = this.props

    if (entries) {
      return null
    }

    return (
      <Loading />
    )
  }

  renderPlaceholder () {
    const { entries } = this.props

    if (!entries || entries.length !== 0) {
      return null
    }

    return (
      <Placeholder
        imageName={'task'}
        title={I18n.t('entries.placeholderTitle')}
        description={I18n.t('entries.placeholderDescription')}
        handler={this.getNewEntryHandler('primary', true)}
       />
    )
  }

  renderEntries () {
    const {
      userMe,
      team,
      currentEntry,
      entries
    } = this.props

    if (!entries || !entries.length) {
      return null
    }

    return (
      <EntriesList
        userMe={userMe}
        team={team}
        willOpenEntry={currentEntry}
        entries={entries}
        onCellClick={this.handleEntryCellClick}
        onCellModalClose={this.handleEntryModalClose}
      />
    )
  }

  render () {
    return (
      <Card
        className={this.getRootClassNames()}
        title={this.getCardTitle()}
        handler={this.getNewEntryHandler('text')}
      >
        {this.renderLoading()}
        {this.renderPlaceholder()}
        {this.renderEntries()}

        {this.renderNewEntryModal()}
      </Card>
    )
  }

}
