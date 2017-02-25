import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import { reduxForm } from 'redux-form'

import {
  createTeamAction
} from 'modules'

import { TeamCreate as TeamCreateView } from './team-create.view'
import { teamCreateValidate } from './team-create.validate'

const mapStateToProps = (state) => ({})

const mapDispatchToProps = (dispatch) => ({
  actions: bindActionCreators({
    createTeam: createTeamAction
  }, dispatch)
})

const makeContainer = (component) => {
  const connectedComponent = connect(
    mapStateToProps,
    mapDispatchToProps
  )(component)

  return reduxForm({
    form: 'teamCreateForm',
    validate: teamCreateValidate
  })(connectedComponent)
}

export const TeamCreate = makeContainer(TeamCreateView)
