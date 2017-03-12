import React, { Component, PropTypes } from 'react'
import cx from 'classnames'

import { MenuSelectorItem } from './menu-selector-item.view'

import './menu-selector.view.styl'

// @TODO: Handle Keydown
// - REF: https://github.com/SuneBear/kpass/issues/44
export class MenuSelector extends Component {

  static propTypes = {
    className: PropTypes.string,
    onClick: PropTypes.func,
    onChange: PropTypes.func,
    isMulti: PropTypes.bool,
    identityKey: PropTypes.string,
    hasSelected: PropTypes.array,
    extraList: PropTypes.array,
    dataList: PropTypes.arrayOf(
      PropTypes.shape({
        className: PropTypes.string,
        isDisabled: PropTypes.bool,
        type: PropTypes.string,
        value: PropTypes.string,
        title: PropTypes.string,
        description: PropTypes.string,
        avatarUrl: PropTypes.string,
        color: PropTypes.string,
        error: PropTypes.string,
        iconName: PropTypes.string,
        onClick: PropTypes.func
      })
    ).isRequired,
    ItemRenderer: PropTypes.oneOfType([
      PropTypes.element,
      PropTypes.func
    ])
  }

  static defaultProps = {
    isMulti: false,
    identityKey: 'value',
    hasSelected: [],
    ItemRenderer: MenuSelectorItem
  }

  constructor (props) {
    super(props)
    this.refPopup = Object.create(null)
  }

  getRootClassNames () {
    return cx(
      'menuSelector',
      this.props.className
    )
  }

  isItemSelected (item) {
    const {
      hasSelected,
      identityKey
    } = this.props

    const conditions = [
      () => hasSelected.indexOf(item[identityKey]) >= 0
    ]

    return conditions.every((fn) => fn())
  }

  isItemSelecteable (item) {
    const conditions = [
      () => !item.isDisabled,
      () => typeof item.value !== undefined,
      () => item.type !== 'divider'
    ]

    return conditions.every((fn) => fn())
  }

  getItemKey (item) {
    const {
      identityKey
    } = this.props

    return item[identityKey]
  }

  handleItemClick (item) {
    const {
      isMulti,
      hasSelected,
      onChange
    } = this.props

    if (item.onClick) {
      item.onClick(item)
    }

    if (!this.isItemSelecteable(item)) {
      return
    }

    const itemKey = this.getItemKey(item)

    const getNextHasSelected = (list, key, isMulti) => {
      if (!isMulti) {
        return [key]
      }

      if (list.indexOf(key) >= 0) {
        return list.filter(v => v === key)
      }

      return list.concat(key)
    }

    const nextHasSelected = getNextHasSelected(
      hasSelected, itemKey, isMulti
    )

    if (onChange) {
      onChange(nextHasSelected, item)
    }
  }

  handleClick = () => {
    const {
      onClick,
      isMulti
    } = this.props

    if (!isMulti && onClick) {
      onClick()
    }
  }

  genFilteredList (sourceList) {
    return sourceList
      .map((item) => Object.assign({}, item, {
        isSelected: this.isItemSelected(item),
        handleClick: this.handleItemClick.bind(this, item)
      }))
  }

  renderSelectorList () {
    const {
      dataList,
      ItemRenderer
    } = this.props

    if (!dataList) {
      return null
    }

    const filteredDataList = this.genFilteredList(dataList)

    if (filteredDataList.length === 0) {
      return null
    }

    return (
      <div className={'menuSelectorList'}>
        {filteredDataList.map((item, i) => {
          return <ItemRenderer key={i} {...item} />
        })}
      </div>
    )
  }

  renderExtraList () {
    const {
      extraList,
      ItemRenderer
    } = this.props

    if (!extraList) {
      return null
    }

    const dividerItem = { type: 'divider' }

    const filteredExtraList = this.genFilteredList(
      [dividerItem].concat(extraList)
    )

    return (
      <div className={'menuExtraList'}>
        {filteredExtraList.map((item, i) => {
          return <ItemRenderer key={i} {...item} />
        })}
      </div>
    )
  }

  render () {
    return (
      <div
        className={this.getRootClassNames()}
        onClick={this.handleClick}
      >
        {this.renderSelectorList()}
        {this.renderExtraList()}
      </div>
    )
  }

}
