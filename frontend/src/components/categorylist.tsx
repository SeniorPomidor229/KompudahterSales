import React, { useEffect, useState } from 'react';

import { Category } from '../types/product';

const CategoryList: React.FC = () => {
  const [categories, setCategories] = useState<Category[]>([]);

  useEffect(() => {
    fetchCategories();
  }, []);

  const fetchCategories = () => {
    fetch('127.0.0.1:8000/Category')
      .then((response) => response.json())
      .then((data) => setCategories(data))
      .catch((error) => console.error('Error fetching categories:', error));
  };

  return (
    <ul>
      {categories.map((category) => (
        <li key={category._id}>{category.name}</li>
      ))}
    </ul>
  );
};

export default CategoryList;
