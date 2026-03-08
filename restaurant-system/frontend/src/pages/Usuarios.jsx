import { useEffect, useState } from 'react'
import {
  Container,
  Row,
  Col,
  Button,
  Table,
  Form,
  Card,
  Alert,
  Spinner,
  Badge
} from 'react-bootstrap'
import { getUsuarios, createUsuario, bulkUsuarios } from '../services/usuarios'

const emptyForm = {
  nombre: '',
  correo: '',
  contrasenaHash: '',
  direccion: '',
  roles: ''
}

export default function Usuarios() {
  const [usuarios, setUsuarios] = useState([])
  const [loading, setLoading] = useState(true)
  const [showForm, setShowForm] = useState(false)
  const [form, setForm] = useState(emptyForm)
  const [error, setError] = useState(null)

  const cargar = () => {
    setLoading(true)
    getUsuarios()
      .then((r) => setUsuarios(r.data || []))
      .catch(() => setError('Error al cargar usuarios'))
      .finally(() => setLoading(false))
  }

  useEffect(() => { cargar() }, [])

  const handleSubmit = async (e) => {
    e.preventDefault()

    const payload = {
      nombre: form.nombre,
      correo: form.correo,
      contrasenaHash: form.contrasenaHash,
      direccion: form.direccion,
      roles: form.roles
        .split(',')
        .map((r) => r.trim())
        .filter(Boolean),
    }

    try {
      await createUsuario(payload)
      setShowForm(false)
      setForm(emptyForm)
      cargar()
    } catch (err) {
      console.error(err)
      alert('Error al crear usuario')
    }
  }

  return (
    <Container className="mt-3">
      <Row className="mb-3 align-items-center">
        <Col>
          <h2>Usuarios</h2>
        </Col>
        <Col className="text-end">
          <Button
            variant="secondary"
            className="me-2"
            onClick={async () => {
              if (!window.confirm('¿Cargar usuarios de demo (bulk insert)?')) return
              try {
                await bulkUsuarios()
                cargar()
              } catch {
                alert('Error en bulk insert')
              }
            }}
          >
            Cargar Demo
          </Button>

          <Button
            variant={showForm ? "outline-danger" : "primary"}
            onClick={() => setShowForm(!showForm)}
          >
            {showForm ? 'Cancelar' : '+ Nuevo'}
          </Button>
        </Col>
      </Row>

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
                      onChange={(e) =>
                        setForm({ ...form, nombre: e.target.value })
                      }
                    />
                  </Form.Group>
                </Col>

                <Col md={6}>
                  <Form.Group className="mb-3">
                    <Form.Label>Correo</Form.Label>
                    <Form.Control
                      required
                      type="email"
                      value={form.correo}
                      onChange={(e) =>
                        setForm({ ...form, correo: e.target.value })
                      }
                    />
                  </Form.Group>
                </Col>

                <Col md={6}>
                  <Form.Group className="mb-3">
                    <Form.Label>Contraseña</Form.Label>
                    <Form.Control
                      required
                      type="password"
                      value={form.contrasenaHash}
                      onChange={(e) =>
                        setForm({ ...form, contrasenaHash: e.target.value })
                      }
                    />
                  </Form.Group>
                </Col>

                <Col md={6}>
                  <Form.Group className="mb-3">
                    <Form.Label>Dirección</Form.Label>
                    <Form.Control
                      value={form.direccion}
                      onChange={(e) =>
                        setForm({ ...form, direccion: e.target.value })
                      }
                    />
                  </Form.Group>
                </Col>

                <Col md={12}>
                  <Form.Group className="mb-3">
                    <Form.Label>Roles (separados por coma)</Form.Label>
                    <Form.Control
                      placeholder="cliente, admin"
                      value={form.roles}
                      onChange={(e) =>
                        setForm({ ...form, roles: e.target.value })
                      }
                    />
                  </Form.Group>
                </Col>
              </Row>

              <Button type="submit" variant="success">
                Guardar Usuario
              </Button>
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
      ) : usuarios.length === 0 ? (
        <Alert variant="secondary">No hay usuarios registrados.</Alert>
      ) : (
        <Card className="shadow-sm">
          <Card.Body>
            <Table striped hover responsive>
              <thead>
                <tr>
                  <th>Nombre</th>
                  <th>Correo</th>
                  <th>Dirección</th>
                  <th>Roles</th>
                  <th>Registro</th>
                </tr>
              </thead>
              <tbody>
                {usuarios.map((u) => (
                  <tr key={u.id}>
                    <td>{u.Nombre}</td>
                    <td>{u.Correo}</td>
                    <td>{u.Direccion || '-'}</td>
                    <td>
                      {u.Roles?.length
                        ? u.Roles.map((r, index) => (
                            <Badge bg="primary" key={index} className="me-1">
                              {r}
                            </Badge>
                          ))
                        : '-'}
                    </td>
                    <td>
                      {u.FechaRegistro &&
                      u.FechaRegistro !== "0001-01-01T00:00:00Z"
                        ? new Date(u.FechaRegistro).toLocaleDateString()
                        : '-'}
                    </td>
                  </tr>
                ))}
              </tbody>
            </Table>
          </Card.Body>
        </Card>
      )}
    </Container>
  )
}
