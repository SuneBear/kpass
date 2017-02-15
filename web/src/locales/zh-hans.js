import { mixinDeep } from 'utils'
import { enUS } from './en-us'

const _zhHans = {
  account: {
    password: '请输入密码'
  },

  zPlaceholder: '你好，世界'
}

export const zhHans = mixinDeep({}, enUS, _zhHans)
