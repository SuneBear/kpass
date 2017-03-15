import React, { Component, PropTypes } from 'react'
import { I18n, Translate } from 'react-redux-i18n'
import CopyToClipboard from 'react-copy-to-clipboard'
import { Field, propTypes as formPropTypes } from 'redux-form'
import cx from 'classnames'

import config from 'config'
import { createEmptyPromise } from 'utils'
import { Button, Icon, Input, FieldText, Tooltip, toast } from 'uis'
import { getInvitePathByCode } from '../../index'

import './member-invite.view.styl'

const INVITE_LINK_VALID_MINUTES = 20

export class MemberInvite extends Component {

  static propTypes = {
    className: PropTypes.string,
    team: PropTypes.object,
    username: PropTypes.string,
    inviteCode: PropTypes.string,
    onUsernameChange: PropTypes.func,
    actions: PropTypes.object,
    ...formPropTypes
  }

  getRootClassnames () {
    return cx(
      'memberInviteView',
      this.props.className
    )
  }

  getInviteLink () {
    const { inviteCode } = this.props

    if (!inviteCode) {
      return ''
    }

    const invitePath = getInvitePathByCode(inviteCode)

    return `${config.APP_HOST}${invitePath}`
  }

  handleSubmit = (values) => {
    const {
      team,
      actions
    } = this.props

    const formPromise = createEmptyPromise()

    actions.createTeamMember({
      teamId: team.id,
      body: {
        userID: values.username
      },
      formPromise
    })

    return formPromise
  }

  renderMemberInviteForm () {
    const { onUsernameChange } = this.props
    const { handleSubmit, pristine, valid, submitting } = this.props

    return (
      <form onSubmit={handleSubmit(this.handleSubmit)}>
        <Field
          withRef
          ref={'usernameInput'}
          name={'username'}
          component={FieldText}
          placeholder={I18n.t('account.username')}
          onChange={onUsernameChange}
        />
        <Button
          block
          type={'primary'}
          htmlType={'submit'}
          disabled={pristine || !valid}
          loading={submitting}
        >
          <Translate value={'teamMembers.inviteGenerate'} />
        </Button>
      </form>
    )
  }

  getCopyHandler (value) {
    const handleCopy = () => {
      toast.success({
        message: I18n.t('teamMembers.inviteLinkCopySucceed')
      })
    }

    return (
      <Tooltip title={I18n.t('teamMembers.inviteLinkCopy')}>
        <CopyToClipboard text={value} onCopy={handleCopy}>
          <Icon name={'copy'} />
        </CopyToClipboard>
      </Tooltip>
    )
  }

  renderInviteLink () {
    const { username } = this.props

    const inviteLink = this.getInviteLink()
    const addonAfter = this.getCopyHandler(inviteLink)

    if (!inviteLink) {
      return null
    }

    return (
      <div className={'memberInviteLink'}>
        <div className={'memberInviteLinkDivider'} />
        <div className={'input-label'}>
          <Translate value={'teamMembers.inviteLink'} />
        </div>
        <Input
          disabled
          defaultValue={inviteLink}
          addonAfter={addonAfter}
        />
        <div className={'memberInviteLinkDescription'}>
          <Translate
            value={'teamMembers.inviteLinkDescription'}
            validMinutes={INVITE_LINK_VALID_MINUTES}
          />
          <div className={'username'}>{username}</div>
        </div>
      </div>
    )
  }

  render () {
    return (
      <div
        className={this.getRootClassnames()}
      >
        {this.renderMemberInviteForm()}
        {this.renderInviteLink()}
      </div>
    )
  }

}
