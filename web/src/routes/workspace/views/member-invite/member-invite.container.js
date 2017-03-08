import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import { reduxForm } from 'redux-form'

import {
  currentTeamSelector,
  currentTeamMembersSelector,
  createCurrentTeamMemberAction
} from '../../modules'

import { MemberInvite as MemberInviteView } from './member-invite.view'
import { memberInviteValidate } from './member-invite.validate'

const mapStateToProps = (state) => ({
  team: currentTeamSelector(state),
  teamMembers: currentTeamMembersSelector(state)
})

const mapDispatchToProps = (dispatch) => ({
  actions: bindActionCreators({
    createTeamMember: createCurrentTeamMemberAction
  }, dispatch)
})

const makeContainer = (component) => {
  const connectedComponent = connect(
    mapStateToProps,
    mapDispatchToProps
  )(component)

  return reduxForm({
    form: 'memberInviteForm',
    validate: memberInviteValidate
  })(connectedComponent)
}

export const MemberInvite = makeContainer(MemberInviteView)
