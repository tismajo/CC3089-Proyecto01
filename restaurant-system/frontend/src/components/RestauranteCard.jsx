import { Card, Badge } from 'react-bootstrap'

export default function RestauranteCard({ restaurante }) {
  return (
    <Card className="mb-3 shadow-sm">
      <Card.Body>
        <Card.Title>{restaurante.Nombre}</Card.Title>

        {restaurante.Descripcion && (
          <Card.Text className="text-muted">
            {restaurante.Descripcion}
          </Card.Text>
        )}

        {restaurante.Categorias?.length > 0 && (
          <div className="mb-2">
            {restaurante.Categorias.map((c) => (
              <Badge bg="warning" text="dark" key={c} className="me-1">
                {c}
              </Badge>
            ))}
          </div>
        )}

        {restaurante.Estado && (
          <Badge bg={restaurante.Estado === 'activo' ? 'success' : 'danger'}>
            {restaurante.Estado}
          </Badge>
        )}
      </Card.Body>
    </Card>
  )
}
