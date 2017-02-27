import React, { Component, PropTypes } from 'react'
import { I18n, Translate } from 'react-redux-i18n'
import { Field, propTypes as formPropTypes } from 'redux-form'
import cx from 'classnames'

import {
  Button, Toggle, Modal,
  FieldText
} from 'uis'
import { Card } from 'views'

import './team-settings.view.styl'

export class TeamSettings extends Component {

  static propTypes = {
    className: PropTypes.string,
    currentTeam: PropTypes.object,
    actions: PropTypes.object,
    ...formPropTypes
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

  handleTeamRenameSubmit = (values) => {
    const {
      actions
    } = this.props

    // @TODO: Implementation
    actions.updateCurrentTeam({
      name: values.teamName
    })
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

  renderTeamRenameForm () {
    const {
      handleSubmit,
      pristine,
      valid,
      submitting,

      currentTeam
    } = this.props

    return (
      <form
        onSubmit={handleSubmit(this.handleTeamRenameSubmit)}
      >
        <Field
          name={'teamName'}
          component={FieldText}
          placeholder={I18n.t('team.teamName')}
          defaultValue={currentTeam.name}
        />
        <Button
          block
          type={'primary'}
          htmlType={'submit'}
          disabled={pristine || !valid}
          loading={submitting}
        >
          <Translate value={'action.save'} />
        </Button>
      </form>
    )
  }

  renderTeamRenameModal () {
    return (
      <Modal
        ref={this.saveTeamRenameModalRef}
        title={I18n.t('teamSettings.rename')}
        size={'small'}
      >
        {this.renderTeamRenameForm()}
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
    const { currentTeam } = this.props

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
            checked={currentTeam.isFrozen}
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
