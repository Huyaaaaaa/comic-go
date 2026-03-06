import { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
import api from '../api'

export default function Tags() {
  const [tags, setTags] = useState([])
  const [categories, setCategories] = useState([])

  useEffect(() => {
    api.get('/tags').then((r) => setTags(r.data || []))
    api.get('/categories').then((r) => setCategories(r.data || []))
  }, [])

  return (
    <div>
      {categories.length > 0 && (
        <div className="mb-8">
          <h2 className="text-lg font-bold mb-3">分类</h2>
          <div className="flex flex-wrap gap-2">
            {categories.map((c) => (
              <Link key={c.id} to={`/?category_id=${c.id}`} className="px-4 py-2 bg-white rounded-lg shadow-sm hover:shadow text-sm">
                {c.name}
              </Link>
            ))}
          </div>
        </div>
      )}
      <div>
        <h2 className="text-lg font-bold mb-3">标签</h2>
        <div className="flex flex-wrap gap-2">
          {tags.map((t) => (
            <Link key={t.id} to={`/?tag_id=${t.id}`} className="px-4 py-2 bg-white rounded-lg shadow-sm hover:shadow text-sm">
              {t.name}
            </Link>
          ))}
        </div>
      </div>
    </div>
  )
}
