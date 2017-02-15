import { validateUsername, validatePassword } from '../modules'

export const signInValidate = (values, props) => {
  return {
    username: validateUsername(values.username),
    password: validatePassword(values.password)
  }
}
