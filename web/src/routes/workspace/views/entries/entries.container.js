import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'

import { userMeSelector } from 'modules'
import {
  readCurrentTeamEntriesAction,
  currentUserPermissionsSelector,
  currentTeamSelector,
  currentTeamSortedEntriesSelector
} from '../../modules'
import { Entries as EntriesView } from './entries.view'

const mapStateToProps = (state) => ({
  userMe: userMeSelector(state),
  team: currentTeamSelector(state),
  entries: currentTeamSortedEntriesSelector(state),
  userPermissions: currentUserPermissionsSelector(state)
})

const mapDispatchToProps = (dispatch) => ({
  actions: bindActionCreators({
    readCurrentTeamEntries: readCurrentTeamEntriesAction
  }, dispatch)
})

const makeContainer = (component) => {
  return connect(
    mapStateToProps,
    mapDispatchToProps
  )(component)
}

export const Entries = makeContainer(EntriesView)
