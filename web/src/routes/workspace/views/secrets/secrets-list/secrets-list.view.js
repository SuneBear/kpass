import React, { Component, PropTypes } from 'react'
import cx from 'classnames'

import { getCreatorPermissions } from 'utils'
import { SecretsListCell } from '../secrets-list-cell'

import './secrets-list.view.styl'

export class SecretsList extends Component {

  static propTypes = {
    className: PropTypes.string,
    userMe: PropTypes.object,
    team: PropTypes.object,
    entry: PropTypes.object,
    secrets: PropTypes.array,
    actions: PropTypes.object
  }

  getRootClassnames () {
    return cx(
      'secretsListView',
      this.props.className
    )
  }

  renderCell (secret) {
    const {
      userMe,
      team,
      entry,
      actions
    } = this.props

    const creatorPermissions = getCreatorPermissions(
      secret.userID,
      userMe,
      team
    )

    return (
      <SecretsListCell
        key={secret.id}
        entry={entry}
        secret={secret}
        creatorPermissions={creatorPermissions}
        actions={actions}
       />
    )
  }

  renderCells () {
    const {
      secrets
    } = this.props

    return secrets.map((secret) => {
      return this.renderCell(secret)
    })
  }

  render () {
    return (
      <div className={this.getRootClassnames()}>
        {this.renderCells()}
      </div>
    )
  }

}
