import React, { Component, PropTypes } from 'react'
import { Link } from 'react-router'
import { I18n, Translate } from 'react-redux-i18n'
import { Field, propTypes as formPropTypes } from 'redux-form'
import cx from 'classnames'

import { createEmptyPromise } from 'utils'
import { Button, FieldText } from 'uis'
import { ACCOUNT_BASE_PATH, SIGN_UP_PATH } from '../index'

import './sign-in.view.styl'

export class SignIn extends Component {

  static propTypes = {
    className: PropTypes.string,
    location: PropTypes.object,
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
      location,
      actions
    } = this.props

    const formPromise = createEmptyPromise()

    actions.signInUser({
      username: values.username,
      password: values.password,
      location,
      formPromise
    })

    return formPromise
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
    const { location } = this.props

    return (
      <div className={this.getRootClassNames()}>
        <div className={'accountLayoutForm'}>
          {this.renderSignInForm()}
        </div>
        <Link
          className={'accountLayoutFooter'}
          to={`${ACCOUNT_BASE_PATH}/${SIGN_UP_PATH}${location.search}`}
        >
          <Translate value={'account.signUpTip'} />
        </Link>
      </div>
    )
  }

}
