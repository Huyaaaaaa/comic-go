import { useState, useEffect } from 'react'
import { useParams, Link, useNavigate } from 'react-router-dom'
import api from '../api'

export default function Detail() {
  const { id } = useParams()
  const navigate = useNavigate()
  const [comic, setComic] = useState(null)
  const [loading, setLoading] = useState(true)
  const [isFav, setIsFav] = useState(false)

  useEffect(() => {
    api.get(`/comics/${id}`).then((r) => {
      setComic(r.data)
    }).finally(() => setLoading(false))
  }, [id])

  const toggleFav = async () => {
    try {
      if (isFav) {
        await api.delete(`/user/favorites/${id}`)
        setIsFav(false)
      } else {
        await api.post(`/user/favorites/${id}`)
        setIsFav(true)
      }
    } catch {
      navigate('/login')
    }
  }

  const rate = async (score) => {
    try {
      const r = await api.post(`/comics/${id}/rate`, { score })
      setComic((prev) => ({ ...prev, rating: r.data.rating, rating_count: r.data.count }))
    } catch {
      navigate('/login')
    }
  }

  if (loading) return <div className="text-center py-20 text-gray-400">加载中...</div>
  if (!comic) return <div className="text-center py-20 text-gray-400">漫画不存在</div>

  return (
    <div>
      <button onClick={() => navigate(-1)} className="text-blue-600 mb-4 hover:underline">← 返回</button>
      <div className="flex flex-col md:flex-row gap-6">
        <img src={comic.cover_url} alt={comic.title} className="w-64 h-80 object-cover rounded-lg shadow" />
        <div className="flex-1">
          <h1 className="text-2xl font-bold mb-2">{comic.title}</h1>
          {comic.subtitle && <p className="text-gray-600 mb-2">{comic.subtitle}</p>}
          {comic.author && <p className="text-sm text-gray-500 mb-2">作者: {comic.author}</p>}
          {comic.category_name && <p className="text-sm text-gray-500 mb-2">分类: {comic.category_name}</p>}
          <div className="flex items-center gap-1 mb-3">
            {[...Array(10)].map((_, i) => (
              <button key={i} onClick={() => rate(i + 1)} className="text-lg hover:scale-125 transition">
                {i < Math.round(comic.rating || 0) ? '⭐' : '☆'}
              </button>
            ))}
            <span className="text-sm text-gray-500 ml-2">{comic.rating?.toFixed(1)}({comic.rating_count})</span>
          </div>
          <div className="flex flex-wrap gap-2 mb-4">
            {comic.tags?.map((t) => (
              <Link key={t.id} to={`/?tag_id=${t.id}`} className="px-2 py-1 bg-gray-100 rounded text-sm hover:bg-blue-100">
                {t.name}
              </Link>
            ))}
          </div>
          <div className="flex gap-3">
            <Link to={`/comic/${id}/read`} className="bg-blue-600 text-white px-6 py-2 rounded hover:bg-blue-700">
              阅读
            </Link>
            <button onClick={toggleFav} className={`px-6 py-2 rounded border ${isFav ? 'bg-red-50 border-red-300 text-red-600' : 'hover:bg-gray-50'}`}>
              {isFav ? '已收藏' : '收藏'}
            </button>
          </div>
          <div className="mt-4 text-xs text-gray-400">
            <p>创建: {comic.created_at}</p>
            <p>更新: {comic.updated_at}</p>
          </div>
        </div>
      </div>
    </div>
  )
}
