import React, { Component, PropTypes } from 'react'
import { I18n, Translate } from 'react-redux-i18n'
import cx from 'classnames'

import { isMe } from 'utils'
import { Button, Loading, Modal } from 'uis'
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
    userPermissions: PropTypes.object,
    actions: PropTypes.object
  }

  constructor (props) {
    super(props)

    this.state = {
      inviteCode: null
    }
  }

  saveMemberInviteModalRef = (ref) => {
    this.memberInviteModalRef = ref
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

    if (!teamMembers) {
      return 0
    }

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
    const { userPermissions } = this.props

    if (!userPermissions.createTeamMember) {
      return null
    }

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

  handleRemoveMember = ({ member }) => {
    const { userMe, team, actions } = this.props

    actions.removeMember({
      team,
      teamId: team.id,
      memberId: member.id,
      isMe: isMe(member, userMe)
    })
  }

  handleAddMemberClick = () => {
    this.memberInviteModalRef.open()
  }

  handleMemberInviteSubmitSuccess = (response) => {
    this.state.inviteCode = response.code
    this.setState(this.state)
  }

  handleUsernameChange = () => {
    this.state.inviteCode = null
    this.setState(this.state)
  }

  renderMemberInviteModal () {
    return (
      <Modal
        ref={this.saveMemberInviteModalRef}
        title={I18n.t('teamMembers.invite')}
        size={'small'}
        onClose={this.handleUsernameChange}
      >
        <MemberInvite
          inviteCode={this.state.inviteCode}
          onUsernameChange={this.handleUsernameChange}
          onSubmitSuccess={this.handleMemberInviteSubmitSuccess}
        />
      </Modal>
    )
  }

  renderMemberList () {
    const {
      userMe,
      team,
      teamMembers,
      userPermissions
    } = this.props

    return (
      <MembersList
        userMe={userMe}
        team={team}
        members={teamMembers}
        userPermissions={userPermissions}
        onRemoveMember={this.handleRemoveMember}
      />
    )
  }

  renderLoading () {
    const { teamMembers } = this.props

    if (teamMembers) {
      return null
    }

    return (
      <Loading />
    )
  }

  render () {
    return (
      <Card
        withoutPadding
        className={this.getRootClassNames()}
        title={this.getCardTitle()}
        handler={this.getAddMemberHandler()}
      >
        {this.renderLoading()}
        {this.renderMemberList()}

        {this.renderMemberInviteModal()}
      </Card>
    )
  }

}
