import React, { Component, PropTypes } from 'react'
import cx from 'classnames'

import { isOwner, isMe } from 'utils'
import { MembersListCell } from '../members-list-cell'

import './members-list.view.styl'

export class MembersList extends Component {

  static propTypes = {
    className: PropTypes.string,
    userMe: PropTypes.object,
    team: PropTypes.object,
    members: PropTypes.array,
    userPermissions: PropTypes.object,
    onRemoveMember: PropTypes.func
  }

  getRootClassnames () {
    return cx(
      'membersListView',
      this.props.className
    )
  }

  sortMembers (ownerId, userMeId, members) {
    return members
      .sort((prev, next) => {
        if (prev.id === userMeId) {
          return -1
        } else if (next.id === userMeId) {
          return 1
        } else {
          return 0
        }
      })
      .sort((prev, next) => {
        if (prev.id === ownerId) {
          return -1
        } else if (next.id === ownerId) {
          return 1
        } else {
          return 0
        }
      })
  }

  renderCell (member) {
    const {
      userMe,
      team,
      userPermissions,
      onRemoveMember
    } = this.props

    return (
      <MembersListCell
        key={member.id}
        member={member}
        isOwner={isOwner(team, member)}
        isMe={isMe(member, userMe)}
        userPermissions={userPermissions}
        onRemoveMember={onRemoveMember}
       />
    )
  }

  renderCells () {
    const {
      userMe,
      team,
      members
    } = this.props

    if (!members) {
      return null
    }

    const sortedMembers = this.sortMembers(
      team.userId,
      userMe.id,
      members
    )

    return sortedMembers.map((member) => {
      return this.renderCell(member)
    })
  }

  render () {
    return (
      <div className={this.getRootClassnames()}>
        {this.renderCells()}
      </div>
    )
  }

}
