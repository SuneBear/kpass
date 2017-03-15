import { checkValidators, required } from 'modules'

export const validateSecretName = (...args) => {
  const validators = [required]
  return checkValidators(validators)(...args)
}

export const secretMakeValidate = (values, props) => {
  return {
    name: validateSecretName(values.name)
  }
}
