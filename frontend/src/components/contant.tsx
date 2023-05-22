import React from 'react';
import { Container, Row, Col } from 'react-bootstrap';

const Contact: React.FC = () => {
  return (
    <div className="contact">
      <Container>
        <Row>
          <Col>
            <h2>Контакты</h2>
            <p>
              Если у вас есть вопросы, предложения или вам требуется дополнительная информация, вы можете связаться с
              нами по следующим контактным данным:
            </p>
          </Col>
        </Row>
        <Row>
          <Col>
            <h4>Телефон</h4>
            <p>+7 123 456 789</p>
          </Col>
          <Col>
            <h4>Email</h4>
            <p>info@example.com</p>
          </Col>
        </Row>
        <Row>
          <Col>
            <h4>Адрес</h4>
            <p>ул. Примерная, д. 123, г. Примерово, Россия</p>
          </Col>
          <Col>
            <h4>График работы</h4>
            <p>Понедельник - Пятница: 9:00 - 18:00</p>
          </Col>
        </Row>
      </Container>
    </div>
  );
};

export default Contact;