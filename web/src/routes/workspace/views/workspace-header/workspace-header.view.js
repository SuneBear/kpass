import React, { Component, PropTypes } from 'react'
import { I18n } from 'react-redux-i18n'

import { Icon, Dropdown, MenuSelector } from 'uis'
import { Avatar, Logo } from 'views'
import { isPublicTeam } from 'utils'
import { getWorkspaceBashPath } from '../../index'

import './workspace-header.view.styl'

export class WorkspaceHeader extends Component {

  static propTypes = {
    actions: PropTypes.object,
    userMe: PropTypes.object,
    teams: PropTypes.array,
    currentTeam: PropTypes.object
  }

  renderWorkspaceInfo () {
    const { currentTeam } = this.props

    return (
      <Dropdown
        className={'workspaceSwitcherDropdown'}
        content={this.getWorkspaceSwitcher()}
      >
        <div className={'workspaceInfo workspaceSwitcherHandler'}>
          <div className={'workspaceName'}>
            {currentTeam.name}
          </div>
          <Icon className={'handlerIcon'} name={'chevron-down'} />
        </div>
      </Dropdown>
    )
  }

  handleSwitchWorkspace = (teamSelector) => {
    const { teams } = this.props
    const { push } = this.props.actions
    const nextTeam = teams.filter(team => teamSelector.value === team.id)[0]
    push(getWorkspaceBashPath(nextTeam))
  }

  handleCreateTeam = () => {
    // Not implemented
    console.log('Create Team')
  }

  getWorkspaceSwitcher () {
    const { currentTeam, teams } = this.props

    const dataList = teams.map((team) => ({
      className: 'workspaceSwitcherItem',
      value: team.id,
      title: team.name,
      iconName: isPublicTeam(team) ? 'building' : 'user',
      onClick: this.handleSwitchWorkspace
    }))

    const extraList = [
      {
        className: 'workspaceSwitcherItem',
        title: I18n.t('team.newTeam'),
        iconName: 'plus',
        onClick: this.handleCreateTeam
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

  renderAppLogo () {
    return (
      <div className={'workspaceLogo'}>
        <Logo className={'defaultLogo'} height={23} />
      </div>
    )
  }

  renderUserInfo () {
    const { id, avatar } = this.props.userMe

    return (
      <Dropdown
        content={this.getUserInfoDropdownMenu()}
        placement={'bottomRight'}
        offset={[-8, 8]}
      >
        <div className={'workspaceUserInfo'}>
          <Avatar className={'infoAvatar'} url={avatar} size={'small'} />
          <span className={'infoUsername'}>{id}</span>
        </div>
      </Dropdown>
    )
  }

  getUserInfoDropdownMenu () {
    const dataList = [
      { title: I18n.t('account.settings'), onClick: () => { console.log('settings') } },
      { type: 'divider' },
      { title: I18n.t('account.signOut'), onClick: () => { console.log('signOut') } }
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
        {this.renderAppLogo()}
        {this.renderUserInfo()}
      </div>
    )
  }

}
