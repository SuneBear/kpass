import { connect } from 'react-redux'

import { Team as TeamView } from './team.view'

const mapStateToProps = (state) => ({})

const mapDispatchToProps = (dispatch) => ({})

const makeContainer = (component) => {
  return connect(
    mapStateToProps,
    mapDispatchToProps
  )(component)
}

export const Team = makeContainer(TeamView)
