import React, { Component, PropTypes } from 'react'

const placeConfig = {
  left: {
    points: ['cr', 'cl']
  },
  right: {
    points: ['cl', 'cr']
  },
  top: {
    points: ['bc', 'tc']
  },
  bottom: {
    points: ['tc', 'bc']
  },
  topLeft: {
    points: ['bl', 'tl']
  },
  topRight: {
    points: ['br', 'tr']
  },
  bottomRight: {
    points: ['tr', 'br']
  },
  bottomLeft: {
    points: ['tl', 'bl']
  },
  leftTop: {
    points: ['tr', 'tl']
  },
  leftBottom: {
    points: ['br', 'bl']
  },
  rightTop: {
    points: ['tl', 'tr']
  },
  rightBottom: {
    points: ['bl', 'br']
  }
}

export class Position extends Component {

  static propTypes = {
    children: PropTypes.element,
    placement: PropTypes.oneOf([
      'bottom', 'bottomLeft', 'bottomRight',
      'top', 'topLeft', 'topRight',
      'left', 'leftTop', 'leftBottom',
      'right', 'rightTop', 'rightBottom'
    ]),
    offset: PropTypes.arrayOf(PropTypes.number),
    zIndex: PropTypes.number,
    adjust: PropTypes.bool,
    fixed: PropTypes.bool,
    onPosition: PropTypes.func
  }

  static defaultProps = {
    placement: 'bottom',
    offset: [0, 0],
    zIndex: 1000,
    adjust: true,
    fixed: false
  }

  getAlignConfig () {
    const {
      offset,
      placement
    } = this.props

    const points = placeConfig[placement].points

    const overflow = {
      adjustX: this.props.adjust,
      adjustY: this.props.adjust
    }

    return {
      points,
      overflow,
      offset
    }
  }

  render () {
    const children = this.props.children
    const child = React.Children.only(children)

    const style = {
      position: this.props.fixed ? 'fixed' : 'absolute',
      ...child.props.style
    }

    return React.cloneElement(child, {
      popupAlign: this.getAlignConfig(),
      popupPlacement: this.props.placement,
      builtinPlacements: placeConfig,
      onPopupAlign: this.props.onPosition,
      zIndex: this.props.zIndex,
      popupStyle: style
    })
  }

}
