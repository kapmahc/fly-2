import React from 'react'
import PropTypes from 'prop-types'
import { Layout, Breadcrumb } from 'antd'

import Header from '../components/Header'
import Footer from '../components/Footer'

const { Content } = Layout

const Widget = ({children}) => (
  <Layout className="layout">
    <Header />
    <Content style={{ padding: '0 50px' }}>
      <Breadcrumb style={{ margin: '12px 0' }}>
        <Breadcrumb.Item>Home</Breadcrumb.Item>
        <Breadcrumb.Item>List</Breadcrumb.Item>
        <Breadcrumb.Item>App</Breadcrumb.Item>
      </Breadcrumb>
      <div style={{ background: '#fff', padding: 24, minHeight: 280 }}>
        Content
        <br />
        {children}
      </div>
    </Content>
    <Footer />
  </Layout>
)

Widget.propTypes = {
  children: PropTypes.node.isRequired
}

export default Widget
