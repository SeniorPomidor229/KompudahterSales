export interface Category {
  _id: string;
  name: string;
  description: string;
}

export interface Product {
  _id: string;
  name: string;
  photo_url: string;
  description: string;
  price: number;
  category: Category;
}
