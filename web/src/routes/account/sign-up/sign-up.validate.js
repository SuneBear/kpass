import { validateUsername, validatePassword } from '../modules'

export const signUpValidate = (values, props) => {
  return {
    username: validateUsername(values.username),
    password: validatePassword(values.password)
  }
}
