import React, { Component, PropTypes } from 'react'
import cx from 'classnames'

import { capitalize } from 'utils'
import { Icon } from 'uis'

import './placeholder.view.styl'

const PLACEHOLDER_IMAGES = {
  building: require('./assets/building.svg'),
  checklist: require('./assets/checklist.svg'),
  favorite: require('./assets/favorite.svg'),
  member: require('./assets/member.svg'),
  notice: require('./assets/notice.svg'),
  post: require('./assets/post.svg'),
  task: require('./assets/task.svg')
}

export class Placeholder extends Component {

  static propTypes = {
    className: PropTypes.string,
    size: PropTypes.oneOf(['small', 'normal']),
    iconName: PropTypes.string,
    imageName: PropTypes.string,
    title: PropTypes.string,
    description: PropTypes.string,
    handler: PropTypes.element
  }

  static defaultProps = {
    size: 'normal'
  }

  getRootClassnames () {
    return cx(
      'placeholderView',
      [`size${capitalize(this.props.size)}`],
      this.props.className
    )
  }

  getImageByName (imageName) {
    return PLACEHOLDER_IMAGES[imageName]
  }

  renderImage () {
    const { imageName } = this.props

    const imageUrl = this.getImageByName(imageName)

    if (!imageUrl) {
      return null
    }

    return (
      <img className={'placeholderImage'} src={imageUrl} />
    )
  }

  renderIcon () {
    const { iconName } = this.props

    if (!iconName) {
      return null
    }

    return (
      <Icon className={'placeholderIcon'} name={iconName} />
    )
  }

  renderTitle () {
    const { title } = this.props

    if (!title) {
      return null
    }

    return (
      <div className={'placeholderTitle'}>
        {title}
      </div>
    )
  }

  renderDescription () {
    const { description } = this.props

    if (!description) {
      return null
    }

    return (
      <div className={'placeholderDescription'}>
        {description}
      </div>
    )
  }

  renderHandler () {
    const { handler } = this.props

    if (!handler) {
      return null
    }

    return (
      <div className={'placeholderHandler'}>
        {handler}
      </div>
    )
  }

  render () {
    return (
      <div className={this.getRootClassnames()}>
        {this.renderImage()}
        {this.renderIcon()}
        {this.renderTitle()}
        {this.renderDescription()}
        {this.renderHandler()}
      </div>
    )
  }

}
