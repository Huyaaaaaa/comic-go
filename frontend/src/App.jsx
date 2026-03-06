import { BrowserRouter, Routes, Route } from 'react-router-dom'
import Navbar from './components/Navbar'
import Home from './pages/Home'
import Detail from './pages/Detail'
import Reader from './pages/Reader'
import Tags from './pages/Tags'
import Login from './pages/Login'
import Register from './pages/Register'

export default function App() {
  return (
    <BrowserRouter>
      <Navbar />
      <div className="max-w-7xl mx-auto px-4 py-4">
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/comic/:id" element={<Detail />} />
          <Route path="/comic/:id/read" element={<Reader />} />
          <Route path="/tags" element={<Tags />} />
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
        </Routes>
      </div>
    </BrowserRouter>
  )
}
