import React, { Component, PropTypes } from 'react'
import { I18n, Translate } from 'react-redux-i18n'
import { Field, propTypes as formPropTypes } from 'redux-form'
import cx from 'classnames'

import { Button, FieldText } from 'uis'
import { createEmptyPromise } from 'utils'

import './team-rename.view.styl'

export class TeamRename extends Component {

  static propTypes = {
    className: PropTypes.string,
    team: PropTypes.object,
    actions: PropTypes.object,
    ...formPropTypes
  }

  getRootClassnames () {
    return cx(
      'teamRenameView',
      this.props.className
    )
  }

  handleSubmit = (values) => {
    const {
      team,
      actions
    } = this.props

    const {
      name
    } = values

    const formPromise = createEmptyPromise()

    actions.updateTeam({
      teamId: team.id,
      body: { name },
      formPromise
    })

    return formPromise
  }

  renderTeamRenameForm () {
    const {
      handleSubmit,
      pristine,
      valid,
      submitting
    } = this.props

    return (
      <form
        onSubmit={handleSubmit(this.handleSubmit)}
      >
        <Field
          name={'name'}
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
          <Translate value={'action.save'} />
        </Button>
      </form>
    )
  }

  render () {
    return (
      <div className={this.getRootClassnames()} >
        {this.renderTeamRenameForm()}
      </div>
    )
  }

}
