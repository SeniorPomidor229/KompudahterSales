import React, { useState } from 'react';
import { Form, Button } from 'react-bootstrap';

const Login: React.FC = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleLogin = () => {
    // Логика для обработки входа пользователя
    console.log('Выполняется вход:', email, password);
  };

  const handleRegister = () => {
    // Логика для обработки регистрации пользователя
    console.log('Выполняется регистрация:', email, password);
  };

  return (
    <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '40vh' }}>
      <div style={{ width: '300px' }}>
        <h2 style={{ textAlign: 'center', marginBottom: '10px' }}>Вход</h2>
        <Form>
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

          <div style={{ display: 'flex', justifyContent: 'center', flexDirection:'column', gap:'15px', marginTop:'15px' }}>
            <Button variant="primary" onClick={handleLogin} >
              Войти
            </Button>
            <Button variant="secondary" onClick={handleRegister}>
              Зарегистрироваться
            </Button>
          </div>
        </Form>
      </div>
    </div>
  );
};

export default Login;
