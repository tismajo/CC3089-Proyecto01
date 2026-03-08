import { BrowserRouter, Routes, Route } from 'react-router-dom'
import Navbar from './components/Navbar'
import Dashboard from './pages/Dashboard'
import Restaurantes from './pages/Restaurantes'
import Ordenes from './pages/Ordenes'
import Usuarios from './pages/Usuarios'
import Resenas from './pages/Resenas'

export default function App() {
  return (
    <BrowserRouter>
      <Navbar />
      <main className="min-h-screen bg-gray-100">
        <Routes>
          <Route path="/" element={<Dashboard />} />
          <Route path="/restaurantes" element={<Restaurantes />} />
          <Route path="/ordenes" element={<Ordenes />} />
          <Route path="/usuarios" element={<Usuarios />} />
          <Route path="/resenas" element={<Resenas />} />
        </Routes>
      </main>
    </BrowserRouter>
  )
}
