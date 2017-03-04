import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import { reduxForm } from 'redux-form'

import { signUpUserAction } from 'modules'
import { SignUp as SignUpView } from './sign-up.view'
import { signUpValidate } from './sign-up.validate'

const mapStateToProps = (state) => ({})

const mapDispatchToProps = (dispatch) => ({
  actions: bindActionCreators({
    signUpUser: signUpUserAction
  }, dispatch)
})

const createContainer = (component) => {
  const connectedComponent = connect(
    mapStateToProps,
    mapDispatchToProps
  )(component)

  return reduxForm({
    form: 'signUpForm',
    validate: signUpValidate
  })(connectedComponent)
}

export const SignUp = createContainer(SignUpView)
