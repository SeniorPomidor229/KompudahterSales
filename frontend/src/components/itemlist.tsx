import React, { useEffect, useState } from 'react';
import { Card, Container, Row, Col } from 'react-bootstrap';
import { Product } from '../types/product';

const ProductsList: React.FC = () => {
  const [products, setProducts] = useState<Product[]>([]);

  useEffect(() => {
    fetchProducts();
  }, []);

  const fetchProducts = () => {
    fetch('127.0.0.1:8000/Product')
      .then((response) => response.json())
      .then((data) => setProducts(data))
      .catch((error) => console.error('Error fetching products:', error));
  };

  return (
    <Container>
      <Row>
        {products.map((product) => (
          <Col key={product._id} md={4} sm={6} className="mb-4">
            <Card>
              <Card.Img variant="top" src={product.photo_url} />
              <Card.Body>
                <Card.Title>{product.name}</Card.Title>
                <Card.Text>{product.description}</Card.Text>
                <Card.Text>Цена: ${product.price}</Card.Text>
                <Card.Text>Категория: {product.category.name}</Card.Text>
              </Card.Body>
            </Card>
          </Col>
        ))}
      </Row>
    </Container>
  );
};

export default ProductsList;
