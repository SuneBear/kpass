export const getEntryCategories = () => ([
  {
    name: 'Login',
    color: '#F58E3D'
  },

  {
    name: 'Network',
    color: '#797EC9'
  },

  {
    name: 'Software License',
    color: '#75C940'
  },

  {
    name: 'Secure Note',
    color: '#FFE738'
  },

  {
    name: 'Server',
    color: '#FF4F3E'
  }
])

export const getEntryCategoryNames = () => {
  return getEntryCategories().map(
    category => category.name
  )
}

export const getEntryCategoryOptions = () => {
  return getEntryCategories().map(
    category => ({
      value: category.name,
      color: category.color
    })
  )
}

export const getEntryCategoryByName = (name) => {
  return getEntryCategories().find(
    category => category.name === name
  )
}
