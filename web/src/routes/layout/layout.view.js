import React, { Component, PropTypes } from 'react'
import DocumentTitle from 'react-document-title'

import config from 'config'

export class Layout extends Component {

  static PropTypes = {
    userMe: PropTypes.object,
    actions: PropTypes.object
  }

  componentDidMount () {
    this.props.actions.readUserMe()
  }

  render () {
    return (
      <DocumentTitle title={config.NAME}>
        <div className='rootLayout'>
          { this.props.children }
        </div>
      </DocumentTitle>
    )
  }

}
