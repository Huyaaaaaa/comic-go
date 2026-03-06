import { Link } from 'react-router-dom'

export default function TagList({ tags, activeId }) {
  return (
    <div className="flex flex-wrap gap-2">
      <Link
        to="/"
        className={`px-3 py-1 rounded-full text-sm border ${
          !activeId ? 'bg-blue-600 text-white border-blue-600' : 'hover:bg-gray-100'
        }`}
      >
        全部
      </Link>
      {tags.map((tag) => (
        <Link
          key={tag.id}
          to={`/?tag_id=${tag.id}`}
          className={`px-3 py-1 rounded-full text-sm border ${
            activeId === String(tag.id) ? 'bg-blue-600 text-white border-blue-600' : 'hover:bg-gray-100'
          }`}
        >
          {tag.name}
        </Link>
      ))}
    </div>
  )
}
