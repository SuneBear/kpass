import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'

import { currentTeamSelector, updateCurrentTeamAction } from '../../modules'
import { TeamSettings as TeamSettingsView } from './team-settings.view'

const mapStateToProps = (state) => ({
  team: currentTeamSelector(state)
})

const mapDispatchToProps = (dispatch) => ({
  actions: bindActionCreators({
    updateTeam: updateCurrentTeamAction
  }, dispatch)
})

const makeContainer = (component) => {
  return connect(
    mapStateToProps,
    mapDispatchToProps
  )(component)
}

export const TeamSettings = makeContainer(TeamSettingsView)
