import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'

import { userMeSelector, readUserMeAction } from 'modules'
import { Layout as LayoutView } from './layout.view'

const mapStateToProps = (state) => ({
  userMe: userMeSelector(state)
})

const mapDispatchToProps = (dispatch) => ({
  actions: bindActionCreators({
    readUserMe: readUserMeAction
  }, dispatch)
})

const makeContainer = (component) => {
  return connect(
    mapStateToProps,
    mapDispatchToProps
  )(component)
}

export const Layout = makeContainer(LayoutView)
