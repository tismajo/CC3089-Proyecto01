import { useEffect, useState } from 'react'
import { getMejoresRestaurantes, getVentasPorMes } from '../services/reportes'
import { getRestaurantes } from '../services/restaurantes'
import { getOrdenes } from '../services/ordenes'
import { getUsuarios } from '../services/usuarios'
import { getResenas } from '../services/resenas'
import {
  BarChart, Bar, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer, Legend,
} from 'recharts'

function StatCard({ label, value, color }) {
  return (
    <div className={`bg-white rounded-xl shadow p-5 border-l-4 ${color}`}>
      <p className="text-sm text-gray-500">{label}</p>
      <p className="text-3xl font-bold text-gray-800 mt-1">{value}</p>
    </div>
  )
}

export default function Dashboard() {
  const [mejores, setMejores] = useState([])
  const [ventas, setVentas] = useState([])
  const [stats, setStats] = useState({ restaurantes: 0, ordenes: 0, usuarios: 0, resenas: 0 })
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState(null)

  useEffect(() => {
    Promise.all([
      getMejoresRestaurantes(),
      getVentasPorMes(),
      getRestaurantes(),
      getOrdenes(),
      getUsuarios(),
      getResenas(),
    ])
      .then(([r1, r2, rest, ord, usu, res]) => {
        setMejores(r1.data || [])
        setVentas(r2.data || [])
        setStats({
          restaurantes: (rest.data || []).length,
          ordenes: (ord.data || []).length,
          usuarios: (usu.data || []).length,
          resenas: (res.data || []).length,
        })
      })
      .catch(() => setError('No se pudo conectar al servidor'))
      .finally(() => setLoading(false))
  }, [])

  if (loading) return <p className="p-8 text-gray-500">Cargando...</p>

  if (error) return (
    <div className="p-8">
      <div className="bg-red-50 border border-red-200 text-red-700 rounded-xl p-4">
        <p className="font-semibold">Sin conexión al servidor</p>
        <p className="text-sm mt-1">{error}</p>
      </div>
    </div>
  )

  const ventasData = ventas.map((v) => ({
    name: `${v._id?.restaurante || 'N/A'} (${v._id?.mes ?? ''})`,
    total: v.totalVentas,
  }))

  const mejoresData = mejores.map((r) => ({
    name: r.nombre || r._id,
    calificacion: parseFloat((r.promedioCalificacion || r.promedio || 0).toFixed(1)),
  }))

  return (
    <div className="p-6 max-w-7xl mx-auto">
      <h1 className="text-2xl font-bold text-gray-800 mb-6">Dashboard</h1>

      {/* Stats */}
      <div className="grid grid-cols-2 lg:grid-cols-4 gap-4 mb-6">
        <StatCard label="Restaurantes" value={stats.restaurantes} color="border-orange-500" />
        <StatCard label="Órdenes" value={stats.ordenes} color="border-blue-500" />
        <StatCard label="Usuarios" value={stats.usuarios} color="border-green-500" />
        <StatCard label="Reseñas" value={stats.resenas} color="border-purple-500" />
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        {/* Ventas por mes */}
        <div className="bg-white rounded-xl shadow p-5">
          <h2 className="text-lg font-semibold text-gray-700 mb-4">Ventas por Mes</h2>
          {ventasData.length === 0 ? (
            <p className="text-gray-400 text-sm">Sin datos aún</p>
          ) : (
            <ResponsiveContainer width="100%" height={260}>
              <BarChart data={ventasData}>
                <CartesianGrid strokeDasharray="3 3" />
                <XAxis dataKey="name" tick={{ fontSize: 10 }} />
                <YAxis />
                <Tooltip formatter={(v) => `$${Number(v).toFixed(2)}`} />
                <Bar dataKey="total" fill="#ea580c" radius={[4, 4, 0, 0]} name="Total" />
              </BarChart>
            </ResponsiveContainer>
          )}
        </div>

        {/* Mejores restaurantes */}
        <div className="bg-white rounded-xl shadow p-5">
          <h2 className="text-lg font-semibold text-gray-700 mb-4">Mejores Restaurantes</h2>
          {mejoresData.length === 0 ? (
            <p className="text-gray-400 text-sm">Sin datos aún</p>
          ) : (
            <ResponsiveContainer width="100%" height={260}>
              <BarChart data={mejoresData} layout="vertical">
                <CartesianGrid strokeDasharray="3 3" />
                <XAxis type="number" domain={[0, 5]} />
                <YAxis dataKey="name" type="category" tick={{ fontSize: 11 }} width={120} />
                <Tooltip />
                <Legend />
                <Bar dataKey="calificacion" fill="#fb923c" radius={[0, 4, 4, 0]} name="Calificación" />
              </BarChart>
            </ResponsiveContainer>
          )}
        </div>
      </div>
    </div>
  )
}
