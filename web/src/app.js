import React, { Component, PropTypes } from 'react'
import { Provider } from 'react-redux'
import { Router } from 'react-router'

export default class App extends Component {
  static propTypes = {
    routes : PropTypes.object.isRequired,
    store  : PropTypes.object.isRequired
  }

  shouldComponentUpdate () {
    return false
  }

  render () {
    const { routes, store } = this.props
    const { history } = store.enhancedUtils

    return (
      <Provider store={store}>
        <div className='app'>
          <Router history={history} children={routes} />
        </div>
      </Provider>
    )
  }
}
