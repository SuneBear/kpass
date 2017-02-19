import React, { Component, PropTypes } from 'react'

import './account-layout.view.styl'

export class AccountLayout extends Component {

  static PropTypes = {
    children: PropTypes.element
  }

  render () {
    return (
      <div className='accountLayout'>
        <div className={'accountLayoutViewport'}>
          <img className={'accountLayoutLogo'} src={require('assets/logo.svg')} />
          {this.props.children}
        </div>
      </div>
    )
  }

}
