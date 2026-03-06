import { useState } from 'react'
import { useNavigate, Link } from 'react-router-dom'
import api from '../api'

export default function Login() {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState('')
  const navigate = useNavigate()

  const submit = async (e) => {
    e.preventDefault()
    setError('')
    try {
      const r = await api.post('/auth/login', { username, password })
      localStorage.setItem('token', r.data.token)
      localStorage.setItem('user', JSON.stringify(r.data.user))
      navigate('/')
      window.location.reload()
    } catch (err) {
      setError(err.response?.data?.error || '登录失败')
    }
  }

  return (
    <div className="max-w-sm mx-auto mt-20">
      <h1 className="text-2xl font-bold mb-6 text-center">登录</h1>
      {error && <p className="text-red-500 text-sm mb-4 text-center">{error}</p>}
      <form onSubmit={submit} className="space-y-4">
        <input type="text" placeholder="用户名" value={username} onChange={(e) => setUsername(e.target.value)}
          className="w-full px-4 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500" required />
        <input type="password" placeholder="密码" value={password} onChange={(e) => setPassword(e.target.value)}
          className="w-full px-4 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500" required />
        <button type="submit" className="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700">登录</button>
      </form>
      <p className="text-center mt-4 text-sm text-gray-500">
        没有账号？<Link to="/register" className="text-blue-600 hover:underline">注册</Link>
      </p>
    </div>
  )
}
