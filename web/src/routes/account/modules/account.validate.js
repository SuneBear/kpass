import {
  checkValidators,
  required,
  minLength,
  maxLength,
  retype
} from 'modules'

export const validateUsername = (...args) => {
  const validators = [required, minLength(3), maxLength(20)]
  return checkValidators(validators)(...args)
}

export const validatePassword = (...args) => {
  const validators = [required, minLength(3), maxLength(20)]
  return checkValidators(validators)(...args)
}

export const validatePasswordRetype = (...args) => {
  const validators = [required, retype]
  return checkValidators(validators)(...args)
}
