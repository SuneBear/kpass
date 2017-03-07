import React, { Component, PropTypes } from 'react'
import cx from 'classnames'

// @REF: https://ant.design/components/upload/
import { Upload as AntUpload } from 'antd'

import config from 'config'
import { request } from 'utils'

const UPLOAD_HOST = `${config.FILE_HOST}/upload`

// @TODO: Custom request client & handle error
export class Upload extends Component {

  static propTypes = {
    url: PropTypes.string.isRequired,
    ...AntUpload.propTypes
  }

  static defaultProps = {
    prefixCls: 'upload',
    showUploadList: false
  }

  getHeaders () {
    return {
      Authorization: request.options.headers.Authorization
    }
  }

  getRootClassNames () {
    return cx(
      this.props.className
    )
  }

  render () {
    const {
      url,
      ...props
    } = this.props

    return (
      <AntUpload
        className={this.getRootClassNames()}
        action={`${UPLOAD_HOST}/${url}`}
        headers={this.getHeaders()}
        {...props}
      />
    )
  }
}
