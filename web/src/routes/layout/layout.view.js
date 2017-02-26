import React, { Component, PropTypes } from 'react'
import DocumentTitle from 'react-document-title'

import config from 'config'

export class Layout extends Component {

  static propTypes = {
    children: PropTypes.element
  }

  render () {
    return (
      <DocumentTitle title={config.NAME}>
        <div className={'rootLayout'}>
          {this.props.children}
        </div>
      </DocumentTitle>
    )
  }

}
