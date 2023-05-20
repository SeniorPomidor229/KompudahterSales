import React from 'react';
import { Navbar, Nav } from 'react-bootstrap';
import { Link } from 'react-router-dom';
import logo  from '../imgs/logo.png'

const MyNavbar: React.FC = () => {
  return (
    <Navbar bg="dark" variant="dark" expand="lg" className="pl-5">
      <Navbar.Brand as={Link} to="/">
        <img
          src={logo}
          width="120"
          height="120"
          className="d-inline-block align-top"
          alt="Logo"
        />
      </Navbar.Brand>
      <Navbar.Toggle aria-controls="basic-navbar-nav" />
      <Navbar.Collapse id="basic-navbar-nav">
        <Nav className="mr-auto">
          <Nav.Link as={Link} to="/home">
            Главная
          </Nav.Link>
          <Nav.Link as={Link} to="/about">
            О нас
          </Nav.Link>
          <Nav.Link as={Link} to="/services">
            Категории
          </Nav.Link>
          <Nav.Link as={Link} to="/contact">
            Контакты
          </Nav.Link>
        </Nav>
      </Navbar.Collapse>
    </Navbar>
  );
};

export default MyNavbar;
