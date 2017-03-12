import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'

import { joinTeamAction } from 'modules'
import { MemberJoin as MemberJoinView } from './member-join.view'

const mapStateToProps = (state) => ({})

const mapDispatchToProps = (dispatch) => ({
  actions: bindActionCreators({
    joinTeam: joinTeamAction
  }, dispatch)
})

const makeContainer = (component) => {
  return connect(
    mapStateToProps,
    mapDispatchToProps
  )(component)
}

export const MemberJoin = makeContainer(MemberJoinView)
