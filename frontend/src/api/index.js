import axios from 'axios'

const api = axios.create({
    baseURL: '/api/v1',
    timeout: 10000,
    headers: { 'Content-Type': 'application/json' }
})

export const getStylings = () => api.get('/stylings')
export const getPackages = () => api.get('/packages')
export const getPhotos = () => api.get('/photos')
export const getReviews = () => api.get('/reviews')
export const createAppointment = (data) => api.post('/appointments', data)

export default api
