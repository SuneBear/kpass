import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import { reduxForm } from 'redux-form'

import {
  currentTeamSelector,
  currentEntrySelector,
  createCurrentTeamEntryAction,
  updateCurrentTeamEntryAction
} from '../../modules'

import { EntryMake as EntryMakeView } from './entry-make.view'
import { entryMakeValidate } from './entry-make.validate'

const mapStateToProps = (state) => ({
  team: currentTeamSelector(state),
  entry: currentEntrySelector(state)
})

const mapDispatchToProps = (dispatch) => ({
  actions: bindActionCreators({
    createEntry: createCurrentTeamEntryAction,
    updateEntry: updateCurrentTeamEntryAction
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
