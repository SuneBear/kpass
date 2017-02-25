import React, { Component, PropTypes } from 'react'
import { Link } from 'react-router'
import { I18n, Translate } from 'react-redux-i18n'
import { Field, propTypes as formPropTypes } from 'redux-form'
import cx from 'classnames'

import { Button, FieldText } from 'uis'

import './sign-in.view.styl'

export class SignIn extends Component {

  static propTypes = {
    className: PropTypes.string,
    actions: PropTypes.object,
    ...formPropTypes
  }

  getRootClassNames () {
    return cx(
      'signInView',
      this.props.className
    )
  }

  handleSubmit = (values) => {
    const {
      actions
    } = this.props

    actions.signInUser({
      username: values.username,
      password: values.password
    })

  }

  renderSignInForm () {
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
          className={'SignInHandler'}
          type={'primary'}
          htmlType={'submit'}
          disabled={!anyTouched || !valid}
          loading={submitting}
          block
        >
          <Translate value={'account.signIn'} />
        </Button>
      </form>
    )
  }

  render () {
    return (
      <div className={this.getRootClassNames()}>
        <div className={'accountLayoutForm'}>
          {this.renderSignInForm()}
        </div>
        <Link className={'accountLayoutFooter'} to={'/account/sign-up'}>
          <Translate value={'account.signUpTip'} />
        </Link>
      </div>
    )
  }

}
