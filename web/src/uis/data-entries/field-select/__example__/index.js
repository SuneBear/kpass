import React, { Component } from 'react'
import { Field, reduxForm, propTypes as formPropTypes } from 'redux-form'

import { FieldSelect } from '../index'

/**
 * Example Code:
 *
 * import { FieldSelectExample } from 'uis/data-entries/field-select/__example__'
 * <FieldSelectExample />
 */

class FieldSelectExampleView extends Component {

  static propTypes = {
    ...formPropTypes
  }

  render () {
    const options = [
      { value: 'lucy', title: 'Lucy Go' },
      { value: 'jack' }
    ]

    const defaultInput = {
      value: 'jack'
    }

    return (
      <div className={'fieldSelectExample'}>
        <form>
          <Field
            name={'fieldSelect'}
            input={defaultInput}
            component={FieldSelect}
            options={options}
          />
        </form>
      </div>
    )
  }

}

export const FieldSelectExample = reduxForm({
  form: 'fieldSelectExample'
})(FieldSelectExampleView)
