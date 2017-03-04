import { checkValidators, required } from 'modules'

export const validateEntryName = (...args) => {
  const validators = [required]
  return checkValidators(validators)(...args)
}

export const entryMakeValidate = (values, props) => {
  return {
    name: validateEntryName(values.name)
  }
}
