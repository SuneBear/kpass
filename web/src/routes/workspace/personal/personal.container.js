import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'

import { userMeSelector, readUserMeAction } from 'modules'
import { Personal as PersonalView } from './personal.view'

const mapStateToProps = (state) => ({
  userMe: userMeSelector(state)
})

const mapDispatchToProps = (dispatch) => ({
})

const makeContainer = (component) => {
  return connect(
    mapStateToProps,
    mapDispatchToProps
  )(component)
}

export const Personal = makeContainer(PersonalView)
