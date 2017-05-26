import React, {Component} from 'react'
import PropTypes from 'prop-types'
import { connect } from 'react-redux'
import { Menu, Icon } from 'semantic-ui-react'

import {signOut, toggleSideBar} from '../actions'


class Widget extends Component {
  state = { activeItem: 'home'}

  handleItemClick = (e, { name }) => this.setState({ activeItem: name })

  render() {
    const { activeItem } = this.state
    const {toggleSideBar} = this.props
    return (
      <Menu inverted fixed="top" color="teal">
        <Menu.Item icon onClick={toggleSideBar}>
          <Icon name="content"/>
          &nbsp;
          <span className="text">Menu</span>
        </Menu.Item>
        <Menu.Item name='home' active={activeItem === 'home'} onClick={this.handleItemClick} />
        <Menu.Item name='messages' active={activeItem === 'messages'} onClick={this.handleItemClick} />
        <Menu.Item name='friends' active={activeItem === 'friends'} onClick={this.handleItemClick} />
      </Menu>
    )
  }
}


Widget.propTypes = {
  user: PropTypes.object.isRequired,
  signOut: PropTypes.func.isRequired,
  toggleSideBar: PropTypes.func.isRequired
}


export default connect(
  state => ({user: state.currentUser}),
  {signOut, toggleSideBar},
)(Widget)
