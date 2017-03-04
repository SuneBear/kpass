import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import { reduxForm } from 'redux-form'

import { updateCurrentTeamAction } from '../../modules'
import { TeamRename as TeamRenameView } from './team-rename.view'
import { teamRenameValidate } from './team-rename.validate'

const mapStateToProps = (state) => ({})

const mapDispatchToProps = (dispatch) => ({
  actions: bindActionCreators({
    updateTeam: updateCurrentTeamAction
  }, dispatch)
})

const makeContainer = (component) => {
  const connectedComponent = connect(
    mapStateToProps,
    mapDispatchToProps
  )(component)

  return reduxForm({
    form: 'teamRenameForm',
    validate: teamRenameValidate
  })(connectedComponent)
}

export const TeamRename = makeContainer(TeamRenameView)
