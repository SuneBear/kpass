import { connect } from 'react-redux'

import { Entries as EntriesView } from './entries.view'

const mapStateToProps = (state) => ({})

const mapDispatchToProps = (dispatch) => ({})

const makeContainer = (component) => {
  return connect(
    mapStateToProps,
    mapDispatchToProps
  )(component)
}

export const Entries = makeContainer(EntriesView)
