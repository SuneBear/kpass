import React from 'react'
// @REF: https://github.com/react-component/notification
import Notification from 'rc-notification'

import { capitalize } from 'utils/string'
import { ToastContent } from './toast.content.view'

import './toast.view.styl'

// Constants
const DEFAULF_DURATION = 5
const ANIMATION_DURATION = 268
const MAX_NOTICES = 5
const PLACEMENT = 'bottomLeft'

// Toast
// TODO: Optimize animation
const toastManager = Notification.newInstance({
  prefixCls: 'toast',
  className: `placement${capitalize(PLACEMENT)}`
})

// @FIXME: Notice DOM element is removed before animation completed
const moveNotice = () => {
  const notices = toastManager.component.state.notices
  if (notices.length > MAX_NOTICES) {
    window.setTimeout(() => {
      const oldestNoticeIndex = notices.length - MAX_NOTICES - 1
      const oldestNoticeKey = notices[oldestNoticeIndex].key
      toastManager.removeNotice(oldestNoticeKey)
    }, ANIMATION_DURATION)
  }
}

const showNotice = (options) => {
  const defaultOptions = {
    closable: true,
    duration: DEFAULF_DURATION,
    content: <ToastContent {...options} />
  }
  options = {
    ...defaultOptions,
    ...options
  }
  toastManager.notice(options)
  window.setTimeout(() => moveNotice(), 0)
}

// API
const api = {

  show (options) {
    showNotice(options)
  },

  remove (key) {
    if (toastManager) {
      toastManager.removeNotice(key)
    }
  },

  destroy () {
    if (toastManager) {
      toastManager.destroy()
    }
  }

}

const noticeStatus = ['success', 'info', 'warning', 'error']
noticeStatus.map((type) => {
  api[type] = (args) => api.show({
    ...args,
    type
  })
})

// Export
export const toast = api
