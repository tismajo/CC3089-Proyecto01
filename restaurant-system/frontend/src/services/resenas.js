import api from './api'

export const getResenas = () => api.get('/resenas')
export const createResena = (data) => api.post('/resenas', data)
