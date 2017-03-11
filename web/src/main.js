import React from 'react'
import ReactDOM from 'react-dom'
import { I18n } from 'react-redux-i18n'
import { Observable } from 'rxjs/Observable'
import moment from 'moment'

import { request, subscribeHTTPError, cookie } from './utils'
import { toast } from './uis'
import { createStore, initialState } from './store'
import { setUserMeIdAction, setUserEntitiesAction } from './store/modules'
import App from './app'

import './styles/main.styl'

// ========================================================
// Store Instantiation
// ========================================================
const store = createStore(initialState)

// ========================================================
// Render Setup
// ========================================================
const MOUNT_NODE = document.getElementById('root')

let render = () => {
  const routes = require('./routes/index').default(store)
  const state = store.getState()

  // Moment
  moment.locale(state.i18n.locale)

  ReactDOM.render(
    <App store={store} routes={routes} />,
    MOUNT_NODE
  )
}

// This code is excluded from production bundle
if (__DEV__) {
  if (module.hot) {
    // Development render functions
    const renderApp = render
    const renderError = (error) => {
      const RedBox = require('redbox-react').default

      ReactDOM.render(<RedBox error={error} />, MOUNT_NODE)
    }

    // Wrap render in try/catch
    render = () => {
      try {
        renderApp()
      } catch (error) {
        console.error(error)
        renderError(error)
      }
    }

    // Setup hot module replacement
    module.hot.accept('./routes/index', () =>
      setImmediate(() => {
        ReactDOM.unmountComponentAtNode(MOUNT_NODE)
        render()
      })
    )
  }
}

// ========================================================
// Go!
// ========================================================

// @Launch: Auto login & render app
const username = cookie('kp_username')

if (username) {
  request.get(`user/${username}`)
    .take(1)
    .catch((error) => {
      render()
      return Observable.throw(error)
    })
    .subscribe((response) => {
      store.dispatch(setUserMeIdAction({
        userMeId: username
      }))
      store.dispatch(setUserEntitiesAction({
        entities: {
          [`${username}`]: response
        }
      }))
      render()
    })
} else {
  render()
}

// @SideEffect: Handle HTTP Error
subscribeHTTPError((res) => {
  const status = res.error.status
  switch (status) {
    case 401:
      return toast.error({
        message: I18n.t('account.unauthorized')
      })
    case 403:
      return toast.error({
        message: I18n.t('team.isFrozen')
      })
  }
})
