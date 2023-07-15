
import Callroom from '../pages/Callroom'
import Login from '../pages/Login'
import Signup from '../pages/Signup'
import SDPs from '../pages/SDPs'
import { Route, Routes } from 'react-router-dom'
import Nav from '../components/Nav'
import { AuthMiddleware } from './Middleware'


function App() {


  return (
    <>
      <Nav />
      <Routes>
        <Route path='/callroom' element={<AuthMiddleware><Callroom /></AuthMiddleware>} />
        <Route path='/login' element={<Login />} />
        <Route path='/signup' element={<Signup />} />
        <Route path='/available-rooms' element={<AuthMiddleware><SDPs /></AuthMiddleware>} />
      </Routes>
    </>
  )
}

export default App
