import 'rxjs/add/observable/interval'
import 'rxjs/add/operator/share'

import { Observable } from 'rxjs/Observable'
import moment from 'moment'

export const minuteClock$ = Observable.interval(60 * 1000).share()

moment.defineLocale('zh-Hans', {
  months: '一月_二月_三月_四月_五月_六月_七月_八月_九月_十月_十一月_十二月'.split('_'),
  monthsShort: '1月_2月_3月_4月_5月_6月_7月_8月_9月_10月_11月_12月'.split('_'),
  weekdays: '周日_周一_周二_周三_周四_周五_周六'.split('_'),
  weekdaysShort: '周日_周一_周二_周三_周四_周五_周六'.split('_'),
  weekdaysMin: '日_一_二_三_四_五_六'.split('_'),
  longDateFormat: {
    LT: 'H:mm',
    L: 'YYYY年MMMD日',
    LL: 'YYYY年MMMD日',
    LLL: 'YYYY年MMMD日LT',
    LLLL: 'YYYY年MMMD日ddddLT',
    l: 'YYYY年MMMD日',
    ll: 'YYYY年MMMD日',
    lll: 'YYYY年MMMD日 LT',
    llll: 'YYYY年MMMD日ddddLT'
  },
  relativeTime: {
    future: '%s内',
    past: '%s前',
    s: '几秒',
    m: '1分钟',
    mm: '%d分钟',
    h: '1小时',
    hh: '%d小时',
    d: '1天',
    dd: '%d天',
    M: '1个月',
    MM: '%d个月',
    y: '1年',
    yy: '%d年'
  },
  week: {
    dow: 1 // Monday is the first day of the week
  },
  calendar: {
    sameDay: '[今天]',
    nextDay: '[明天]',
    lastDay: '[昨天]',
    lastWeek: function () {
      var prefix, startOfWeek
      startOfWeek = moment().startOf('week')
      prefix = this.unix() < startOfWeek.unix() ? '[上]' : ''
      return prefix + 'dddd'
    },
    nextWeek: function () {
      var prefix, startOfWeek
      startOfWeek = moment().startOf('week')
      prefix = this.unix() - startOfWeek.unix() >= 7 * 24 * 3600 ? '[下]' : ''
      return prefix + 'dddd'
    },
    sameElse: function () {
      if (this.year() === moment().year()) {
        return 'MMMD日'
      } else {
        return 'LL'
      }
    }
  },
  meridiem: function (hour, minute, isLower) {
    let hm = hour * 100 + minute
    if (hm < 500) {
      return '凌晨'
    } else if (hm < 900) {
      return '早上'
    } else if (hm < 1130) {
      return '上午'
    } else if (hm < 1230) {
      return '中午'
    } else if (hm < 1800) {
      return '下午'
    } else {
      return '晚上'
    }
  },
  ordinal: function (number, period) {
    switch (period) {
      case 'd':
      case 'D':
      case 'DDD':
        return number + '日'
      case 'M':
        return number + '月'
      case 'w':
      case 'W':
        return number + '周'
      default:
        return number
    }
  }
})

moment.defineLocale('en-US', {
  longDateFormat: {
    LT: 'H:mm',
    L: 'YYYY/MM/DD',
    LL: 'MMM DD, YYYY',
    LLL: 'LL LT',
    LLLL: 'dddd, LLL'
  },
  calendar: {
    lastDay: '[Yesterday]',
    sameDay: '[Today]',
    nextDay: '[Tomorrow]',
    lastWeek: function () {
      var prefix, startOfWeek
      startOfWeek = moment().startOf('week')
      prefix = this.unix() < startOfWeek.unix() ? '[Last] ' : ''
      return prefix + 'dddd'
    },
    nextWeek: function () {
      var prefix, startOfWeek
      startOfWeek = moment().startOf('week')
      prefix = this.unix() - startOfWeek.unix() >= 7 * 24 * 3600 ? '[Next] ' : ''
      return prefix + 'dddd'
    },
    sameElse: function () {
      if (this.year() === moment().year()) {
        return 'MMM D'
      } else {
        return 'LL'
      }
    }
  }
})
