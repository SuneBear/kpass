import React, { Component, PropTypes } from 'react'
import { Link } from 'react-router'
import { I18n, Translate } from 'react-redux-i18n'
import { Field } from 'redux-form'
import cx from 'classnames'

import { Button, FieldText } from 'uis'

import './sign-up.view.styl'

const sleep = ms => new Promise(resolve => setTimeout(resolve, ms))

export class SignUp extends Component {

  static propTypes = {
    className: PropTypes.string,
    actions: PropTypes.object
  }

  getRootClassNames () {
    return cx(
      'signUpView',
      this.props.className
    )
  }

  handleSubmit = (values) => {
    const {
      actions
    } = this.props

    actions.signUpUser({
      username: values.username,
      password: values.password
    })

    return sleep(1000)
      .then(() => {})
  }

  renderSignUpForm () {
    const { handleSubmit, anyTouched, valid, submitting } = this.props

    return (
      <form onSubmit={handleSubmit(this.handleSubmit)}>
        <Field
          name='username'
          component={FieldText}
          placeholder={I18n.t('account.username')}
          autoFocus
        />
        <Field
          name='password'
          component={FieldText}
          type={'password'}
          placeholder={I18n.t('account.password')}
        />
        <Button
          className={'SignUpHandler'}
          type={'primary'}
          htmlType={'submit'}
          disabled={!anyTouched || !valid}
          loading={submitting}
          block
        >
          <Translate value={'account.signUp'} />
        </Button>
      </form>
    )
  }

  render () {
    return (
      <div className={this.getRootClassNames()}>
        <div className={'accountLayoutForm'}>
          {this.renderSignUpForm()}
        </div>
        <Link className={'accountLayoutFooter'} to={'/account/sign-in'}>
          <Translate value={'account.signInTip'} />
        </Link>
      </div>
    )
  }

}
