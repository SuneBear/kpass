import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'

import {
  privateTeamIdSelector,
  mountCurrentTeamAction,
  unmountCurrentTeamAction
} from '../modules'
import { Personal as PersonalView } from './personal.view'

const mapStateToProps = (state) => ({
  privateTeamId: privateTeamIdSelector(state)
})

const mapDispatchToProps = (dispatch) => ({
  actions: bindActionCreators({
    mountCurrentTeam: mountCurrentTeamAction,
    unmountCurrentTeam: unmountCurrentTeamAction
  }, dispatch)
})

const makeContainer = (component) => {
  return connect(
    mapStateToProps,
    mapDispatchToProps
  )(component)
}

export const Personal = makeContainer(PersonalView)
