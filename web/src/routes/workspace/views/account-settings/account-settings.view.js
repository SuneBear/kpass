import React, { Component, PropTypes } from 'react'
import { I18n, Translate } from 'react-redux-i18n'
import cx from 'classnames'

import { Button, Input, Upload, toast } from 'uis'
import { getFileUrl } from 'utils'
import { Avatar } from 'views'

import './account-settings.view.styl'

export class AccountSettings extends Component {

  static propTypes = {
    className: PropTypes.string,
    userMe: PropTypes.object,
    actions: PropTypes.object
  }

  getRootClassnames () {
    return cx(
      'accountSettingsView',
      this.props.className
    )
  }

  handleUploadAvatarChange = (upload) => {
    const { updateUser } = this.props.actions
    const status = upload.file.status
    const body = upload.file.response

    if (status === 'error') {
      toast.error({
        message: I18n.t('accountSettings.uploadAvatarFailed')
      })
    }

    if (status === 'done') {
      updateUser({
        body
      })
      toast.success({
        message: I18n.t('accountSettings.uploadAvatarSucceed')
      })
    }
  }

  renderSettingAvatar () {
    const { userMe } = this.props

    return (
      <div className={'settingsSectionItem'}>
        <div className={'settingsSectionItemLabel'}>
          <Translate value={'member.avatar'} />
        </div>
        <div className={'settingsSectionItemContent'}>
          <Avatar url={getFileUrl(userMe.avatar)} size={'large'} />
          <Upload
            url={'avatar'}
            accept={'image/*'}
            onChange={this.handleUploadAvatarChange}
          >
            <Button type={'text'}>
              <Translate value={'accountSettings.uploadAvatar'} />
            </Button>
          </Upload>
        </div>
      </div>
    )
  }

  renderSettingName () {
    const { userMe } = this.props

    return (
      <div className={'settingsSectionItem'}>
        <div className={'settingsSectionItemLabel'}>
          <Translate value={'member.name'} />
        </div>
        <div className={'settingsSectionItemContent'}>
          <Input value={userMe.id} disabled />
        </div>
      </div>
    )
  }

  renderPublicProfileSection () {
    return (
      <div className={'settingsSection'}>
        <div className={'settingsSectionTitle'}>
          <Translate value={'accountSettings.publicProfile'} />
        </div>
        <div className={'settingsSectionContent'}>
          {this.renderSettingAvatar()}
          {this.renderSettingName()}
        </div>
      </div>
    )
  }

  render () {
    return (
      <div className={this.getRootClassnames()}>
        {this.renderPublicProfileSection()}
      </div>
    )
  }

}
