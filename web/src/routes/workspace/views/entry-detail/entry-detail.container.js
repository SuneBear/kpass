import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'

import { EntryDetail as EntryDetailView } from './entry-detail.view'

const mapStateToProps = (state) => ({})

const mapDispatchToProps = (dispatch) => ({
  actions: bindActionCreators({
  }, dispatch)
})

const makeContainer = (component) => {
  return connect(
    mapStateToProps,
    mapDispatchToProps
  )(component)
}

export const EntryDetail = makeContainer(EntryDetailView)
