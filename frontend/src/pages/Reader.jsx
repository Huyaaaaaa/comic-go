import { useState, useEffect } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import api from '../api'

export default function Reader() {
  const { id } = useParams()
  const navigate = useNavigate()
  const [images, setImages] = useState([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    api.get(`/comics/${id}/images`).then((r) => {
      setImages(r.data || [])
    }).finally(() => setLoading(false))
  }, [id])

  if (loading) return <div className="text-center py-20 text-gray-400">加载中...</div>

  return (
    <div className="max-w-3xl mx-auto">
      <div className="sticky top-14 z-40 bg-white/90 backdrop-blur py-2 px-4 flex items-center justify-between border-b">
        <button onClick={() => navigate(`/comic/${id}`)} className="text-blue-600 hover:underline">← 返回详情</button>
        <span className="text-sm text-gray-500">共 {images.length} 张</span>
      </div>
      <div className="flex flex-col items-center">
        {images.map((img, i) => (
          <img
            key={img.id || i}
            src={img.local_path || img.url}
            alt={`第 ${img.sort + 1} 页`}
            className="w-full max-w-2xl"
            loading="lazy"
          />
        ))}
      </div>
      <div className="text-center py-8 text-gray-400">— 完 —</div>
    </div>
  )
}
