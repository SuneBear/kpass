import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import { reduxForm } from 'redux-form'

import {
  currentEntrySelector,
  createCurrentEntrySecretAction,
  updateCurrentEntrySecretAction
} from '../../modules'

import { SecretMake as SecretMakeView } from './secret-make.view'
import { secretMakeValidate } from './secret-make.validate'

const mapStateToProps = (state) => ({
  entry: currentEntrySelector(state)
})

const mapDispatchToProps = (dispatch) => ({
  actions: bindActionCreators({
    createSecret: createCurrentEntrySecretAction,
    updateSecret: updateCurrentEntrySecretAction
  }, dispatch)
})

const makeContainer = (component) => {
  const connectedComponent = connect(
    mapStateToProps,
    mapDispatchToProps
  )(component)

  return reduxForm({
    form: 'secretMakeForm',
    validate: secretMakeValidate
  })(connectedComponent)
}

export const SecretMake = makeContainer(SecretMakeView)
