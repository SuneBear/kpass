import React, { Component, PropTypes } from 'react'
import { Translate } from 'react-redux-i18n'
import cx from 'classnames'

import { Button, Loading, Modal } from 'uis'
import { Card } from 'views'
import { EntryItem } from '../entry-item'
import { Secrets } from '../secrets'
import { SecretMake } from '../secret-make'

import './entry-detail.view.styl'

export class EntryDetail extends Component {

  static propTypes = {
    className: PropTypes.string,
    entry: PropTypes.object,
    creatorPermissions: PropTypes.object,
    actions: PropTypes.object
  }

  componentWillMount () {
    const { entry, actions } = this.props

    actions.readEntry({
      entryId: entry.id
    })
  }

  getRootClassnames () {
    return cx(
      'entryDetailView',
      this.props.className
    )
  }

  saveAddSecretModalRef = (ref) => {
    this.addSecretModalRef = ref
  }

  handleAddSecretClick = () => {
    this.addSecretModalRef.open()
  }

  handleAddSecretSubmitSuccess = () => {
    this.addSecretModalRef.close()
  }

  renderEntryHeader () {
    const {
      entry,
      creatorPermissions
    } = this.props

    return (
      <div className={'entryDetailHeader'}>
        <EntryItem
          entry={entry}
          creatorPermissions={creatorPermissions}
        />
      </div>
    )
  }

  renderAddSecretHandler () {
    const { creatorPermissions } = this.props

    if (!creatorPermissions.createSecret) {
      return null
    }

    return (
      <div className={'entryDetailContentSection'}>
        <Button
          block
          icon={'plus-sign'}
          type={'normal'}
          onClick={this.handleAddSecretClick}
        >
          <Translate value={'secret.new'} />
        </Button>
      </div>
    )
  }

  renderAddSecretModal () {
    return (
      <Modal
        ref={this.saveAddSecretModalRef}
        className={'secretMakeModal'}
      >
        <SecretMake
          onSubmitSuccess={this.handleAddSecretSubmitSuccess}
        />
      </Modal>
    )
  }

  renderLoading () {
    const { entry } = this.props

    if (entry.secrets) {
      return null
    }

    return (
      <Card className={'entryDetailContentSection'}>
        <Loading />
      </Card>
    )
  }

  renderSecrets () {
    const { entry } = this.props

    if (!entry.secrets) {
      return null
    }

    return (
      <div className={'entryDetailContentSection'}>
        <Secrets />
      </div>
    )
  }

  render () {
    return (
      <div
        className={this.getRootClassnames()}
      >
        {this.renderEntryHeader()}
        <div className={'entryDetailDivider'} />
        <div className={'entryDetailContent'}>
          {this.renderAddSecretHandler()}
          {this.renderLoading()}
          {this.renderSecrets()}
        </div>

        {this.renderAddSecretModal()}
      </div>
    )
  }

}
