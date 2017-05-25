import React from 'react'
import PropTypes from 'prop-types'


const Widget = ({children}) => (
  <div>
    <div>dashboard header</div>
    {children}
    <div>footer</div>
  </div>
)

Widget.propTypes = {
  children: PropTypes.node.isRequired
}

export default Widget
