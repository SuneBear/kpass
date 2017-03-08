import React, { Component, PropTypes } from 'react'
import cx from 'classnames'

import { Button } from 'uis'
import { EntryItem } from '../entry-item'

import './entry-detail.view.styl'

export class EntryDetail extends Component {

  static propTypes = {
    className: PropTypes.string,
    entry: PropTypes.object,
    creatorPermissions: PropTypes.object,
    actions: PropTypes.object
  }

  getRootClassnames () {
    return cx(
      'entryDetailView',
      this.props.className
    )
  }

  renderEntryItem () {
    const {
      entry,
      creatorPermissions
    } = this.props

    return (
      <EntryItem
        entry={entry}
        creatorPermissions={creatorPermissions}
      />
    )
  }

  render () {
    return (
      <div
        className={this.getRootClassnames()}
      >
        <div className={'entryDetailHeader'}>
          {this.renderEntryItem()}
        </div>
        <div className={'entryDetailDivider'} />
        <div className={'entryDetailContent'}>
          <Button block type={'normal'}>@TODO: Secret</Button>
        </div>
      </div>
    )
  }

}
