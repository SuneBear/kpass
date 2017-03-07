import React, { Component, PropTypes } from 'react'
import { I18n, Translate } from 'react-redux-i18n'
import { Field, propTypes as formPropTypes } from 'redux-form'
import cx from 'classnames'

import { Button, FieldText, FieldSelect } from 'uis'
import { createEmptyPromise } from 'utils'

import { getEntryCategories } from '../../shared'

import './entry-make.view.styl'

export class EntryMake extends Component {

  static propTypes = {
    className: PropTypes.string,
    team: PropTypes.object,
    entryId: PropTypes.string,
    action: PropTypes.oneOf(['create', 'update']),
    actions: PropTypes.object,
    ...formPropTypes
  }

  static defaultProps = {
    action: 'create'
  }

  getRootClassnames () {
    return cx(
      'entryMakeView',
      this.props.className
    )
  }

  handleSubmit = (values) => {
    const {
      team,
      entryId,
      action,
      actions
    } = this.props

    const isCreate = action === 'create'
    const formPromise = createEmptyPromise()

    if (isCreate) {
      actions.createEntry({
        teamId: team.id,
        body: values,
        formPromise
      })
    } else {
      actions.updateEntry({
        body: values,
        entryId,
        formPromise
      })
    }

    return formPromise
  }

  renderEntryMakeForm () {
    const { handleSubmit, pristine, valid, submitting } = this.props
    const { action } = this.props

    const isCreate = action === 'create'

    return (
      <form onSubmit={handleSubmit(this.handleSubmit)}>
        <Field
          name={'name'}
          component={FieldText}
          placeholder={I18n.t('entry.entryName')}
        />
        <Field
          name={'category'}
          component={FieldSelect}
          label={I18n.t('entry.entryCategory')}
          options={getEntryCategories()}
        />
        <Button
          block
          type={'primary'}
          htmlType={'submit'}
          disabled={pristine || !valid}
          loading={submitting}
        >
          <Translate value={isCreate ? 'action.create' : 'action.edit'} />
        </Button>
      </form>
    )
  }

  render () {
    return (
      <div
        className={this.getRootClassnames()}
      >
        {this.renderEntryMakeForm()}
      </div>
    )
  }

}
