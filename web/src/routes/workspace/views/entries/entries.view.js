import React, { Component, PropTypes } from 'react'
import { I18n } from 'react-redux-i18n'
import cx from 'classnames'

import { Card } from 'views'

import './entries.view.styl'

export class Entries extends Component {

  static propTypes = {
    className: PropTypes.string
  }

  getRootClassNames () {
    return cx(
      'entriesView',
      this.props.className
    )
  }

  render () {
    return (
      <Card className={this.getRootClassNames()} title={I18n.t('entries.title')}>
        <div>Workspace Entries</div>
      </Card>
    )
  }

}
