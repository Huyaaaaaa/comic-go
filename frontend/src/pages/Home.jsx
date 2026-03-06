import { useState, useEffect } from 'react'
import { useSearchParams } from 'react-router-dom'
import Masonry from 'react-masonry-css'
import api from '../api'
import ComicCard from '../components/ComicCard'
import TagList from '../components/TagList'
import Pagination from '../components/Pagination'

const SORT_OPTIONS = [
  { value: 'id', label: '默认' },
  { value: 'newest', label: '最新' },
  { value: 'updated', label: '最近更新' },
  { value: 'rating', label: '评分' },
  { value: 'favorites', label: '收藏' },
]

export default function Home() {
  const [searchParams, setSearchParams] = useSearchParams()
  const [comics, setComics] = useState([])
  const [tags, setTags] = useState([])
  const [page, setPage] = useState(Number(searchParams.get('page')) || 1)
  const [totalPages, setTotalPages] = useState(1)
  const [sort, setSort] = useState(searchParams.get('sort') || 'id')
  const [loading, setLoading] = useState(true)

  const tagId = searchParams.get('tag_id')
  const categoryId = searchParams.get('category_id')

  useEffect(() => {
    api.get('/tags').then((r) => setTags(r.data || []))
  }, [])

  useEffect(() => {
    setLoading(true)
    const params = { page, sort, page_size: 40 }
    if (tagId) params.tag_id = tagId
    if (categoryId) params.category_id = categoryId
    api.get('/comics', { params }).then((r) => {
      setComics(r.data.comics || [])
      setTotalPages(r.data.total_pages || 1)
    }).finally(() => setLoading(false))
  }, [page, sort, tagId, categoryId])

  const changePage = (p) => {
    setPage(p)
    setSearchParams((prev) => { prev.set('page', p); return prev })
    window.scrollTo(0, 0)
  }

  return (
    <div>
      <TagList tags={tags} activeId={tagId} />
      <div className="flex items-center gap-3 mt-4 mb-4">
        <span className="text-sm text-gray-500">排序:</span>
        {SORT_OPTIONS.map((opt) => (
          <button
            key={opt.value}
            onClick={() => { setSort(opt.value); setPage(1) }}
            className={`text-sm px-2 py-1 rounded ${sort === opt.value ? 'bg-blue-600 text-white' : 'text-gray-600 hover:bg-gray-100'}`}
          >
            {opt.label}
          </button>
        ))}
      </div>

      {loading ? (
        <div className="text-center py-20 text-gray-400">加载中...</div>
      ) : comics.length === 0 ? (
        <div className="text-center py-20 text-gray-400">暂无数据</div>
      ) : (
        <Masonry
          breakpointCols={{ default: 5, 1100: 4, 768: 3, 500: 2 }}
          className="flex -ml-4 w-auto"
          columnClassName="pl-4 bg-clip-padding"
        >
          {comics.map((c) => (
            <div key={c.id} className="mb-4">
              <ComicCard comic={c} />
            </div>
          ))}
        </Masonry>
      )}

      <Pagination page={page} totalPages={totalPages} onChange={changePage} />
    </div>
  )
}
