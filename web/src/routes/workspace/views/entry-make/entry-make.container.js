import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import { reduxForm } from 'redux-form'

import {
  currentEntrySelector,
  currentTeamSelector,
  updateCurrentEntryAction,
  createCurrentTeamEntryAction
} from '../../modules'

import { EntryMake as EntryMakeView } from './entry-make.view'
import { entryMakeValidate } from './entry-make.validate'

const mapStateToProps = (state) => ({
  entry: currentEntrySelector(state),
  team: currentTeamSelector(state)
})

const mapDispatchToProps = (dispatch) => ({
  actions: bindActionCreators({
    createEntry: createCurrentTeamEntryAction,
    updateEntry: updateCurrentEntryAction
  }, dispatch)
})

const makeContainer = (component) => {
  const connectedComponent = connect(
    mapStateToProps,
    mapDispatchToProps
  )(component)

  return reduxForm({
    form: 'entryMakeForm',
    validate: entryMakeValidate
  })(connectedComponent)
}

export const EntryMake = makeContainer(EntryMakeView)
