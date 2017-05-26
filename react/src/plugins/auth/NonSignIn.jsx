import React from 'react'
import PropTypes from 'prop-types'
import {Grid, Header, Divider, List} from 'semantic-ui-react'
import {FormattedMessage} from 'react-intl'
import {Link} from 'react-router-dom'

import Layout from '../../layouts/Application'
import {NonSignInLinks} from '../../constants'

const Widget = ({title, children}) => (
  <Layout>
    <Grid.Row centered columns={2}>
      <Grid.Column>
        <Header><FormattedMessage id={title}/></Header>
        <Divider />
        {children}
        <br />
        <List>
          {
            NonSignInLinks.map((l, i) => <List.Item as={Link} to={l.href} key={i} icon={l.icon} content={<FormattedMessage id={l.label} />} />)
          }
        </List>
      </Grid.Column>
    </Grid.Row>
  </Layout>
)

Widget.propTypes = {
  title: PropTypes.string.isRequired,
  children: PropTypes.node.isRequired
}

export default Widget
