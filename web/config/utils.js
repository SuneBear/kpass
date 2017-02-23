const project = require('./project.config')

exports.assetsPath = function (_path) {
  return project.compiler_public_path + _path
}
