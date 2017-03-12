import { connect } from 'react-redux'
import { push } from 'react-router-redux'
import { bindActionCreators } from 'redux'

import { userMeSelector } from 'modules'
import {
  currentTeamSelector,
  currentTeamSortedEntriesSelector,
  currentTeamEntriesFilterSelector,
  currentEntrySelector,
  currentEntriesBasePathSelector,
  currentUserPermissionsSelector,

  setCurrentEntryAction,
  setCurrentTeamEntriesFilterAction
} from '../../modules'
import { Entries as EntriesView } from './entries.view'

const mapStateToProps = (state) => ({
  userMe: userMeSelector(state),
  team: currentTeamSelector(state),
  currentEntry: currentEntrySelector(state),
  entries: currentTeamSortedEntriesSelector(state),
  entriesBashPath: currentEntriesBasePathSelector(state),
  entriesFilter: currentTeamEntriesFilterSelector(state),
  userPermissions: currentUserPermissionsSelector(state)
})

const mapDispatchToProps = (dispatch) => ({
  actions: bindActionCreators({
    setCurrentEntry: setCurrentEntryAction,
    setCurrentFilter: setCurrentTeamEntriesFilterAction,
    push
  }, dispatch)
})

const makeContainer = (component) => {
  return connect(
    mapStateToProps,
    mapDispatchToProps
  )(component)
}

export const Entries = makeContainer(EntriesView)
