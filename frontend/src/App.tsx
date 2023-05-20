import './App.css';
import { Provider } from 'react-redux';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { store } from './redux/store';
import CustomNavbar from './components/navbar';
import ProductsList from './components/itemlist';
import CategoryList from './components/categorylist';


function App() {
  return (
    <Provider store={store}>

      <Router>

        <CustomNavbar />
        <CategoryList/>

        <main>
          <Routes>
            <Route path='/home' element={<ProductsList/>} />
          </Routes>
        </main>

      </Router>
    </Provider>
  );
}

export default App;
