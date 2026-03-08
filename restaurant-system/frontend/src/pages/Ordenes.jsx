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
  getOrdenes,
  createOrden,
  procesarPendientes
} from '../services/ordenes'
import { getUsuarios } from '../services/usuarios'
import { getRestaurantes } from '../services/restaurantes'
import OrdenCard from '../components/OrdenCard'

const emptyItem = { nombre: '', precioUnitario: '', cantidad: 1 }

export default function Ordenes() {
  const [ordenes, setOrdenes] = useState([])
  const [usuarios, setUsuarios] = useState([])
  const [restaurantes, setRestaurantes] = useState([])
  const [loading, setLoading] = useState(true)
  const [showForm, setShowForm] = useState(false)
  const [usuarioId, setUsuarioId] = useState('')
  const [restauranteId, setRestauranteId] = useState('')
  const [items, setItems] = useState([{ ...emptyItem }])
  const [error, setError] = useState(null)

  const cargar = () => {
    setLoading(true)

    Promise.all([getOrdenes(), getUsuarios(), getRestaurantes()])
      .then(([o, u, r]) => {
        setOrdenes(o.data || [])
        setUsuarios(u.data || [])
        setRestaurantes(r.data || [])
      })
      .catch(() => setError('Error al cargar datos'))
      .finally(() => setLoading(false))
  }

  useEffect(() => { cargar() }, [])

  const updateItem = (i, field, value) => {
    const next = [...items]
    next[i] = { ...next[i], [field]: value }
    setItems(next)
  }

  const handleSubmit = async (e) => {
    e.preventDefault()

    const payload = {
      usuario_id: usuarioId,
      restaurante_id: restauranteId,
      items: items.map((it) => ({
        nombre: it.nombre,
        precio_unitario: parseFloat(it.precioUnitario),
        cantidad: parseInt(it.cantidad),
      })),
    }

    try {
      await createOrden(payload)
      setShowForm(false)
      setItems([{ ...emptyItem }])
      cargar()
    } catch {
      alert('Error al crear orden')
    }
  }

  return (
    <Container className="mt-4">
      <Row className="mb-3 align-items-center">
        <Col>
          <h2>Órdenes</h2>
        </Col>
        <Col className="text-end">
          <Button
            variant="primary"
            className="me-2"
            onClick={async () => {
              if (!window.confirm('¿Procesar todas las pendientes?')) return
              await procesarPendientes()
              cargar()
            }}
          >
            Procesar Pendientes
          </Button>

          <Button
            variant={showForm ? "outline-danger" : "success"}
            onClick={() => setShowForm(!showForm)}
          >
            {showForm ? 'Cancelar' : '+ Nueva Orden'}
          </Button>
        </Col>
      </Row>

      {showForm && (
        <Card className="mb-4 shadow-sm">
          <Card.Body>
            <Form onSubmit={handleSubmit}>
              <Row className="mb-3">
                <Col md={6}>
                  <Form.Select
                    required
                    value={usuarioId}
                    onChange={(e) => setUsuarioId(e.target.value)}
                  >
                    <option value="">Seleccionar usuario</option>
                    {usuarios.map((u) => (
                      <option key={u.id} value={u.id}>
                        {u.Nombre}
                      </option>
                    ))}
                  </Form.Select>
                </Col>

                <Col md={6}>
                  <Form.Select
                    required
                    value={restauranteId}
                    onChange={(e) => setRestauranteId(e.target.value)}
                  >
                    <option value="">Seleccionar restaurante</option>
                    {restaurantes.map((r) => (
                      <option key={r.ID} value={r.ID}>
                        {r.Nombre}
                      </option>
                    ))}
                  </Form.Select>
                </Col>
              </Row>

              {items.map((item, i) => (
                <Row key={i} className="mb-2">
                  <Col>
                    <Form.Control
                      placeholder="Nombre del plato"
                      value={item.nombre}
                      onChange={(e) => updateItem(i, 'nombre', e.target.value)}
                    />
                  </Col>
                  <Col md={2}>
                    <Form.Control
                      type="number"
                      placeholder="Precio"
                      value={item.precioUnitario}
                      onChange={(e) =>
                        updateItem(i, 'precioUnitario', e.target.value)
                      }
                    />
                  </Col>
                  <Col md={2}>
                    <Form.Control
                      type="number"
                      min="1"
                      value={item.cantidad}
                      onChange={(e) =>
                        updateItem(i, 'cantidad', e.target.value)
                      }
                    />
                  </Col>
                </Row>
              ))}

              <Button
                type="button"
                variant="outline-secondary"
                size="sm"
                onClick={() =>
                  setItems([...items, { ...emptyItem }])
                }
                className="mb-3"
              >
                + Agregar Item
              </Button>

              <div>
                <Button type="submit" variant="success">
                  Crear Orden
                </Button>
              </div>
            </Form>
          </Card.Body>
        </Card>
      )}

      {loading ? (
        <div className="text-center mt-4">
          <Spinner animation="border" />
        </div>
      ) : error ? (
        <Alert variant="danger">{error}</Alert>
      ) : ordenes.length === 0 ? (
        <Alert variant="secondary">No hay órdenes registradas.</Alert>
      ) : (
        <Row>
          {ordenes.map((o) => (
            <Col md={6} lg={4} key={o.ID}>
              <OrdenCard orden={o} onRefresh={cargar} />
            </Col>
          ))}
        </Row>
      )}
    </Container>
  )
}
