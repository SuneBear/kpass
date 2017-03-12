import React, { Component, PropTypes } from 'react'
import { I18n } from 'react-redux-i18n'

import { Loading } from 'uis'
import { Card } from 'views'

export class MemberJoin extends Component {

  static propTypes = {
    params: PropTypes.object,
    actions: PropTypes.object
  }

  componentWillMount () {
    const { params, actions } = this.props

    actions.joinTeam({
      body: {
        code: params.inviteCode
      }
    })
  }

  render () {
    return (
      <Card
        className={'memberJoinView'}
        title={I18n.t('teamMembers.joining')}
      >
        <Loading />
      </Card>
    )
  }

}
