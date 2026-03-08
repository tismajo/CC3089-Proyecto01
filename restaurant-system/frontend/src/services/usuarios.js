import api from './api'

export const getUsuarios = () => api.get('/usuarios')
export const createUsuario = (data) => api.post('/usuarios', data)
export const bulkUsuarios = (data) => api.post('/usuarios/bulk', data)
