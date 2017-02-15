import defaultConfig from './default'
import productionConfig from './production'

const config = __DEV__
  ? defaultConfig
  : Object.assign({}, defaultConfig, productionConfig)

export default config
