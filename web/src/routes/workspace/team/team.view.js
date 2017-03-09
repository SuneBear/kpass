import React, { Component, PropTypes } from 'react'
import Animate from 'rc-animate'

import './team.view.styl'

export class Team extends Component {

  static propTypes = {
    params: PropTypes.object,
    children: PropTypes.element,
    location: PropTypes.object,
    teams: PropTypes.array,
    currentTeam: PropTypes.object,
    actions: PropTypes.object
  }

  componentDidMount () {
    const { teamId } = this.props.params

    if (!teamId) {
      return
    }

    this.handleComponentMount(teamId)
  }

  componentWillReceiveProps (nextProps) {
    const { currentTeam } = this.props
    const { teamId } = nextProps.params

    if (
      !teamId ||
      teamId === currentTeam.id
    ) {
      return
    }

    this.handleComponentMount(teamId)
  }

  componentDidUpdate () {
    this.handleInvalidTeam()
  }

  componentWillUnmount () {
    const { actions } = this.props

    actions.unmountCurrentTeam()
  }

  isInvalidTeam (team) {
    const { teams } = this.props

    const conditions = [
      () => teams.length,
      () => typeof team.name === 'undefined'
    ]

    return conditions.every((fn) => fn())
  }

  handleInvalidTeam () {
    const { currentTeam, actions } = this.props

    if (this.isInvalidTeam(currentTeam)) {
      return actions.push('/')
    }
  }

  handleComponentMount (nextTeamId) {
    const { actions } = this.props

    actions.mountCurrentTeam({
      teamId: nextTeamId
    })

    this.handleInvalidTeam()
  }

  render () {
    const { children, location, currentTeam } = this.props

    const key = currentTeam.id
    const childrenKey = location.pathname

    return (
      <div className={'workspaceType teamView'} key={key}>
        <Animate
          className={'workspaceCardLayer'}
          transitionName={'switch-card-transition'}
        >
          {React.cloneElement(children, { key: childrenKey })}
        </Animate>
      </div>
    )
  }

}
