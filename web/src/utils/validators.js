const URL_REGEXP = /^(?:\w+:)?\/\/([^\s.]+\.\S{2}?)\S*$/

export const isObject = (item) => {
  return (item && typeof item === 'object' && !Array.isArray(item) && item !== null)
}

export const isUrl = (str) => URL_REGEXP.test(str)
