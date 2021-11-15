import { post } from './http'

export const getAuthorList = params => post('author/getAuthorList', params)

export const setUser = params => post('user/add', params)
