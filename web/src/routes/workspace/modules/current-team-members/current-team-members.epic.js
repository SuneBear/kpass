import { I18n } from 'react-redux-i18n'
import { push } from 'react-router-redux'
import { combineEpics } from 'redux-observable'
import { Observable } from 'rxjs/Observable'

import { request } from 'utils'
import { toast } from 'uis'
import { deleteTeamSuccessAction, setTeamEntitiesAction } from 'modules'
import { unmountCurrentTeamAction } from '../index'
import {
  createCurrentTeamMemberAction,
  createCurrentTeamMemberSuccessAction,
  createCurrentTeamMemberFailureAction,

  deleteCurrentTeamMemberAction,
  deleteCurrentTeamMemberSuccessAction,
  deleteCurrentTeamMemberFailureAction
} from './current-team-members.reducer'

const createCurrentTeamMemberEpic = (action$) => {
  return action$
    .ofType(`${createCurrentTeamMemberAction}`)
    .switchMap((action) => {
      const { teamId, body, formPromise } = action.payload

      return request
        .post(`teams/${teamId}/invite`, body)
        .takeUntil(action$.ofType(
          `${unmountCurrentTeamAction}`
        ))
        .concatMap((response) => {
          formPromise.resolve(response)

          toast.success({
            message: I18n.t('teamMembers.inviteGenerateSucceed')
          })

          return Observable.of(
            createCurrentTeamMemberSuccessAction()
          )
        })
        .catch((errorMessage) => {
          formPromise.reject(errorMessage)

          const statusCode = errorMessage.error.status

          if (statusCode === 409) {
            toast.error({
              message: I18n.t('teamMembers.inviteRepeated')
            })
          } else {
            toast.error({
              message: I18n.t('teamMembers.inviteFailed')
            })
          }

          return Observable.of(
            createCurrentTeamMemberFailureAction(errorMessage)
          )
        })
    })
}

const deleteCurrentTeamMemberEpic = (action$) => {
  return action$
    .ofType(`${deleteCurrentTeamMemberAction}`)
    .switchMap((action) => {
      const { team, teamId, memberId, isMe } = action.payload

      return request
        .delete(`teams/${teamId}/members/${memberId}`)
        .takeUntil(action$.ofType(
          `${unmountCurrentTeamAction}`
        ))
        .concatMap(() => {
          const successActions = []

          if (isMe) {
            toast.success({
              message: I18n.t('teamMembers.leaveSucceed')
            })

            successActions.push(
              deleteTeamSuccessAction({
                teamId
              }),
              push('/')
            )
          } else {
            toast.success({
              message: I18n.t('teamMembers.removeSucceed')
            })

            // @Hack response
            const parsedTeamEntities = {
              [`${team.id}`]: {
                ...team,
                members: [
                  ...team.members.filter(
                    currentMemberId => currentMemberId !== memberId
                  )
                ]
              }
            }

            successActions.push(
              setTeamEntitiesAction({
                entities: parsedTeamEntities
              })
            )
          }

          return Observable.of(
            ...successActions,
            deleteCurrentTeamMemberSuccessAction()
          )
        })
        .catch((error) => {
          return Observable.of(
            deleteCurrentTeamMemberFailureAction(error)
          )
        })
    })
}

export const currentTeamMembersEpic = combineEpics(
  createCurrentTeamMemberEpic,
  deleteCurrentTeamMemberEpic
)
