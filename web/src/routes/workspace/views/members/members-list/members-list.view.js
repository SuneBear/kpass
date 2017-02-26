import React, { Component, PropTypes } from 'react'
import cx from 'classnames'

import { MembersListCell } from '../members-list-cell'

import './members-list.view.styl'

export class MembersList extends Component {

  static propTypes = {
    className: PropTypes.string,
    userMe: PropTypes.object,
    team: PropTypes.object,
    members: PropTypes.array,
    onLeaveTeam: PropTypes.func,
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

  isOwner (member) {
    const { team } = this.props

    return team.userID === member.id
  }

  isMe (member) {
    const { userMe } = this.props

    return userMe.id === member.id
  }

  getPermissions (member) {
    const { userMe } = this.props

    const permissions = {
      deleteMember: false
    }

    if (this.isOwner(userMe)) {
      permissions.deleteMember = true
    }

    return permissions
  }

  renderCell (member) {
    const {
      onLeaveTeam,
      onRemoveMember
    } = this.props

    return (
      <MembersListCell
        key={member.id}
        member={member}
        onLeaveTeam={onLeaveTeam}
        onRemoveMember={onRemoveMember}
        isMe={this.isMe(member)}
        isOwner={this.isOwner(member)}
        permissions={this.getPermissions(member)}
       />
    )
  }

  renderCells () {
    const {
      userMe,
      team,
      members
    } = this.props

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
