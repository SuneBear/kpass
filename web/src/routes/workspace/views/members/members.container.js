import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'

import { userMeSelector } from 'modules'
import {
  currentTeamSelector,
  currentTeamMembersSelector,
  currentUserPermissionsSelector,
  deleteCurrentTeamMemberAction
} from '../../modules'

import { Members as MembersView } from './members.view'

const mapStateToProps = (state) => ({
  userMe: userMeSelector(state),
  team: currentTeamSelector(state),
  teamMembers: currentTeamMembersSelector(state),
  userPermissions: currentUserPermissionsSelector(state)
})

const mapDispatchToProps = (dispatch) => ({
  actions: bindActionCreators({
    removeMember: deleteCurrentTeamMemberAction
  }, dispatch)
})

const makeContainer = (component) => {
  return connect(
    mapStateToProps,
    mapDispatchToProps
  )(component)
}

export const Members = makeContainer(MembersView)
