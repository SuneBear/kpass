import React, { Component, PropTypes } from 'react'
import cx from 'classnames'
import moment from 'moment'

import { minuteClock$ } from 'utils'

export class Datestamp extends Component {

  static propTypes = {
    className: PropTypes.string,
    date: PropTypes.instanceOf(Date)
  }

  constructor (props) {
    super(props)

    this.state = {
      date: this.format(props.date)
    }
  }

  componentWillReceiveProps (nextProps) {
    let cond = this.props.date.toJSON() !== nextProps.date.toJSON()
    if (cond) {
      this.setState({
        date: this.format(nextProps.date)
      })
    }
  }

  componentDidMount () {
    this.minuteClockSubscription = minuteClock$.subscribe(() => {
      this.setState({
        date: this.format(this.props.date)
      })
    })
  }

  componentWillUnmount () {
    this.minuteClockSubscription.unsubscribe()
  }

  getRootClassnames () {
    return cx(
      'datestampView',
      this.props.className
    )
  }

  format (t) {
    let now = moment()
    let date = moment(t)

    if (date.year() !== now.year()) {
      return date.format('L')
    } else {
      let diff = now.diff(date, 'hour', true)
      if (diff < 1) {
        return date.fromNow()
      } else {
        return date.calendar() + ' ' + date.format('LT')
      }
    }
  }

  render () {
    return (
      <time className={this.getRootClassnames()}>
        {this.state.date}
      </time>
    )
  }

}
