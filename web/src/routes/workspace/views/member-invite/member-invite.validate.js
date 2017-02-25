import { validateUsername } from '../../../account/modules'

export const memberInviteValidate = (values, props) => {
  return {
    username: validateUsername(values.username)
  }
}
