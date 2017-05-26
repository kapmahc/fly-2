import React from 'react'
import {FormattedMessage} from 'react-intl'

import Layout from '../NonSignIn'

const Widget = () => (
  <Layout>
    <div>
      sign in
      <hr />
      <FormattedMessage id="buttons.cancel" />
    </div>
  </Layout>
)

export default Widget
