import { applyMiddleware, compose, createStore as reduxCreateStore } from 'redux'
import { createEpicMiddleware } from 'redux-observable'
import { I18n, syncTranslationWithStore } from 'react-redux-i18n'
import { browserHistory } from 'react-router'
import { syncReduxAndTitle } from 'redux-title'
import { routerMiddleware, syncHistoryWithStore } from 'react-router-redux'

import { subscribeHTTPError } from 'utils'
import { toast } from 'uis'
import { makeRootEpic, makeRootReducer } from './modules'

export const createStore = (initialState = {}) => {
  // ======================================================
  // Middleware Configuration
  // ======================================================
  const middleware = [
    routerMiddleware(browserHistory),
    createEpicMiddleware(makeRootEpic())
  ]

  // ======================================================
  // Store Enhancers
  // ======================================================
  const enhancers = []

  let composeEnhancers = compose

  if (__DEV__) {
    const composeWithDevToolsExtension = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__
    if (typeof composeWithDevToolsExtension === 'function') {
      composeEnhancers = composeWithDevToolsExtension
    }
  }

  // ======================================================
  // Store Instantiation and HMR Setup
  // ======================================================
  const store = reduxCreateStore(
    makeRootReducer(),
    initialState,
    composeEnhancers(
      applyMiddleware(...middleware),
      ...enhancers
    )
  )

  // Sync with Store
  syncTranslationWithStore(store)
  syncReduxAndTitle(store,
    (state) => store.getState().context.title
  )
  const history = syncHistoryWithStore(browserHistory, store, {
    selectLocationState: (state) => state.context.routing
  })

  // @SideEffect: Handle HTTP Error
  subscribeHTTPError((res) => {
    const status = res.error.status
    switch (status) {
      case 401:
        return toast.error({
          message: I18n.t('account.unauthorized')
        })
    }
  })

  // @Property: Async Reducers
  store.asyncReducers = {}

  // @Property: Enhanced Utils
  store.enhancedUtils = {
    history
  }

  if (module.hot) {
    module.hot.accept('./modules/root.reducer', () => {
      const reducers = require('./modules/root.reducer').makeRootReducer
      store.replaceReducer(reducers(store.asyncReducers))
    })
  }

  return store
}
