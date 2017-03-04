import React, { Component, PropTypes } from 'react'
import { I18n, Translate } from 'react-redux-i18n'
import cx from 'classnames'

import { Button, Modal } from 'uis'
import { Card, Placeholder } from 'views'
import { EntriesList } from './entries-list'
import { EntryMake } from '../entry-make'

import './entries.view.styl'

export class Entries extends Component {

  static propTypes = {
    className: PropTypes.string,
    userMe: PropTypes.object,
    team: PropTypes.object,
    entries: PropTypes.array,
    userPermissions: PropTypes.object,
    actions: PropTypes.object
  }

  componentWillReceiveProps (nextProps) {
    const { team, actions } = this.props

    if (
      !nextProps.team.id ||
      nextProps.team.id === team.id
    ) {
      return
    }

    actions.readCurrentTeamEntries({
      teamId: nextProps.team.id
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

  renderPlaceholder () {
    const { entries } = this.props

    if (entries) {
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
    const { userMe, entries } = this.props

    if (!entries) {
      return null
    }

    return (
      <EntriesList
        userMe={userMe}
        entries={entries}
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
        {this.renderPlaceholder()}
        {this.renderEntries()}

        {this.renderNewEntryModal()}
      </Card>
    )
  }

}
