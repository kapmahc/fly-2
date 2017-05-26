import React from 'react'
import PropTypes from 'prop-types'

import Header from '../components/Header'
import Footer from '../components/Footer'

const Widget = ({children}) => (
  <div>
    <Header />
    <div className="ui container" style={{marginTop: '4em'}}>      
      {children}
    </div>
    <Footer />
  </div>
)

Widget.propTypes = {
  children: PropTypes.node.isRequired
}

export default Widget
