// Util
export const checkValidators = validators => (...args) => {
  for (const validator of validators) {
    const result = validator(...args)
    if (result) {
      return result
    }
  }
  return undefined
}

// Common Validators
export const required = value => value ? undefined : 'Required'

export const number = value => value && isNaN(Number(value)) ? 'Must be a number' : undefined

export const minLength = min => value =>
  value && value.length < min ? `Must be at least ${min}` : undefined

export const maxLength = max => value =>
  value && value.length > max ? `Must be ${max} characters or less` : undefined

export const retype = (retypeValue, value) =>
  retypeValue !== value ? 'Incorrect retyped value' : undefined
