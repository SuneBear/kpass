import React, { Component, PropTypes } from 'react'
import { I18n, Translate } from 'react-redux-i18n'
import cx from 'classnames'

import { Button, Modal } from 'uis'
import { Card } from 'views'
import { MembersList } from './members-list'
import { MemberInvite } from '../member-invite'

import './members.view.styl'

export class Members extends Component {

  static propTypes = {
    className: PropTypes.string,
    userMe: PropTypes.object,
    team: PropTypes.object,
    teamMembers: PropTypes.array,
    actions: PropTypes.object
  }

  refMemberInviteModal = null

  saveRefMemberInviteModal = (ref) => {
    this.refMemberInviteModal = ref
  }

  getRootClassNames () {
    return cx(
      'membersView',
      this.props.className
    )
  }

  getMembersCount () {
    const {
      teamMembers
    } = this.props

    return teamMembers.length
  }

  getCardTitle () {
    return (
      I18n.t('teamMembers.title') +
      ' · ' +
      this.getMembersCount()
    )
  }

  getAddMemberHandler () {
    return (
      <Button
        type={'text'}
        icon={'circle-plus'}
        onClick={this.handleAddMemberClick}
      >
        <Translate value={'member.add'} />
      </Button>
    )
  }

  handleLeaveTeam = () => {
    const { team, userMe } = this.props
    const { leaveTeam } = this.props.actions

    // @TODO: Implementation
    leaveTeam({
      teamId: team.id,
      memberId: userMe.id
    })
  }

  handleRemoveMember = (memberId) => {
    const { team } = this.props
    const { removeTeamMember } = this.props.actions

    // @TODO: Implementation
    removeTeamMember({
      teamId: team.id,
      memberId
    })
  }

  handleAddMemberClick = () => {
    this.refMemberInviteModal.open()
  }

  renderMemberInviteModal () {
    return (
      <Modal
        ref={this.saveRefMemberInviteModal}
        title={I18n.t('teamMembers.invite')}
        size={'small'}
      >
        <MemberInvite />
      </Modal>
    )
  }

  render () {
    const {
      userMe,
      team,
      teamMembers
    } = this.props

    return (
      <Card
        withoutPadding
        className={this.getRootClassNames()}
        title={this.getCardTitle()}
        handler={this.getAddMemberHandler()}
      >
        {this.renderMemberInviteModal()}
        <MembersList
          userMe={userMe}
          team={team}
          members={teamMembers}
          onLeaveTeam={this.handleLeaveTeam}
          onRemoveMember={this.handleRemoveMember}
        />
      </Card>
    )
  }

}
