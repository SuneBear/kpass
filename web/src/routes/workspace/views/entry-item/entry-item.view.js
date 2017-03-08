import React, { Component, PropTypes } from 'react'
import { I18n } from 'react-redux-i18n'
import cx from 'classnames'
import { pick } from 'lodash'
import moment from 'moment'

import { Badge, Datestamp } from 'views'
import { Icon, Dropdown, MenuSelector, Modal, Tooltip } from 'uis'
import { EntryMake } from '../entry-make'
import { getEntryCategoryByValue } from '../../shared'

import './entry-item.view.styl'

export class EntryItem extends Component {

  static propTypes = {
    className: PropTypes.string,
    entry: PropTypes.object,
    creatorPermissions: PropTypes.object,
    extraContent: PropTypes.any,
    actions: PropTypes.object
  }

  getRootClassnames () {
    return cx(
      'entryItemView',
      this.props.className
    )
  }

  saveEditEntryModalRef = (ref) => {
    this.editEntryModalRef = ref
  }

  handleEditEntryClick = () => {
    this.editEntryModalRef.open()
  }

  handleEditEntrySubmitSuccess = () => {
    this.editEntryModalRef.close()
  }

  handleDeleteEntry = () => {
    const { entry, actions } = this.props

    actions.deleteEntry({
      entryId: entry.id
    })
  }

  renderHero () {
    return (
      <div className={'entryItemHero typeEntry'}>
        <Icon name={'lock'} />
      </div>
    )
  }

  renderTitle () {
    const { entry } = this.props

    return (
      <div className={'entryItemTitle'}>
        {entry.name}
      </div>
    )
  }

  renderMeta () {
    return (
      <div className={'entryItemMeta'}>
        {this.renderMetaDate()}
        {this.renderMetaCategory()}
      </div>
    )
  }

  renderMetaDate () {
    const { entry } = this.props

    const dateTip = `${I18n.t('entry.updatedDate')}: ${moment(entry.updated).format('LLL')}`

    return (
      <Tooltip title={dateTip}>
        <div className={'entryItemMetaItem'}>
          <Icon name={'calendar2'} />
          <div className={'entryItemMetaText'}>
            <Datestamp date={moment(entry.updated).toDate()} />
          </div>
        </div>
      </Tooltip>
    )
  }

  renderMetaCategory () {
    const { entry } = this.props

    const category = getEntryCategoryByValue(entry.category)
    const categoryTip = `${I18n.t('entry.category')}: ${category.title}`

    if (!category) {
      return null
    }

    return (
      <Tooltip title={categoryTip}>
        <div className={'entryItemMetaItem'}>
          <Badge bgColor={category.color} />
          <div className={'entryItemMetaText'}>
            {category.title}
          </div>
        </div>
      </Tooltip>
    )
  }

  renderMore () {
    const { creatorPermissions } = this.props

    const conditions = [
      () => creatorPermissions.updateEntry,
      () => creatorPermissions.deleteEntry
    ]

    if (!conditions.some((fn) => fn())) {
      return null
    }

    return (
      <Dropdown
        className={'entryItemMoreDropdown'}
        content={this.getMoreMenu()}
      >
        <Tooltip
          title={I18n.t('base.more')}
          onClick={(e) => e.stopPropagation()}
        >
          <div className={'entryItemMoreHandler'}>
            <Icon name={'more'} />
          </div>
        </Tooltip>
      </Dropdown>
    )
  }

  getMoreMenu () {
    const { creatorPermissions } = this.props

    const dataList = []

    if (creatorPermissions.updateEntry) {
      dataList.push({
        iconName: 'pencil',
        title: I18n.t('entry.edit'),
        onClick: this.handleEditEntryClick
      })
    }

    if (creatorPermissions.deleteEntry) {
      dataList.push({
        iconName: 'trash',
        title: I18n.t('entry.delete'),
        onClick: this.handleDeleteEntry
      })
    }

    return (
      <MenuSelector
        dataList={dataList}
      />
    )
  }

  renderExtra () {
    const { extraContent } = this.props

    if (!extraContent) {
      return null
    }

    return (
      <div className={'entryItemExtra'}>
        {extraContent}
      </div>
    )
  }

  renderEditEntryModal () {
    const { entry } = this.props

    const initialValues = pick(entry, [
      'name', 'category', 'priority'
    ])

    return (
      <Modal
        ref={this.saveEditEntryModalRef}
        title={I18n.t('entry.edit')}
        size={'small'}
      >
        <EntryMake
          action={'update'}
          entryId={entry.id}
          initialValues={initialValues}
          onSubmitSuccess={this.handleEditEntrySubmitSuccess}
        />
      </Modal>
    )
  }

  render () {
    return (
      <div className={this.getRootClassnames()}>
        {this.renderHero()}
        <div className={'entryItemContent'}>
          {this.renderTitle()}
          {this.renderMeta()}
        </div>
        <div className={'entryItemHandlerset'}>
          {this.renderMore()}
          {this.renderExtra()}
        </div>

        {this.renderEditEntryModal()}
      </div>
    )
  }

}
