import { Link } from 'react-router-dom'

export default function ComicCard({ comic }) {
  return (
    <Link to={`/comic/${comic.id}`} className="block group">
      <div className="bg-white rounded-lg shadow hover:shadow-md transition overflow-hidden">
        <div className="aspect-[3/4] overflow-hidden">
          <img
            src={comic.cover_url}
            alt={comic.title}
            className="w-full h-full object-cover group-hover:scale-105 transition"
            loading="lazy"
          />
        </div>
        <div className="p-2">
          <h3 className="text-sm font-medium line-clamp-2 mb-1">{comic.title}</h3>
          <div className="flex items-center gap-2 text-xs text-gray-500">
            <span>⭐ {comic.rating?.toFixed(1) || '0'}({comic.rating_count || 0})</span>
            <span>❤ {comic.favorites || 0}</span>
          </div>
        </div>
      </div>
    </Link>
  )
}
