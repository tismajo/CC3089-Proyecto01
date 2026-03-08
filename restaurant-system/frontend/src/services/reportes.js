import api from './api'

export const getMejoresRestaurantes = () => api.get('/reportes/mejores-restaurantes')
export const getVentasPorMes = () => api.get('/reportes/ventas-por-mes')
