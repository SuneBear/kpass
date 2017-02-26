export const cookie = (name, value, days) => {
  // Get: if value is undefined, get the cookie value
  if (value === undefined) {
    const cookiestring = '; ' + document.cookie
    const cookies = cookiestring.split('; ' + name + '=')
    if (cookies.length === 2) {
      return decodeURIComponent(
        cookies.pop()
          .split(';').shift()
      )
    }
    return null
  } else {
    // Delete: if value is a false boolean or null, we'll treat that as a delete
    if (value === false || value === null) {
      days = -1
    } else {
      // Escape value
      value = encodeURIComponent(value)
    }
    // Set
    let expires
    if (days) {
      const date = new Date()
      date.setTime(date.getTime() + (days * 24 * 60 * 60 * 1000))
      expires = '; expires=' + date.toGMTString()
    } else {
      expires = ''
    }
    document.cookie = name + '=' + value + expires + '; path=/'
    return value
  }
}
