import React, { Component, PropTypes } from 'react'
import cx from 'classnames'

import { capitalize } from 'utils'
import { Icon } from 'uis'
import { Avatar } from 'views'

import './menu-selector-item.view.styl'

export class MenuSelectorItem extends Component {

  static propTypes = {
    className: PropTypes.string,
    type: PropTypes.oneOf(['selector', 'divider']),
    handleClick: PropTypes.func,
    value: PropTypes.string,
    isDisabled: PropTypes.bool,
    isSelected: PropTypes.bool,
    title: PropTypes.string,
    description: PropTypes.string,
    avatarUrl: PropTypes.string,
    color: PropTypes.string,
    error: PropTypes.string,
    iconName: PropTypes.string
  }

  static defaultProps = {
    type: 'selector'
  }

  getRootClassNames () {
    return cx(
      'menuSelectorItem',
      `type${capitalize(this.props.type)}`,
      { isSelected: this.props.isSelected },
      { isDisabled: this.props.isDisabled },
      this.props.className
    )
  }

  renderAvatar () {
    const { avatarUrl } = this.props

    if (!avatarUrl) {
      return null
    }

    return (
      <Avatar url={avatarUrl} />
    )
  }

  renderColorLabel () {
    const { color } = this.props

    if (!color) {
      return null
    }

    const style = {
      backgroundColor: color
    }

    return (
      <div className={'colorLabel'} style={style} />
    )
  }

  renderIcon () {
    const { iconName } = this.props

    if (!iconName) {
      return null
    }

    return (
      <Icon name={iconName} />
    )
  }

  renderError () {
    const { error } = this.props

    if (!error) {
      return null
    }

    return (
      <span className={'errorMessage'}>{error}</span>
    )
  }

  renderMainInfo () {
    const { title, value, description } = this.props
    const text = title || value

    if (!text && !description) {
      return null
    }

    return (
      <div className={'mainInfoWrap'}>
        {this.renderTitle()}
        {this.renderDescription()}
      </div>
    )
  }

  renderTitle () {
    const { title, value } = this.props
    const text = title || value

    return (
      <span className={'title'}>{text}</span>
    )
  }

  renderDescription () {
    if (!this.props.description) {
      return null
    }

    return (
      <span className={'description'}>{this.props.description}</span>
    )
  }

  renderSelectedLabel () {
    if (!this.props.isSelected) {
      return null
    }

    return (
      <Icon name={'tick'} className={'selectedLabel'} />
    )
  }

  render () {
    const { handleClick } = this.props
    return (
      <div
        className={this.getRootClassNames()}
        onClick={handleClick}
      >
        {this.renderAvatar()}
        {this.renderColorLabel()}
        {this.renderIcon()}
        {this.renderMainInfo()}
        {this.renderError()}
        {this.renderSelectedLabel()}
      </div>
    )
  }

}
