import { useState } from 'react'
import { useNavigate, Link } from 'react-router-dom'
import api from '../api'

export default function Register() {
  const [username, setUsername] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState('')
  const navigate = useNavigate()

  const submit = async (e) => {
    e.preventDefault()
    setError('')
    try {
      const r = await api.post('/auth/register', { username, email, password })
      localStorage.setItem('token', r.data.token)
      localStorage.setItem('user', JSON.stringify(r.data.user))
      navigate('/')
      window.location.reload()
    } catch (err) {
      setError(err.response?.data?.error || '注册失败')
    }
  }

  return (
    <div className="max-w-sm mx-auto mt-20">
      <h1 className="text-2xl font-bold mb-6 text-center">注册</h1>
      {error && <p className="text-red-500 text-sm mb-4 text-center">{error}</p>}
      <form onSubmit={submit} className="space-y-4">
        <input type="text" placeholder="用户名" value={username} onChange={(e) => setUsername(e.target.value)}
          className="w-full px-4 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500" required minLength={3} />
        <input type="email" placeholder="邮箱" value={email} onChange={(e) => setEmail(e.target.value)}
          className="w-full px-4 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500" required />
        <input type="password" placeholder="密码 (至少6位)" value={password} onChange={(e) => setPassword(e.target.value)}
          className="w-full px-4 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500" required minLength={6} />
        <button type="submit" className="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700">注册</button>
      </form>
      <p className="text-center mt-4 text-sm text-gray-500">
        已有账号？<Link to="/login" className="text-blue-600 hover:underline">登录</Link>
      </p>
    </div>
  )
}
