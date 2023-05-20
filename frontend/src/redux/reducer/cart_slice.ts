import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import type { RootState } from '../store';
import { Product } from '../../types/product';

interface CartItems {
  value: Product[];
}

const initialState: CartItems = {
  value: [],
};

export const cartSlice = createSlice({
  name: 'cart',
  initialState,
  reducers: {
    add: (state, action: PayloadAction<Product>) => {
      state.value.push(action.payload);
    },
    remove: (state, action: PayloadAction<Product>) => {
      state.value = state.value.filter(item => item._id !== action.payload._id);
    },
  },
});

export const { add, remove } = cartSlice.actions;

export const selectCart = (state: RootState) => state;

export default cartSlice.reducer;
