import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'

import { privateTeamIdSelector, setCurrentTeamAction } from '../modules'
import { Personal as PersonalView } from './personal.view'

const mapStateToProps = (state) => ({
  privateTeamId: privateTeamIdSelector(state)
})

const mapDispatchToProps = (dispatch) => ({
  actions: bindActionCreators({
    setCurrentTeam: setCurrentTeamAction
  }, dispatch)
})

const makeContainer = (component) => {
  return connect(
    mapStateToProps,
    mapDispatchToProps
  )(component)
}

export const Personal = makeContainer(PersonalView)
