export const capitalize = (str) => {
  return str && str[0].toUpperCase() + str.slice(1)
}

export const asteriskify = (str, options = {}) => {
  let filteredStr

  const defaultOptions = {
    char: '*',
    keepSymbols: false
  }

  options = {
    ...defaultOptions,
    ...options
  }

  const regex = options.keepSymbols ? /[a-zA-Z0-9]/g : /(.)/g

  filteredStr = str.replace(regex, options.char)

  return filteredStr
}
