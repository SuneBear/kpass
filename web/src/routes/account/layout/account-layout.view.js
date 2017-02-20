import React, { Component, PropTypes } from 'react'

import { Logo } from 'views'

import './account-layout.view.styl'

export class AccountLayout extends Component {

  static propTypes = {
    children: PropTypes.element
  }

  render () {
    return (
      <div className='accountLayout'>
        <div className={'accountLayoutViewport'}>
          <Logo className={'accountLayoutLogo'} />
          {this.props.children}
        </div>
      </div>
    )
  }

}
