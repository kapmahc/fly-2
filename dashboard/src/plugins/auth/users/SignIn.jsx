import React from 'react'
import {FormattedMessage} from 'react-intl'

import Layout from '../../../layouts/NonSignIn'

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
