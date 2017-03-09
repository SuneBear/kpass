import {
  validateUsername,
  validatePassword,
  validatePasswordRetype
} from '../modules'

export const signUpValidate = (values, props) => {
  return {
    username: validateUsername(values.username),
    password: validatePassword(values.password),
    passwordRetype: validatePasswordRetype(
      values.passwordRetype,
      values.password,
    )
  }
}
