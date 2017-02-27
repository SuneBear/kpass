import { checkValidators, required } from 'modules'

export const validateTeamName = (...args) => {
  const validators = [required]
  return checkValidators(validators)(...args)
}

export const TeamCreateValidate = (values, props) => {
  return {
    teamName: validateTeamName(values.teamName)
  }
}
