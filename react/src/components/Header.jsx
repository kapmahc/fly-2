import React, {Component} from 'react'
import { Collapse, Navbar, NavbarToggler, NavbarBrand, Nav, NavItem, NavLink } from 'reactstrap'

import SearchForm from './SearchForm'

class Widget extends Component {
  constructor(props) {
    super(props)

    this.toggle = this.toggle.bind(this)
    this.state = {
      isOpen: false
    }
  }
  toggle() {
    this.setState({
      isOpen: !this.state.isOpen
    })
  }
  render() {
    return (

      <Navbar color="inverse" inverse toggleable="md" fixed="top">
        <NavbarToggler right onClick={this.toggle} />
        <NavbarBrand href="/">reactstrap</NavbarBrand>
        <Collapse isOpen={this.state.isOpen} navbar>
          <Nav className="mr-auto" navbar>
            <NavItem>
              <NavLink href="/components/">Components</NavLink>
            </NavItem>
            <NavItem>
              <NavLink href="https://github.com/reactstrap/reactstrap">Github</NavLink>
            </NavItem>
          </Nav>
          <SearchForm />
        </Collapse>
      </Navbar>
    )
  }
}

export default Widget
