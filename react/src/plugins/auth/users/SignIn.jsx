import React from 'react'
import {FormattedMessage} from 'react-intl'
import {Link} from 'react-router-dom'
import Layout from '../NonSignIn'

const Widget = () => (
  <Layout title="auth.users.sign-in.title">
    <div>
      sign in
      <hr />
      <FormattedMessage id="buttons.cancel" />

    <Link to="/users/sign-up">aaa</Link>
    </div>
  </Layout>
)

export default Widget
