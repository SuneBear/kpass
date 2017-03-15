import React, { Component, PropTypes } from 'react'
import { I18n, Translate } from 'react-redux-i18n'
import { Field, propTypes as formPropTypes } from 'redux-form'
import cx from 'classnames'

import { createEmptyPromise } from 'utils'
import { Icon, Button, FieldText } from 'uis'

import './secret-make.view.styl'

export class SecretMake extends Component {

  static propTypes = {
    className: PropTypes.string,
    entry: PropTypes.object,
    secret: PropTypes.object,
    action: PropTypes.oneOf(['create', 'update']),
    actions: PropTypes.object,
    ...formPropTypes
  }

  static defaultProps = {
    action: 'create'
  }

  getRootClassnames () {
    return cx(
      'secretMakeView',
      this.props.className
    )
  }

  isCreate () {
    const { action } = this.props

    return action === 'create'
  }

  handleSubmit = (values) => {
    const {
      entry,
      secret,
      actions
    } = this.props

    const formPromise = createEmptyPromise()

    if (this.isCreate()) {
      actions.createSecret({
        entry,
        entryId: entry.id,
        body: values,
        formPromise
      })
    } else {
      actions.updateSecret({
        entryId: entry.id,
        secretId: secret.id,
        body: values,
        formPromise
      })
    }

    return formPromise
  }

  renderDivider () {
    return (
      <div className={'secretMakeDivider'} />
    )
  }

  renderHeader () {
    const {
      entry,
      secret
    } = this.props

    const title = this.isCreate()
      ? I18n.t('secret.newTitle', { entryName: entry.name })
      : I18n.t('secret.editTitle', { secretName: secret.name })

    return (
      <div className={'secretMakeHeader'}>
        <div className={'secretMakeHeaderTitle'}>
          {title}
        </div>
        {this.renderDivider()}
      </div>
    )
  }

  renderSecretMakeForm () {
    const { handleSubmit, pristine, valid, submitting } = this.props

    return (
      <form onSubmit={handleSubmit(this.handleSubmit)}>
        <Field
          name={'name'}
          component={FieldText}
          placeholder={I18n.t('secret.secretNamePlaceholder')}
        />
        <div className={'fieldGroup'}>
          <Field
            name={'note'}
            component={FieldText}
            type={'textarea'}
            autosize={{ minRows: 3 }}
            placeholder={I18n.t('secret.secretNotePlaceholder')}
          />
          <Field
            name={'password'}
            component={FieldText}
            prefix={<Icon name={'lock'} />}
            placeholder={I18n.t('secret.secretPasswordPlaceholder')}
          />
          <Field
            name={'url'}
            component={FieldText}
            prefix={<Icon name={'link'} />}
            placeholder={I18n.t('secret.secretUrlPlaceholder')}
          />
        </div>
        {this.renderDivider()}
        <Button
          block
          type={'primary'}
          htmlType={'submit'}
          disabled={pristine || !valid}
          loading={submitting}
        >
          <Translate value={this.isCreate() ? 'action.add' : 'action.edit'} />
        </Button>
      </form>
    )
  }

  render () {
    return (
      <div
        className={this.getRootClassnames()}
      >
        {this.renderHeader()}
        <div className={'secretMakeContent'}>
          {this.renderSecretMakeForm()}
        </div>
      </div>
    )
  }

}
