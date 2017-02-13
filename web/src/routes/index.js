import ExampleRoute from './example'

export const createRoutes = (store) => ({
  path          : '/',
  indexRoute    : { onEnter: (nextState, replace) => replace('/example') },
  childRoutes   : [
    ExampleRoute(store)
  ]
})

export default createRoutes
