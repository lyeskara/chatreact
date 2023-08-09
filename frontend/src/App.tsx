
import Login from './pages/Login'
import Signup from './pages/Signup'
import { Route, Routes } from 'react-router-dom'
import { AuthMiddleware } from '../utils/authorizationMiddleware'
import { LoginMiddleware } from '../utils/loginRedirectMiddleware'
import Rooms from './pages/Rooms'
import ChatRoom from './pages/ChatRoom'
function App() {


  return (
    <>
      <Routes>
        <Route path='/' element={<AuthMiddleware><Rooms /></AuthMiddleware>} />
        <Route path='/login' element={<LoginMiddleware><Login /></LoginMiddleware>} />
        <Route path='/signup' element={<LoginMiddleware><Signup /></LoginMiddleware>} />
        <Route path='/:room' element={<AuthMiddleware><ChatRoom /></AuthMiddleware>} />
      </Routes>
    </>
  )
}

export default App
