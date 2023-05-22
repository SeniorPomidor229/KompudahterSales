import React, { useState } from 'react';
import { Form, Button } from 'react-bootstrap';

const Register: React.FC = () => {
  const [firstName, setFirstName] = useState('');
  const [lastName, setLastName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleRegister = () => {
    // Логика для обработки регистрации пользователя
    const newUser = {
      firstName,
      lastName,
      email,
      password
    };
    console.log('Выполняется регистрация:', newUser);
  };

  return (
    <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '50vh' }}>
      <div style={{ width: '300px' }}>
      <h2 style={{ textAlign: 'center', marginBottom: '10px' }}>Регистрация</h2>
      <Form>
        <Form.Group controlId="formFirstName">
          <Form.Label>Имя</Form.Label>
          <Form.Control
            type="text"
            placeholder="Введите имя"
            value={firstName}
            onChange={(e) => setFirstName(e.target.value)}
          />
        </Form.Group>

        <Form.Group controlId="formLastName">
          <Form.Label>Фамилия</Form.Label>
          <Form.Control
            type="text"
            placeholder="Введите фамилию"
            value={lastName}
            onChange={(e) => setLastName(e.target.value)}
          />
        </Form.Group>

        <Form.Group controlId="formEmail">
          <Form.Label>Email</Form.Label>
          <Form.Control
            type="email"
            placeholder="Введите email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
        </Form.Group>

        <Form.Group controlId="formPassword">
          <Form.Label>Пароль</Form.Label>
          <Form.Control
            type="password"
            placeholder="Введите пароль"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </Form.Group>

        <Button variant="primary" onClick={handleRegister} style={{ display: 'flex', justifyContent: 'center', flexDirection:'column', gap:'15px', marginTop:'15px' }}>
          Зарегистрироваться
        </Button>
      </Form>
      </div>
    </div>
  );
};

export default Register;
