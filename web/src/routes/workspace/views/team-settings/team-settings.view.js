import React, { Component, PropTypes } from 'react'
import { I18n, Translate } from 'react-redux-i18n'
import cx from 'classnames'

import { Button, Toggle, Modal } from 'uis'
import { Card } from 'views'
import { TeamRename } from '../team-rename'

import './team-settings.view.styl'

export class TeamSettings extends Component {

  static propTypes = {
    className: PropTypes.string,
    team: PropTypes.object,
    actions: PropTypes.object
  }

  getRootClassnames () {
    return cx(
      'teamSettingsView',
      this.props.className
    )
  }

  saveTeamRenameModalRef = (ref) => {
    this.teamRenameModalRef = ref
  }

  handleTeamRenameClick = () => {
    this.teamRenameModalRef.open()
  }

  handleFreezeToggleChange = (checked) => {
    const {
      actions
    } = this.props

    // @TODO: Implementation
    actions.updateCurrentTeam({
      isFrozen: checked
    })
  }

  renderTeamRenameModal () {
    const { team } = this.props

    return (
      <Modal
        ref={this.saveTeamRenameModalRef}
        title={I18n.t('teamSettings.rename')}
        size={'small'}
      >
        <TeamRename
          team={team}
          initialValues={team}
        />
      </Modal>
    )
  }

  renderTeamRenameSection () {
    return (
      <div className={'settingSection'}>
        <div className={'settingSectionTitle'}>
          <Translate value={'teamSettings.rename'} />
        </div>
        <div className={'settingSectionDescription'}>
          <Translate value={'teamSettings.renameDescription'} />
        </div>
        <div className={'settingSectionHandler'}>
          <Button
            ghost
            size={'small'}
            onClick={this.handleTeamRenameClick}
          >
            <Translate value={'teamSettings.rename'} />
          </Button>
        </div>
      </div>
    )
  }

  renderFreezeToggleSection () {
    const { team } = this.props

    return (
      <div className={'settingSection'}>
        <div className={'settingSectionTitle'}>
          <Translate value={'teamSettings.freeze'} />
        </div>
        <div className={'settingSectionDescription'}>
          <Translate value={'teamSettings.freezeDescription'} />
        </div>
        <div className={'settingSectionHandler'}>
          <Toggle
            checked={team.isFrozen}
            onChange={this.handleFreezeToggleChange}
          />
        </div>
      </div>
    )
  }

  render () {
    return (
      <Card
        className={this.getRootClassnames()}
        title={I18n.t('teamSettings.title')}
      >
        {this.renderTeamRenameSection()}
        {this.renderFreezeToggleSection()}

        {this.renderTeamRenameModal()}
      </Card>
    )
  }

}
