import React, { Component, PropTypes } from 'react'
import { I18n } from 'react-redux-i18n'

import { Icon, Dropdown, Modal, MenuSelector } from 'uis'
import { isPublicTeam, getFileUrl } from 'utils'
import { Avatar, Logo } from 'views'
import { TeamCreate } from '../team-create'
import { AccountSettings } from '../account-settings'
import { getWorkspaceBashPath } from '../../index'

import './workspace-header.view.styl'

export class WorkspaceHeader extends Component {

  static propTypes = {
    userMe: PropTypes.object,
    teams: PropTypes.array,
    currentTeam: PropTypes.object,
    actions: PropTypes.object
  }

  saveTeamCreateModalRef = (ref) => {
    this.teamCreateModalRef = ref
  }

  saveAccountSettingsModalRef = (ref) => {
    this.accountSettingsModalRef = ref
  }

  getTeamName (team) {
    if (isPublicTeam(team)) {
      return team.name
    } else {
      return I18n.t('workspace.personal')
    }
  }

  handleNewTeamClick = () => {
    this.teamCreateModalRef.open()
  }

  handleAccountSettingsClick = () => {
    this.accountSettingsModalRef.open()
  }

  renderTeamCreateModal () {
    return (
      <Modal
        ref={this.saveTeamCreateModalRef}
        title={I18n.t('team.newTeam')}
        size={'small'}
      >
        <TeamCreate />
      </Modal>
    )
  }

  renderAccountSettingsModal () {
    return (
      <Modal
        opened
        ref={this.saveAccountSettingsModalRef}
        title={I18n.t('account.settings')}
      >
        <AccountSettings />
      </Modal>
    )
  }

  renderWorkspaceInfo () {
    const { currentTeam } = this.props

    return (
      <div className={'workspaceInfoWrap'}>
        <Dropdown
          className={'workspaceSwitcherDropdown'}
          content={this.getWorkspaceSwitcher()}
        >
          <div className={'workspaceInfo workspaceSwitcherHandler'} title={currentTeam.name}>
            <div className={'workspaceName'}>
              {this.getTeamName(currentTeam)}
            </div>
            <Icon className={'handlerIcon'} name={'chevron-down'} />
          </div>
        </Dropdown>
      </div>
    )
  }

  handleSwitchWorkspace = (teamSelector) => {
    const { teams } = this.props
    const { push } = this.props.actions
    const nextTeam = teams.filter(team => teamSelector.value === team.id)[0]
    push(getWorkspaceBashPath(nextTeam))
  }

  getWorkspaceSwitcher () {
    const { currentTeam, teams } = this.props

    const dataList = teams.map((team) => ({
      className: 'workspaceSwitcherItem',
      value: team.id,
      title: this.getTeamName(team),
      iconName: isPublicTeam(team) ? 'building' : 'user',
      onClick: this.handleSwitchWorkspace
    }))

    const extraList = [
      {
        className: 'workspaceSwitcherItem',
        title: I18n.t('team.newTeam'),
        iconName: 'plus',
        onClick: this.handleNewTeamClick
      }
    ]

    return (
      <MenuSelector
        dataList={dataList}
        extraList={extraList}
        hasSelected={[currentTeam.id]}
      />
    )
  }

  handleWorkspaceLogoClick = () => {
    const { currentTeam } = this.props
    const { push } = this.props.actions
    push(getWorkspaceBashPath(currentTeam))
  }

  renderWrokspaceLogo () {
    return (
      <div className={'workspaceLogo'} onClick={this.handleWorkspaceLogoClick}>
        <Logo className={'defaultLogo'} height={23} />
      </div>
    )
  }

  renderUserInfo () {
    const { id, avatar } = this.props.userMe

    return (
      <div className={'workspaceUserInfoWrap'}>
        <Dropdown
          content={this.getUserInfoDropdownMenu()}
          placement={'bottomRight'}
          offset={[-8, 8]}
        >
          <div className={'workspaceUserInfo'}>
            <Avatar className={'infoAvatar'} url={getFileUrl(avatar)} size={'small'} />
            <span className={'infoUsername'}>{id}</span>
          </div>
        </Dropdown>
      </div>
    )
  }

  handleSignOutClick = () => {
    this.props.actions.signOutUser()
  }

  getUserInfoDropdownMenu () {
    const dataList = [
      { title: I18n.t('account.settings'), onClick: this.handleAccountSettingsClick },
      { type: 'divider' },
      { title: I18n.t('account.signOut'), onClick: this.handleSignOutClick }
    ]

    return (
      <MenuSelector
        dataList={dataList}
      />
    )
  }

  render () {
    return (
      <div className={'workspaceHeaderView'}>
        {this.renderWorkspaceInfo()}
        {this.renderWrokspaceLogo()}
        {this.renderUserInfo()}

        {this.renderTeamCreateModal()}
        {this.renderAccountSettingsModal()}
      </div>
    )
  }

}
