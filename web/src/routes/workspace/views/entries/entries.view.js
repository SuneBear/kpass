import React, { Component, PropTypes } from 'react'
import { I18n, Translate } from 'react-redux-i18n'
import cx from 'classnames'

import { Button } from 'uis'
import { Card, Placeholder } from 'views'

import './entries.view.styl'

export class Entries extends Component {

  static propTypes = {
    className: PropTypes.string,
    entries: PropTypes.array
  }

  getRootClassNames () {
    return cx(
      'entriesView',
      this.props.className
    )
  }

  handleNewEntryClick () {

  }

  getNewEntryHandler (type, isGhost) {
    return (
      <Button
        type={type}
        ghost={isGhost}
        icon={'circle-plus'}
        onClick={this.handleAddMemberClick}
      >
        <Translate value={'entry.new'} />
      </Button>
    )
  }

  renderPlaceholder () {
    const { entries } = this.props

    if (entries) {
      return null
    }

    return (
      <Placeholder
        imageName={'task'}
        title={I18n.t('entries.placeholderTitle')}
        description={I18n.t('entries.placeholderDescription')}
        handler={this.getNewEntryHandler('primary', true)}
       />
    )
  }

  renderEntries () {
    const { entries } = this.props

    if (!entries) {
      return null
    }

    return (
      <div>Workspace Entries</div>
    )
  }

  render () {
    return (
      <Card
        className={this.getRootClassNames()}
        title={I18n.t('entries.title')}
        handler={this.getNewEntryHandler('text')}
      >
        {this.renderPlaceholder()}
        {this.renderEntries()}
      </Card>
    )
  }

}
