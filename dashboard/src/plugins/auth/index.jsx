import React from 'react'
import { Route } from 'react-router'

import SignIn from './users/SignIn'
import SignUp from './users/SignUp'

export default {
  dashboard: {
    label: 'auth.dashboard.title',
    items: {href:'/users/logs', label:'auth.users.logs.title'}
  },
  routes: [
    <Route key="auth.users.sign-in" path="/users/sign-in" component={SignIn} />,
    <Route key="auth.users.sign-up" path="/users/sign-up" component={SignUp} />
  ],
}
