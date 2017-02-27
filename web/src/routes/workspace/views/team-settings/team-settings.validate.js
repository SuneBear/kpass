import { validateTeamName } from '../team-create/team-create.validate'

export const teamSettingsValidate = (values, props) => {
  return {
    teamName: validateTeamName(values.teamName)
  }
}
