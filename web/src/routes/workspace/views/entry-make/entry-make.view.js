import React, { Component, PropTypes } from 'react'
import { I18n, Translate } from 'react-redux-i18n'
import { Field, propTypes as formPropTypes } from 'redux-form'
import cx from 'classnames'

import { Button, FieldText, FieldSelect } from 'uis'

import { getEntryCategoryOptions } from '../../shared'

import './entry-make.view.styl'

export class EntryMake extends Component {

  static propTypes = {
    className: PropTypes.string,
    team: PropTypes.object,
    entry: PropTypes.object,
    currentAction: PropTypes.oneOf(['create', 'update']),
    actions: PropTypes.object,
    ...formPropTypes
  }

  static defaultProps = {
    currentAction: 'create'
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
      currentAction,
      actions
    } = this.props

    if (currentAction === 'create') {
      // @TODO: Implementation
      actions.createEntry({
        teamId: team.id,
        ...values
      })
    } else {
      // @TODO: Implementation
      actions.updateEntry({
        teamId: team.id,
        ...values
      })
    }
  }

  renderEntryMakeForm () {
    const { handleSubmit, pristine, valid, submitting } = this.props
    const { entry } = this.props

    return (
      <form onSubmit={handleSubmit(this.handleSubmit)}>
        <Field
          name={'entryName'}
          component={FieldText}
          placeholder={I18n.t('entry.entryName')}
        />
        <Field
          name={'entryCategory'}
          component={FieldSelect}
          label={I18n.t('entry.entryCategory')}
          defaultValue={entry.category || 'Login'}
          options={getEntryCategoryOptions()}
        />
        <Button
          block
          type={'primary'}
          htmlType={'submit'}
          disabled={pristine || !valid}
          loading={submitting}
        >
          <Translate value={'action.create'} />
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
