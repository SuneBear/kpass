import React, { Component, PropTypes } from 'react'
import cx from 'classnames'
import marked from 'marked'

import './readable.view.styl'

export class Readable extends Component {

  static propTypes = {
    className: PropTypes.string,
    type: PropTypes.oneOf(['html', 'markdown']),
    content: PropTypes.oneOfType([
      PropTypes.element,
      PropTypes.string
    ]),
    options: PropTypes.object
  }

  static defaultProps = {
    type: 'html'
  }

  getRootClassNames () {
    return cx(
      'readableView',
      this.props.className
    )
  }

  renderContent () {
    let { type, content, options } = this.props

    if (type === 'markdown') {
      const markdownOptions = {
        breaks: true,
        sanitize: true,
        smartLists: true,
        ...options
      }
      content = marked(content, markdownOptions)
    }

    return (
      <div
        className={'readableContent'}
        dangerouslySetInnerHTML={{ __html: content }}
      />
    )
  }

  render () {
    return (
      <div className={this.getRootClassNames()}>
        {this.renderContent()}
      </div>
    )
  }

}
