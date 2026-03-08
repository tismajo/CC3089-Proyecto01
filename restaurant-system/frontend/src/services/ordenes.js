import api from './api'

export const getOrdenes = () => api.get('/ordenes')
export const createOrden = (data) => api.post('/ordenes', data)
export const cancelarOrden = (id) => api.put(`/ordenes/${id}/cancelar`)
export const procesarPendientes = () => api.put('/ordenes/masivo')
export const deleteOrden = (id) => api.delete(`/ordenes/${id}`)
