
import Login from '../pages/Login'
import Signup from '../pages/Signup'
import { Route, Routes } from 'react-router-dom'
import Nav from '../components/Nav'
import { AuthMiddleware } from './Middleware'
import  Messanger from '../pages/Messenger'
import ChatRoom from '../pages/ChatRoom'
function App() {


  return (
    <>
      <Nav />
      <Routes>
        <Route path='/' element={<AuthMiddleware><Messanger /></AuthMiddleware>} />
        <Route path='/login' element={<Login />} />
        <Route path='/signup' element={<Signup />} />
        <Route path='/:room' element={<AuthMiddleware><ChatRoom /></AuthMiddleware>} />
      </Routes>
    </>
  )
}

export default App
