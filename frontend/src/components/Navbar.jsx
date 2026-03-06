import { Link, useNavigate } from 'react-router-dom'
import { useState, useEffect } from 'react'

export default function Navbar() {
  const [user, setUser] = useState(null)
  const navigate = useNavigate()

  useEffect(() => {
    const stored = localStorage.getItem('user')
    if (stored) setUser(JSON.parse(stored))
  }, [])

  const logout = () => {
    localStorage.removeItem('token')
    localStorage.removeItem('user')
    setUser(null)
    navigate('/')
  }

  return (
    <nav className="bg-white shadow-sm sticky top-0 z-50">
      <div className="max-w-7xl mx-auto px-4 flex items-center justify-between h-14">
        <div className="flex items-center gap-6">
          <Link to="/" className="text-xl font-bold text-blue-600">漫画站</Link>
          <Link to="/" className="text-gray-600 hover:text-blue-600">首页</Link>
          <Link to="/tags" className="text-gray-600 hover:text-blue-600">标签</Link>
        </div>
        <div className="flex items-center gap-4">
          {user ? (
            <>
              <span className="text-sm text-gray-600">{user.username}</span>
              <button onClick={logout} className="text-sm text-red-500 hover:text-red-700">退出</button>
            </>
          ) : (
            <>
              <Link to="/login" className="text-sm text-gray-600 hover:text-blue-600">登录</Link>
              <Link to="/register" className="text-sm bg-blue-600 text-white px-3 py-1 rounded hover:bg-blue-700">注册</Link>
            </>
          )}
        </div>
      </div>
    </nav>
  )
}
