export default function Pagination({ page, totalPages, onChange }) {
  if (totalPages <= 1) return null

  const pages = []
  const start = Math.max(1, page - 3)
  const end = Math.min(totalPages, page + 3)
  for (let i = start; i <= end; i++) pages.push(i)

  return (
    <div className="flex items-center justify-center gap-1 mt-6">
      <button
        onClick={() => onChange(page - 1)}
        disabled={page <= 1}
        className="px-3 py-1 rounded border text-sm disabled:opacity-30 hover:bg-gray-100"
      >
        上一页
      </button>
      {start > 1 && <span className="px-2 text-gray-400">...</span>}
      {pages.map((p) => (
        <button
          key={p}
          onClick={() => onChange(p)}
          className={`px-3 py-1 rounded border text-sm ${
            p === page ? 'bg-blue-600 text-white border-blue-600' : 'hover:bg-gray-100'
          }`}
        >
          {p}
        </button>
      ))}
      {end < totalPages && <span className="px-2 text-gray-400">...</span>}
      <button
        onClick={() => onChange(page + 1)}
        disabled={page >= totalPages}
        className="px-3 py-1 rounded border text-sm disabled:opacity-30 hover:bg-gray-100"
      >
        下一页
      </button>
    </div>
  )
}
