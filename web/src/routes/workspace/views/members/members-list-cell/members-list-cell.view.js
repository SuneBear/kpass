import React, { Component, PropTypes } from 'react'
import { I18n, Translate } from 'react-redux-i18n'
import cx from 'classnames'

import { getFileUrl } from 'utils'
import { Button } from 'uis'
import { Avatar } from 'views'

import './members-list-cell.view.styl'

export class MembersListCell extends Component {

  static propTypes = {
    className: PropTypes.string,
    member: PropTypes.object,
    isMe: PropTypes.bool,
    isOwner: PropTypes.bool,
    userPermissions: PropTypes.object,
    onRemoveMember: PropTypes.func
  }

  getRootClassnames () {
    return cx(
      'membersListCellView',
      this.props.className
    )
  }

  handleRemoveMember = () => {
    const { onRemoveMember, member } = this.props

    onRemoveMember({
      member
    })
  }

  renderUserInfo () {
    const { member } = this.props

    return (
      <div className={'userInfoSection'}>
        <Avatar url={getFileUrl(member.avatar)} />
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
    const { isOwner, isMe, userPermissions } = this.props

    let handler = null
    if (isMe && !isOwner) { // @ALT: isMe
      const handlerText = isOwner
        ? I18n.t('teamMembers.leaveAndDisband')
        : I18n.t('teamMembers.leave')

      handler = (
        <Button
          ghost
          type={'danger'}
          size={'small'}
          onClick={this.handleRemoveMember}
        >
          {handlerText}
        </Button>
      )
    }

    if (!isMe && userPermissions.deleteTeamMember) {
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
