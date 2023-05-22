import './App.css';
import { Provider } from 'react-redux';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { store } from './redux/store';
import CustomNavbar from './components/navbar';
import ProductsList from './components/itemlist';
import CategoryList from './components/categorylist';
import Login from './components/login';
import Register from './components/register';
import Footer from './components/footer';
import AboutUs from './components/aboutus';
import Contact from './components/contant';


function App() {
  return (
    <Provider store={store}>
      <div className="app-container">
        <Router>
          <div className="content">
            <CustomNavbar />
          </div>

          <main>
            <Routes>
              <Route path='/home' element={<ProductsList />} />
              <Route path='/login' element={<Login />} />
              <Route path='/signup' element={<Register />} />
              <Route path='/about' element={<AboutUs/>}/>
              <Route path='/contact' element={<Contact/>}/>
              <Route path='/services' element={<CategoryList/>} />
            </Routes>
          </main>

          <Footer />

        </Router>
      </div>
    </Provider>
  );
}

export default App;
