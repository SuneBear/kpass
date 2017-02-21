import React, { Component, PropTypes } from 'react'
import { Link } from 'react-router'
import { Translate } from 'react-redux-i18n'
import cx from 'classnames'

import { Icon } from 'uis'

import './workspace-sidebar.view.styl'

export class WorkspaceSidebar extends Component {

  static propTypes = {
    className: PropTypes.string,
    team: PropTypes.object,
    basePath: PropTypes.string
  }

  getRootClassnames () {
    return cx(
      'workspaceSidebarView',
      this.props.className
    )
  }

  renderSecretNav () {
    const { basePath } = this.props

    return (
      <div className={'navGroup'}>
        <div className={'navGroupTitle'}>
          <Translate value={'pageType.secret'} />
        </div>
        <div className={'navGroupList'}>
          <Link to={`${basePath}/entries`} className={'navItem'} activeClassName={'isActive'}>
            <Icon name={'lock'} />
            <Translate value={'entries.title'} />
          </Link>
        </div>
      </div>
    )
  }

  renderTeamNav () {
    const { team, basePath } = this.props

    if (team.visibility === 'private') {
      return null
    }

    return (
      <div className={'navGroup'}>
        <div className={'navGroupTitle'}>
          <Translate value={'pageType.team'} />
        </div>
        <div className={'navGroupList'}>
          <Link to={`${basePath}/members`} className={'navItem'} activeClassName={'isActive'}>
            <Icon name={'users'} />
            <Translate value={'teamMembers.title'} />
          </Link>
        </div>
      </div>
    )
  }

  render () {
    return (
      <div className={this.getRootClassnames()}>
        <div className={'sidebarNav'}>
          {this.renderSecretNav()}
          {this.renderTeamNav()}
        </div>
      </div>
    )
  }

}
