const path = require('path')
const project = require('./project.config')

exports.assetsPath = function (_path) {
  return path.posix.join(
    project.compiler_public_path, // .replace(/^\//, '')
    _path
  )
}
