import React, { Component, PropTypes } from 'react'
import { I18n, Translate } from 'react-redux-i18n'
import { Field, propTypes as formPropTypes } from 'redux-form'
import cx from 'classnames'

import { createEmptyPromise } from 'utils'
import { Button, FieldText } from 'uis'

import './team-create.view.styl'

export class TeamCreate extends Component {

  static propTypes = {
    className: PropTypes.string,
    actions: PropTypes.object,
    ...formPropTypes
  }

  getRootClassnames () {
    return cx(
      'teamCreateView',
      this.props.className
    )
  }

  handleSubmit = (values) => {
    const {
      actions
    } = this.props

    const formPromise = createEmptyPromise()

    // @TODO: Implementation
    actions.createTeam({
      body: {
        name: values.teamName
      },
      formPromise
    })

    return formPromise
  }

  componentDidMount () {
    this.refs.teamNameInput
      .getRenderedComponent().focus()
  }

  renderMemberInviteForm () {
    const { handleSubmit, pristine, valid, submitting } = this.props

    return (
      <form onSubmit={handleSubmit(this.handleSubmit)}>
        <Field
          withRef
          ref={'teamNameInput'}
          name={'teamName'}
          component={FieldText}
          placeholder={I18n.t('team.teamName')}
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
        {this.renderMemberInviteForm()}
      </div>
    )
  }

}
