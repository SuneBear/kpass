import { validateTeamName } from '../team-create/team-create.validate'

export const teamRenameValidate = (values, props) => {
  return {
    name: validateTeamName(values.name)
  }
}
