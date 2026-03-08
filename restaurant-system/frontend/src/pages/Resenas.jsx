import { useEffect, useState } from 'react'
import { getResenas, createResena } from '../services/resenas'
import { getUsuarios } from '../services/usuarios'
import { getRestaurantes } from '../services/restaurantes'
import ResenaCard from '../components/ResenaCard'

const emptyForm = { calificacion: 5, comentario: '', usuario_id: '', restaurante_id: '' }

export default function Resenas() {
  const [resenas, setResenas] = useState([])
  const [usuarios, setUsuarios] = useState([])
  const [restaurantes, setRestaurantes] = useState([])
  const [loading, setLoading] = useState(true)
  const [showForm, setShowForm] = useState(false)
  const [form, setForm] = useState(emptyForm)
  const [error, setError] = useState(null)

  const cargar = () => {
    setLoading(true)
    Promise.all([getResenas(), getUsuarios(), getRestaurantes()])
      .then(([res, u, r]) => {
        setResenas(res.data || [])
        setUsuarios(u.data || [])
        setRestaurantes(r.data || [])
      })
      .catch(() => setError('Error al cargar datos'))
      .finally(() => setLoading(false))
  }

  useEffect(() => { cargar() }, [])

  const handleSubmit = async (e) => {
    e.preventDefault()
    try {
      await createResena({
        ...form,
        calificacion: parseInt(form.calificacion),
      })
      setShowForm(false)
      setForm(emptyForm)
      cargar()
    } catch {
      alert('Error al crear reseña')
    }
  }

  return (
    <div className="p-6 max-w-7xl mx-auto">
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-2xl font-bold text-gray-800">Reseñas</h1>
        <button
          onClick={() => setShowForm(!showForm)}
          className="bg-orange-600 text-white px-4 py-2 rounded-lg text-sm hover:bg-orange-700"
        >
          {showForm ? 'Cancelar' : '+ Nueva Reseña'}
        </button>
      </div>

      {showForm && (
        <form onSubmit={handleSubmit} className="bg-white rounded-xl shadow p-5 mb-6 grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label className="text-xs text-gray-500 block mb-1">Usuario</label>
            <select required value={form.usuario_id} onChange={(e) => setForm({ ...form, usuario_id: e.target.value })}
              className="border rounded-lg px-3 py-2 text-sm w-full focus:outline-none focus:ring-2 focus:ring-orange-400">
              <option value="">Seleccionar</option>
              {usuarios.map((u) => <option key={u._id} value={u._id}>{u.nombre}</option>)}
            </select>
          </div>
          <div>
            <label className="text-xs text-gray-500 block mb-1">Restaurante</label>
            <select required value={form.restaurante_id} onChange={(e) => setForm({ ...form, restaurante_id: e.target.value })}
              className="border rounded-lg px-3 py-2 text-sm w-full focus:outline-none focus:ring-2 focus:ring-orange-400">
              <option value="">Seleccionar</option>
              {restaurantes.map((r) => <option key={r._id} value={r._id}>{r.nombre}</option>)}
            </select>
          </div>
          <div>
            <label className="text-xs text-gray-500 block mb-1">Calificación (1-5)</label>
            <input type="number" min="1" max="5" value={form.calificacion} onChange={(e) => setForm({ ...form, calificacion: e.target.value })}
              className="border rounded-lg px-3 py-2 text-sm w-full focus:outline-none focus:ring-2 focus:ring-orange-400" />
          </div>
          <div>
            <label className="text-xs text-gray-500 block mb-1">Comentario</label>
            <input placeholder="Tu comentario" value={form.comentario} onChange={(e) => setForm({ ...form, comentario: e.target.value })}
              className="border rounded-lg px-3 py-2 text-sm w-full focus:outline-none focus:ring-2 focus:ring-orange-400" />
          </div>
          <button type="submit" className="md:col-span-2 bg-orange-600 text-white py-2 rounded-lg hover:bg-orange-700 text-sm">
            Guardar Reseña
          </button>
        </form>
      )}

      {loading ? (
        <p className="text-gray-400">Cargando...</p>
      ) : error ? (
        <p className="text-red-500">{error}</p>
      ) : resenas.length === 0 ? (
        <p className="text-gray-400">No hay reseñas registradas.</p>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
          {resenas.map((r) => (
            <ResenaCard key={r._id} resena={r} />
          ))}
        </div>
      )}
    </div>
  )
}
