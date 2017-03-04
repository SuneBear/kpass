import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'

import {
  currentTeamSelector,
  mountCurrentTeamAction,
  unmountCurrentTeamAction
} from '../modules'
import { Team as TeamView } from './team.view'

const mapStateToProps = (state) => ({
  currentTeam: currentTeamSelector(state)
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

export const Team = makeContainer(TeamView)
