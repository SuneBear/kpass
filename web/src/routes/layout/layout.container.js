import { connect } from 'react-redux'

import { Layout as LayoutView } from './layout.view'

const mapStateToProps = (state) => ({})

const mapDispatchToProps = (dispatch) => ({})

const makeContainer = (component) => {
  return connect(
    mapStateToProps,
    mapDispatchToProps
  )(component)
}

export const Layout = makeContainer(LayoutView)
