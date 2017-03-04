import React, { Component, PropTypes } from 'react'
import Animate from 'rc-animate'

import './team.view.styl'

export class Team extends Component {

  static propTypes = {
    params: PropTypes.object,
    children: PropTypes.element,
    location: PropTypes.object,
    currentTeam: PropTypes.object,
    actions: PropTypes.object
  }

  componentDidMount () {
    const { actions } = this.props
    const { teamId } = this.props.params

    if (!teamId) {
      return
    }

    actions.mountCurrentTeam({
      teamId
    })
  }

  componentWillReceiveProps (nextProps) {
    const { currentTeam, actions } = this.props
    const { teamId } = nextProps.params

    if (
      !teamId ||
      teamId === currentTeam.id
    ) {
      return
    }

    actions.mountCurrentTeam({
      teamId
    })
  }

  componentWillUnmount () {
    const { actions } = this.props

    actions.unmountCurrentTeam()
  }

  render () {
    const { children, location } = this.props

    const key = location.pathname

    return (
      <div className={'workspaceType teamView'}>
        <Animate
          className={'workspaceCardLayer'}
          transitionName={'switch-card-transition'}
        >
          {React.cloneElement(children, { key })}
        </Animate>
      </div>
    )
  }

}
