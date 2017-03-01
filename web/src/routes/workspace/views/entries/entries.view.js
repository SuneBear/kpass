import React, { Component, PropTypes } from 'react'
import { I18n, Translate } from 'react-redux-i18n'
import cx from 'classnames'

import { Button, Modal } from 'uis'
import { Card, Placeholder } from 'views'
import { EntryMake } from '../entry-make'

import './entries.view.styl'

export class Entries extends Component {

  static propTypes = {
    className: PropTypes.string,
    entries: PropTypes.array
  }

  getRootClassNames () {
    return cx(
      'entriesView',
      this.props.className
    )
  }

  handleNewEntryClick = () => {
    this.newEntryModalRef.open()
  }

  saveNewEntryModalRef = (ref) => {
    this.newEntryModalRef = ref
  }

  getNewEntryHandler (type, isGhost) {
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
    return (
      <Modal
        ref={this.saveNewEntryModalRef}
        title={I18n.t('entry.new')}
        size={'small'}
      >
        <EntryMake />
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
    const { entries } = this.props

    if (!entries) {
      return null
    }

    return entries.map((entry) => (
      <div>{entry.name}</div>
    ))
  }

  render () {
    return (
      <Card
        className={this.getRootClassNames()}
        title={I18n.t('entries.title')}
        handler={this.getNewEntryHandler('text')}
      >
        {this.renderPlaceholder()}
        {this.renderEntries()}

        {this.renderNewEntryModal()}
      </Card>
    )
  }

}
