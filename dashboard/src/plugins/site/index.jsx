import React from 'react';
import { Route } from 'react-router'

import Home from './Home'
import Dashboard from './Dashboard'

export default {
  dashboard: {
    label: 'site.dashboard.title',
    admin: true,
    items: {href:'/site/info', label:'site.admin.info.title'}
  },
  routes: [
    <Route key="site.home" exact path="/" component={Home}/>,
    <Route key="site.dashboard" path="/dashboard" component={Dashboard} />
  ],
}
