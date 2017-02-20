import React, { Component, PropTypes } from 'react'
import { Translate } from 'react-redux-i18n'

import { Icon, Dropdown } from 'uis'
import { Avatar, Logo } from 'views'

import './workspace-header.view.styl'

export class WorkspaceHeader extends Component {

  static propTypes = {
    userMe: PropTypes.object,
    currentTeam: PropTypes.object
  }

  renderWorkspaceInfo () {
    const { currentTeam } = this.props

    return (
      <div className={'workspaceInfo workspaceSwitcherHandler'}>
        <div className={'workspaceName'}>
          {currentTeam.name}
        </div>
        <Icon className={'handlerIcon'} name={'chevron-down'} />
      </div>
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
        content={this.getUserInfoDropdownItems()}
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

  getUserInfoDropdownItems () {
    return (
      <ul>
        <li><a><Translate value={'account.settings'} /></a></li>
        <div className={'divider'} />
        <li><a><Translate value={'account.signOut'} /></a></li>
      </ul>
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
