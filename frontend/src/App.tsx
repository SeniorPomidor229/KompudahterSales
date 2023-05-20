import './App.css';
import { Provider } from 'react-redux';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { store } from './redux/store';
import CustomNavbar from './components/navbar';


function App() {
  return (
    <Provider store={store}>

      <Router>

        <CustomNavbar />


        <main>
          <Routes>
            
          </Routes>
        </main>

      </Router>
    </Provider>
  );
}

export default App;
