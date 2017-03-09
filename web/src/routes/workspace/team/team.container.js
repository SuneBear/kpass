import { connect } from 'react-redux'
import { push } from 'react-router-redux'
import { bindActionCreators } from 'redux'

import { teamsSelector } from 'modules'
import {
  currentTeamSelector,
  mountCurrentTeamAction,
  unmountCurrentTeamAction
} from '../modules'
import { Team as TeamView } from './team.view'

const mapStateToProps = (state) => ({
  teams: teamsSelector(state),
  currentTeam: currentTeamSelector(state)
})

const mapDispatchToProps = (dispatch) => ({
  actions: bindActionCreators({
    mountCurrentTeam: mountCurrentTeamAction,
    unmountCurrentTeam: unmountCurrentTeamAction,
    push
  }, dispatch)
})

const makeContainer = (component) => {
  return connect(
    mapStateToProps,
    mapDispatchToProps
  )(component)
}

export const Team = makeContainer(TeamView)
