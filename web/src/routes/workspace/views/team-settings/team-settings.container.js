import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import { reduxForm } from 'redux-form'

import { currentTeamSelector, updateCurrentTeamAction } from '../../modules'
import { TeamSettings as TeamSettingsView } from './team-settings.view'
import { teamSettingsValidate } from './team-settings.validate'

const mapStateToProps = (state) => ({
  currentTeam: currentTeamSelector(state)
})

const mapDispatchToProps = (dispatch) => ({
  actions: bindActionCreators({
    updateCurrentTeam: updateCurrentTeamAction
  }, dispatch)
})

const makeContainer = (component) => {
  const connectedComponent = connect(
    mapStateToProps,
    mapDispatchToProps
  )(component)

  return reduxForm({
    form: 'teamSettingsForm',
    validate: teamSettingsValidate
  })(connectedComponent)
}

export const TeamSettings = makeContainer(TeamSettingsView)
