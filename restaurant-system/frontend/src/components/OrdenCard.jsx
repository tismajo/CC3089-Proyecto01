import { Card, Badge, Button } from 'react-bootstrap'
import { cancelarOrden, deleteOrden } from '../services/ordenes'

const estadoVariant = {
  pendiente: 'warning',
  procesando: 'primary',
  completado: 'success',
  cancelado: 'danger',
}

export default function OrdenCard({ orden, onRefresh }) {
  const handleCancelar = async () => {
    if (!window.confirm('¿Cancelar esta orden?')) return
    try {
      await cancelarOrden(orden.ID)
      onRefresh()
    } catch {
      alert('Error al cancelar la orden')
    }
  }

  const handleEliminar = async () => {
    if (!window.confirm('¿Eliminar esta orden permanentemente?')) return
    try {
      await deleteOrden(orden.ID)
      onRefresh()
    } catch {
      alert('Error al eliminar la orden')
    }
  }

  return (
    <Card className="mb-3 shadow-sm">
      <Card.Body>
        <div className="d-flex justify-content-between align-items-start mb-2">
          <div>
            <small className="text-muted font-monospace">
              #{orden.ID?.slice(-6)}
            </small>
            <div className="text-muted">
              {orden.Fecha ? new Date(orden.Fecha).toLocaleDateString() : '—'}
            </div>
          </div>

          <Badge bg={estadoVariant[orden.Estado] || 'secondary'}>
            {orden.Estado}
          </Badge>
        </div>

        <div className="mb-2">
          {orden.Items?.map((item, i) => (
            <div key={i} className="d-flex justify-content-between border-bottom py-1">
              <span>
                {item.Nombre} <small className="text-muted">x{item.Cantidad}</small>
              </span>
              <span>
                $
                {(item.Subtotal ??
                  item.PrecioUnitario * item.Cantidad
                ).toFixed(2)}
              </span>
            </div>
          ))}
        </div>

        <div className="d-flex justify-content-between align-items-center mt-2">
          <strong>${orden.Total?.toFixed(2)}</strong>

          <div className="d-flex gap-2">
            {orden.Estado !== 'cancelado' &&
              orden.Estado !== 'completado' && (
                <Button
                  size="sm"
                  variant="warning"
                  onClick={handleCancelar}
                >
                  Cancelar
                </Button>
              )}

            <Button
              size="sm"
              variant="danger"
              onClick={handleEliminar}
            >
              Eliminar
            </Button>
          </div>
        </div>
      </Card.Body>
    </Card>
  )
}
