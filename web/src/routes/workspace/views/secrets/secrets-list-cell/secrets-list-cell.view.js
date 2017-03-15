import React, { Component, PropTypes } from 'react'
import { I18n, Translate } from 'react-redux-i18n'
import CopyToClipboard from 'react-copy-to-clipboard'
import cx from 'classnames'
import { pick } from 'lodash'

import { isUrl, asteriskify } from 'utils'
import { Icon, Dropdown, MenuSelector, Modal, Tooltip, toast } from 'uis'
import { Card, Readable } from 'views'
import { SecretMake } from '../../secret-make'

import './secrets-list-cell.view.styl'

export class SecretsListCell extends Component {

  static propTypes = {
    className: PropTypes.string,
    entry: PropTypes.object,
    secret: PropTypes.object,
    creatorPermissions: PropTypes.object,
    actions: PropTypes.object
  }

  constructor (props) {
    super(props)

    this.state = {
      showPassword: false
    }
  }

  getRootClassnames () {
    return cx(
      'secretsListCellView',
      this.props.className
    )
  }

  saveEditSecretModalRef = (ref) => {
    this.editSecretModalRef = ref
  }

  handleEditSecretClick = () => {
    this.editSecretModalRef.open()
  }

  handleEditSecretSubmitSuccess = () => {
    this.editSecretModalRef.close()
  }

  handleToggleShowPassword = () => {
    this.state.showPassword = !this.state.showPassword
    this.setState(this.state)
  }

  handleDeleteSecret = () => {
    const { entry, secret, actions } = this.props

    actions.deleteSecret({
      entry,
      entryId: entry.id,
      secretId: secret.id
    })
  }

  hasInfos () {
    const { secret } = this.props

    const conditions = [
      !!secret.password,
      !!secret.url
    ]

    return conditions.some((cond) => cond)
  }

  renderMore () {
    const { creatorPermissions } = this.props

    const conditions = [
      creatorPermissions.updateSecret,
      creatorPermissions.deleteSecret
    ]

    if (!conditions.some((cond) => cond)) {
      return null
    }

    return (
      <Dropdown
        className={'secretsListCellMoreDropdown'}
        content={this.getMoreMenu()}
      >
        <Tooltip
          title={I18n.t('secret.menu')}
        >
          <div className={'secretsListCellMoreHandler'}>
            <Icon name={'chevron-down'} Component={'a'} />
          </div>
        </Tooltip>
      </Dropdown>
    )
  }

  getMoreMenu () {
    const { creatorPermissions } = this.props

    const dataList = []

    if (creatorPermissions.updateSecret) {
      dataList.push({
        iconName: 'pencil',
        title: I18n.t('secret.edit'),
        onClick: this.handleEditSecretClick
      })
    }

    if (creatorPermissions.deleteSecret) {
      dataList.push({
        iconName: 'trash',
        title: I18n.t('secret.delete'),
        onClick: this.handleDeleteSecret
      })
    }

    return (
      <MenuSelector
        dataList={dataList}
      />
    )
  }

  renderDetailDivider () {
    return (
      <div className={'secretDetailDivider'} />
    )
  }

  renderDetailHeader () {
    const { secret } = this.props

    return (
      <div className={'secretDetailHeader'}>
        <div className={'secretDetailHeaderName'}>
          {secret.name}
        </div>
        {this.renderMore()}
      </div>
    )
  }

  renderDetailInfos () {
    if (!this.hasInfos()) {
      return null
    }

    return (
      <div className={'secretDetailInfos'}>
        {this.renderDetailInfoPassword()}
        {this.renderDetailInfoUrl()}
      </div>
    )
  }

  renderDetailInfoPassword () {
    const { secret } = this.props
    const { showPassword } = this.state
    let formattedPassword = secret.password

    if (!secret.password) {
      return null
    }

    if (!showPassword) {
      formattedPassword = asteriskify(formattedPassword)
    }

    const handleCopy = () => {
      toast.success({
        message: I18n.t('secret.secretPasswordCopySucceed')
      })
    }

    const ValueClassName = cx(
      'secretDetailInfosItemValue',
      { isConcealed: !showPassword }
    )

    const togglePasswordClassName = cx(
      { isActive: showPassword }
    )

    return (
      <div className={'secretDetailInfosItem'}>
        <div className={'secretDetailInfosItemLabel'}>
          <Icon name={'lock'} />
          <Translate value={'secret.secretPassword'} />
        </div>
        <div className={ValueClassName}>
          {formattedPassword}
        </div>
        <div className={'secretDetailInfosItemHandler'}>
          <Tooltip title={I18n.t(showPassword
            ? 'secret.secretPasswordHide'
            : 'secret.secretPasswordShow')}
          >
            <Icon
              name={'eye'}
              className={togglePasswordClassName}
              Component={'a'}
              onClick={this.handleToggleShowPassword}
            />
          </Tooltip>
          <Tooltip title={I18n.t('secret.secretPasswordCopy')}>
            <CopyToClipboard text={secret.password} onCopy={handleCopy}>
              <Icon name={'copy'} Component={'a'} />
            </CopyToClipboard>
          </Tooltip>
        </div>
      </div>
    )
  }

  renderDetailInfoUrl () {
    const { secret } = this.props
    const PROTOCOL = 'http://'
    let url = secret.url

    if (!url) {
      return null
    }

    if (!isUrl(url)) {
      url = `${PROTOCOL}${url}`
    }

    const handleOpen = () => {
      const nextWindow = window.open(url, '_blank')
      nextWindow.focus()
    }

    return (
      <div className={'secretDetailInfosItem'}>
        <div className={'secretDetailInfosItemLabel'}>
          <Icon name={'link'} />
          <Translate value={'secret.secretUrl'} />
        </div>
        <div className={'secretDetailInfosItemValue'}>
          {secret.url}
        </div>
        <div className={'secretDetailInfosItemHandler'} onClick={handleOpen}>
          <Tooltip title={I18n.t('secret.secretUrlGoTo')}>
            <Icon name={'share-stroke'} Component={'a'} />
          </Tooltip>
        </div>
      </div>
    )
  }

  renderDetailNote () {
    const { secret } = this.props

    if (!secret.note) {
      return null
    }

    return (
      <div className={'secretDetailNote'}>
        {this.renderDetailDivider()}
        <Readable
          type={'markdown'}
          content={secret.note}
        />
      </div>
    )
  }

  renderDetail () {
    return (
      <div className={'secretDetail'}>
        {this.renderDetailHeader()}
        {this.renderDetailNote()}
        {this.renderDetailInfos()}
      </div>
    )
  }

  renderEditSecretModal () {
    const { secret } = this.props

    const initialValues = pick(secret, [
      'name', 'password', 'url', 'note'
    ])

    return (
      <Modal
        ref={this.saveEditSecretModalRef}
        className={'secretMakeModal'}
      >
        <SecretMake
          action={'update'}
          secret={secret}
          initialValues={initialValues}
          onSubmitSuccess={this.handleEditSecretSubmitSuccess}
        />
      </Modal>
    )
  }

  render () {
    return (
      <Card
        className={this.getRootClassnames()}
      >
        {this.renderDetail()}

        {this.renderEditSecretModal()}
      </Card>
    )
  }

}
