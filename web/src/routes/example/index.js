import { injectViewReducer } from '../../store/reducers'

export default (store) => ({
  path : 'example',
  /*  Async getComponent is only invoked when route matches   */
  getComponent (nextState, cb) {
    require.ensure([], (require) => {
      /*  Webpack - use require callback to define
          dependencies for bundling   */
      const Counter = require('./example.container').default
      const reducer = require('./modules/counter').default

      /*  Add the reducer to the store on key 'counter'  */
      injectViewReducer(store, { key: 'counter', reducer })

      /*  Return getComponent   */
      cb(null, Counter)

    /* Webpack named bundle   */
    }, 'example')
  }
})
