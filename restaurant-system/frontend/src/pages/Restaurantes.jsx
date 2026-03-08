import { useEffect, useState } from 'react'
import {
  Container,
  Row,
  Col,
  Button,
  Form,
  Card,
  Alert,
  Spinner
} from 'react-bootstrap'
import {
  getRestaurantes,
  createRestaurante,
  getRestaurantesCercanos
} from '../services/restaurantes'
import RestauranteCard from '../components/RestauranteCard'

const emptyForm = {
  nombre: '',
  descripcion: '',
  categorias: '',
  estado: 'activo',
  lat: '',
  lng: ''
}

export default function Restaurantes() {
  const [restaurantes, setRestaurantes] = useState([])
  const [loading, setLoading] = useState(true)
  const [showForm, setShowForm] = useState(false)
  const [form, setForm] = useState(emptyForm)
  const [busqueda, setBusqueda] = useState({ lat: '', lng: '', dist: '1000' })
  const [error, setError] = useState(null)

  const cargar = () => {
    setLoading(true)
    getRestaurantes()
      .then((r) => setRestaurantes(r.data || []))
      .catch(() => setError('Error al cargar restaurantes'))
      .finally(() => setLoading(false))
  }

  useEffect(() => { cargar() }, [])

  const handleSubmit = async (e) => {
    e.preventDefault()

    const lat = parseFloat(form.lat)
    const lng = parseFloat(form.lng)

    if (isNaN(lat) || isNaN(lng)) {
      alert('Latitud y longitud deben ser números válidos')
      return
    }

    const payload = {
      nombre: form.nombre,
      descripcion: form.descripcion,
      categorias: form.categorias
        .split(',')
        .map((c) => c.trim())
        .filter(Boolean),
      estado: form.estado,
      ubicacion: {
        type: 'Point',
        coordinates: [lng, lat], // Mongo espera [lng, lat]
      }
    }

    try {
      await createRestaurante(payload)
      setShowForm(false)
      setForm(emptyForm)
      cargar()
    } catch (err) {
      console.error(err)
      alert('Error al crear restaurante')
    }
  }

  const handleBuscarCercanos = async (e) => {
    e.preventDefault()

    const lat = parseFloat(busqueda.lat)
    const lng = parseFloat(busqueda.lng)
    const dist = parseFloat(busqueda.dist || 1000)

    if (isNaN(lat) || isNaN(lng)) {
      alert('Latitud y longitud deben ser números válidos')
      return
    }

    try {
      const r = await getRestaurantesCercanos(lat, lng, dist)
      setRestaurantes(r.data || [])
    } catch (err) {
      console.error(err)
      alert('Error en búsqueda por cercanía')
    }
  }

  return (
    <Container className="mt-4">
      <Row className="mb-3 align-items-center">
        <Col>
          <h2>Restaurantes</h2>
        </Col>
        <Col className="text-end">
          <Button
            variant={showForm ? "outline-danger" : "primary"}
            onClick={() => setShowForm(!showForm)}
          >
            {showForm ? 'Cancelar' : '+ Nuevo'}
          </Button>
        </Col>
      </Row>

      {/* Formulario nuevo restaurante */}
      {showForm && (
        <Card className="mb-4 shadow-sm">
          <Card.Body>
            <Form onSubmit={handleSubmit}>
              <Row>
                <Col md={6}>
                  <Form.Group className="mb-3">
                    <Form.Label>Nombre</Form.Label>
                    <Form.Control
                      required
                      value={form.nombre}
                      onChange={(e) => setForm({ ...form, nombre: e.target.value })}
                    />
                  </Form.Group>
                </Col>

                <Col md={6}>
                  <Form.Group className="mb-3">
                    <Form.Label>Descripción</Form.Label>
                    <Form.Control
                      value={form.descripcion}
                      onChange={(e) => setForm({ ...form, descripcion: e.target.value })}
                    />
                  </Form.Group>
                </Col>

                <Col md={6}>
                  <Form.Group className="mb-3">
                    <Form.Label>Categorías</Form.Label>
                    <Form.Control
                      placeholder="italiana, pizza"
                      value={form.categorias}
                      onChange={(e) => setForm({ ...form, categorias: e.target.value })}
                    />
                  </Form.Group>
                </Col>

                <Col md={6}>
                  <Form.Group className="mb-3">
                    <Form.Label>Estado</Form.Label>
                    <Form.Select
                      value={form.estado}
                      onChange={(e) => setForm({ ...form, estado: e.target.value })}
                    >
                      <option value="activo">Activo</option>
                      <option value="inactivo">Inactivo</option>
                    </Form.Select>
                  </Form.Group>
                </Col>

                <Col md={6}>
                  <Form.Group className="mb-3">
                    <Form.Label>Latitud</Form.Label>
                    <Form.Control
                      value={form.lat}
                      onChange={(e) => setForm({ ...form, lat: e.target.value })}
                    />
                  </Form.Group>
                </Col>

                <Col md={6}>
                  <Form.Group className="mb-3">
                    <Form.Label>Longitud</Form.Label>
                    <Form.Control
                      value={form.lng}
                      onChange={(e) => setForm({ ...form, lng: e.target.value })}
                    />
                  </Form.Group>
                </Col>
              </Row>

              <Button type="submit" variant="success">
                Guardar Restaurante
              </Button>
            </Form>
          </Card.Body>
        </Card>
      )}

      {/* Búsqueda por cercanía */}
      <Card className="mb-4 shadow-sm">
        <Card.Body>
          <Form onSubmit={handleBuscarCercanos}>
            <Row className="align-items-end">
              <Col md={3}>
                <Form.Group>
                  <Form.Label>Latitud</Form.Label>
                  <Form.Control
                    value={busqueda.lat}
                    onChange={(e) => setBusqueda({ ...busqueda, lat: e.target.value })}
                  />
                </Form.Group>
              </Col>

              <Col md={3}>
                <Form.Group>
                  <Form.Label>Longitud</Form.Label>
                  <Form.Control
                    value={busqueda.lng}
                    onChange={(e) => setBusqueda({ ...busqueda, lng: e.target.value })}
                  />
                </Form.Group>
              </Col>

              <Col md={3}>
                <Form.Group>
                  <Form.Label>Distancia (m)</Form.Label>
                  <Form.Control
                    value={busqueda.dist}
                    onChange={(e) => setBusqueda({ ...busqueda, dist: e.target.value })}
                  />
                </Form.Group>
              </Col>

              <Col md={3} className="d-flex gap-2">
                <Button type="submit" variant="dark">
                  Buscar
                </Button>
                <Button variant="outline-secondary" onClick={cargar}>
                  Ver todos
                </Button>
              </Col>
            </Row>
          </Form>
        </Card.Body>
      </Card>

      {loading ? (
        <div className="text-center mt-4">
          <Spinner animation="border" />
        </div>
      ) : error ? (
        <Alert variant="danger">{error}</Alert>
      ) : restaurantes.length === 0 ? (
        <Alert variant="secondary">No hay restaurantes registrados.</Alert>
      ) : (
        <Row>
          {restaurantes.map((r) => (
            <Col md={6} lg={4} key={r.ID}>
              <RestauranteCard restaurante={r} />
            </Col>
          ))}
        </Row>
      )}
    </Container>
  )
}
