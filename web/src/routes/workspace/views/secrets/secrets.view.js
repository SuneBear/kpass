import React, { Component, PropTypes } from 'react'
import { I18n } from 'react-redux-i18n'
import cx from 'classnames'

import { Card, Placeholder } from 'views'
import { SecretsList } from './secrets-list'

import './secrets.view.styl'

export class Secrets extends Component {

  static propTypes = {
    className: PropTypes.string,
    userMe: PropTypes.object,
    team: PropTypes.object,
    entry: PropTypes.object,
    secrets: PropTypes.array,
    actions: PropTypes.object
  }

  getRootClassNames () {
    return cx(
      'secretsView',
      this.props.className
    )
  }

  renderPlaceholder () {
    const { secrets } = this.props

    if (secrets && secrets.length !== 0) {
      return null
    }

    return (
      <Card>
        <Placeholder
          size={'small'}
          iconName={'lock'}
          title={I18n.t('secrets.placeholderTitle')}
         />
      </Card>
    )
  }

  renderSecrets () {
    const {
      userMe,
      team,
      entry,
      secrets,
      actions
    } = this.props

    if (!secrets || !secrets.length) {
      return null
    }

    return (
      <SecretsList
        userMe={userMe}
        team={team}
        entry={entry}
        secrets={secrets}
        actions={actions}
      />
    )
  }

  render () {
    return (
      <div
        className={this.getRootClassNames()}
      >
        {this.renderPlaceholder()}
        {this.renderSecrets()}
      </div>
    )
  }

}
