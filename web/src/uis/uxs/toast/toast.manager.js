import React, { Component, PropTypes } from 'react'
// @REF: https://github.com/react-component/notification
import Notification from 'rc-notification'

import { capitalize } from 'utils/string'
import { ToastContent } from './toast.content.view'

import './toast.view.styl'

// Constants
const DEFAULF_DURATION = 5
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
    setTimeout(() => {
      const oldestNoticeIndex = notices.length - MAX_NOTICES - 1
      const oldestNoticeKey = notices[oldestNoticeIndex].key
      toastManager.removeNotice(oldestNoticeKey)
    }, 168)
  }
}

const showNotice = (options) => {
  const defaultOptions = {
    closable: true,
    duration: DEFAULF_DURATION,
    content: <ToastContent {...options} />
  }
  options = Object.assign({}, defaultOptions, options)
  toastManager.notice(options)
  setTimeout(() => moveNotice(), 0)
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
  api[type] = (args) => api.show(Object.assign({}, args, { type }))
})

// Export
export const toast = api
