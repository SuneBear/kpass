import React, { Component, PropTypes } from 'react'
import { I18n, Translate } from 'react-redux-i18n'
import cx from 'classnames'

import { Button } from 'uis'
import { Avatar } from 'views'

import './members-list-cell.view.styl'

export class MembersListCell extends Component {

  static propTypes = {
    className: PropTypes.string,
    member: PropTypes.object,
    isMe: PropTypes.bool,
    isOwner: PropTypes.bool,
    onLeaveTeam: PropTypes.func,
    onRemoveMember: PropTypes.func,
    permissions: PropTypes.object
  }

  getRootClassnames () {
    return cx(
      'membersListCellView',
      this.props.className
    )
  }

  handleLeaveTeam = () => {
    const { onLeaveTeam } = this.props

    onLeaveTeam()
  }

  handleRemoveMember = () => {
    const { onRemoveMember, member } = this.props

    onRemoveMember({
      memberId: member.id
    })
  }

  renderUserInfo () {
    const { member } = this.props

    return (
      <div className={'userInfoSection'}>
        <Avatar url={member.avatarUrl} />
        <span className={'userName'}>{member.id}</span>
      </div>
    )
  }

  renderRole () {
    const { isOwner } = this.props

    const roleName = isOwner
      ? I18n.t('role.owner')
      : I18n.t('role.member')

    return (
      <div className={'roleSection'}>
        <div className={'roleName'}>{roleName}</div>
      </div>
    )
  }

  renderHandler () {
    const { isOwner, isMe, permissions } = this.props

    let handler = null
    if (isMe) {
      const handlerText = isOwner
        ? I18n.t('team.leaveAndDisband')
        : I18n.t('team.leave')

      handler = (
        <Button
          ghost
          type={'danger'}
          size={'small'}
          onClick={this.handleLeaveTeam}
        >
          {handlerText}
        </Button>
      )
    } else if (permissions.deleteTeamMember) {
      handler = (
        <Button
          ghost
          type={'danger'}
          size={'small'}
          onClick={this.handleRemoveMember}
        >
          <Translate value={'teamMembers.remove'} />
        </Button>
      )
    }

    return (
      <div className={'handlerSection'}>
        {handler}
      </div>
    )
  }

  render () {
    return (
      <div className={this.getRootClassnames()}>
        {this.renderUserInfo()}
        {this.renderRole()}
        {this.renderHandler()}
      </div>
    )
  }

}
