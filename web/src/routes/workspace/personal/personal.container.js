import { connect } from 'react-redux'

import { Personal as PersonalView } from './personal.view'

const mapStateToProps = (state) => ({})

const mapDispatchToProps = (dispatch) => ({})

const makeContainer = (component) => {
  return connect(
    mapStateToProps,
    mapDispatchToProps
  )(component)
}

export const Personal = makeContainer(PersonalView)
