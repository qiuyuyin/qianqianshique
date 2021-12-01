import { post, get } from './http'

export const getAuthorList = params => post('author/getAuthorList', params)

export const searchAll = params => get('all/searchAll', params)
