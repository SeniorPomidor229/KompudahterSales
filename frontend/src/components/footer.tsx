import React from 'react';
import { Container, Row, Col } from 'react-bootstrap';

const Footer: React.FC = () => {
    return (
      <footer className="bg-dark text-light mt-auto">
        <Container>
          <Row>
            <Col md={6}>
              <h5>О нас</h5>
              <p>Некоторая информация о вашей компании или проекте.</p>
            </Col>
            <Col md={6}>
              <h5>Контакты</h5>
              <p>Адрес: город, улица, дом</p>
              <p>Телефон: +7 123 456 789</p>
              <p>Email: info@example.com</p>
            </Col>
          </Row>
        </Container>
        <div className="text-center py-3">
          &copy; {new Date().getFullYear()} Ваше название компании или проекта
        </div>
      </footer>
    );
  };

export default Footer;