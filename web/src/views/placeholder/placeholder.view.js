import React, { Component, PropTypes } from 'react'
import cx from 'classnames'

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
    imageName: PropTypes.string,
    title: PropTypes.string,
    description: PropTypes.string,
    handler: PropTypes.element
  }

  getRootClassnames () {
    return cx(
      'placeholderView',
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
        {this.renderTitle()}
        {this.renderDescription()}
        {this.renderHandler()}
      </div>
    )
  }

}
