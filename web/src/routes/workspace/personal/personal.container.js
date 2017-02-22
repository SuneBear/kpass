import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'

import { setCurrentTeamAction } from 'modules'
import { Personal as PersonalView } from './personal.view'

const mapStateToProps = (state) => ({})

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
