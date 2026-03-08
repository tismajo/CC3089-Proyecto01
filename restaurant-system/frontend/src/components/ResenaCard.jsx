export default function ResenaCard({ resena }) {
  const stars =
    '★'.repeat(resena.calificacion || 0) +
    '☆'.repeat(5 - (resena.calificacion || 0))

  const fechaFormateada = resena.fecha
    ? new Date(resena.fecha).toLocaleDateString()
    : 'Sin fecha'

  return (
    <div className="bg-white rounded-xl shadow p-4 border border-gray-100">
      <div className="flex justify-between items-center mb-1">
        <span className="text-orange-400 text-lg">{stars}</span>
        <span className="text-xs text-gray-400">{fechaFormateada}</span>
      </div>
      <p className="text-gray-700 text-sm">{resena.comentario}</p>
    </div>
  )
}