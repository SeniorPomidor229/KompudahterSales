import { Product } from "./product";

interface Order {
    _id: string;
    userId: string;
    totalPrice: number;
    products: Product[];
    createdAt: Date;
}