import api from './api'

export const getRestaurantes = () => api.get('/restaurantes')
export const createRestaurante = (data) => api.post('/restaurantes', data)
export const getRestaurantesCercanos = (lat, lng, dist) =>
  api.get('/restaurantes/cercanos', { params: { lat, lng, dist } })
