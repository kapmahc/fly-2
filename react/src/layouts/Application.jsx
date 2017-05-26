import React from 'react'
import PropTypes from 'prop-types'
import {Container, Grid} from 'semantic-ui-react'

import Sidebar from '../components/Sidebar'
import Footer from '../components/Footer'
import Header from '../components/Header'

const Widget = ({children}) => (
  <Sidebar>
    <Header />
    <div style={{marginTop: '3em'}}/>
    <Container>      
      <Grid>{children}</Grid>
    </Container>
    <Footer />
  </Sidebar>
)

Widget.propTypes = {
  children: PropTypes.node.isRequired
}

export default Widget
