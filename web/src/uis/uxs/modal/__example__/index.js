import React, { Component } from 'react'

import { Button } from 'uis'
import { Modal } from '../index'

/**
 * Example Code:
 *
 * import { ModalExample } from 'uis/uxs/modal/__example__'
 * <ModalExample />
 */

export class ModalExample extends Component {

  constructor (props) {
    super(props)

    this.state = {
      opened: false
    }
    this.ref = Object.create(null)
  }

  ref: {
    modal: Modal,
    nestedModal: Modal
  }

  open = () => {
    this.ref.modal.open()
  }

  saveRef = (el) => {
    this.ref.modal = el
  }

  saveRefNested = (el) => {
    this.ref.nestedModal = el
  }

  openNested = () => {
    this.ref.nestedModal.open()
  }

  closeNested = () => {
    this.ref.nestedModal.close()
  }

  onOpen () {
    console.warn('modal opened')
  }

  onClose () {
    console.warn('modal closed')
  }

  render () {
    const nestedModalStyle = { width: '100%', textAlign: 'right' }
    const nestedModalFooter = (
      <div style={nestedModalStyle}>
        <Button size={'small'} onClick={this.closeNested}>Close Me</Button>
      </div>
    )
    return (
      <div>
        <button onClick={this.open}>Open Modal</button>
        <Modal
          title={'Normal Modal'}
          ref={this.saveRef}
        >
          <div>
            <div>
              <button onClick={this.openNested}>
                One More
              </button>
            </div>
            <Modal
              centered
              title={'Nested Modal'}
              size={'small'}
              onOpen={this.onOpen}
              onClose={this.onClose}
              footer={nestedModalFooter}
              ref={this.saveRefNested}
            >
              <div>
                Message for you
              </div>
            </Modal>
          </div>
        </Modal>
      </div>
    )
  }

}
