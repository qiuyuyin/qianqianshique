import { post, get } from './http'

export const login = params => post('base/login', params)

export const dd = id => get(`base/register/${id}`)

export const register = params => post('base/register', params)
