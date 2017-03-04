import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'

import {
  deleteCurrentTeamEntryAction
} from '../../modules'
import { EntryItem as EntryItemView } from './entry-item.view'

const mapStateToProps = (state) => ({})

const mapDispatchToProps = (dispatch) => ({
  actions: bindActionCreators({
    deleteCurrentTeamEntry: deleteCurrentTeamEntryAction
  }, dispatch)
})

const makeContainer = (component) => {
  return connect(
    mapStateToProps,
    mapDispatchToProps
  )(component)
}

export const EntryItem = makeContainer(EntryItemView)
