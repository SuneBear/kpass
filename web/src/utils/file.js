import config from 'config'

export const getFileUrl = (baseUrl) => {
  if (!baseUrl) {
    return null
  }
  const downloadHost = `${config.FILE_HOST}`
  const fileUrl = `${downloadHost}${baseUrl}`
  return fileUrl
}
