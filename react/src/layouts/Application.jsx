import React from 'react'
import PropTypes from 'prop-types'
import { Container } from 'semantic-ui-react'

import Header from '../components/Header'
import Footer from '../components/Footer'

const Widget = ({children}) => (
  <div>
    <Header />
    <div style={{marginTop: '3em'}}/>
    <Container>
      {children}
    </Container>
    <Footer />
  </div>
)

Widget.propTypes = {
  children: PropTypes.node.isRequired
}

export default Widget
