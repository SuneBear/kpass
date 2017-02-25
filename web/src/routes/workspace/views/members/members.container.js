import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'

import { userMeSelector } from 'modules'
import {
  currentTeamSelector,
  currentTeamMembersSelector,
  leaveTeamAction,
  removeTeamMemberAction
} from '../../modules'

import { Members as MembersView } from './members.view'

const mapStateToProps = (state) => ({
  userMe: userMeSelector(state),
  team: currentTeamSelector(state),
  teamMembers: currentTeamMembersSelector(state)
})

const mapDispatchToProps = (dispatch) => ({
  actions: bindActionCreators({
    leaveTeam: leaveTeamAction,
    removeTeamMember: removeTeamMemberAction
  }, dispatch)
})

const makeContainer = (component) => {
  return connect(
    mapStateToProps,
    mapDispatchToProps
  )(component)
}

export const Members = makeContainer(MembersView)
