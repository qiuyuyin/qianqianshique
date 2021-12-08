import { post, get } from './http'

export const getAuthorList = params => post('author/getAuthorList', params)

export const getAuthorByID = id => get(`author/getAuthorByID/${id}`)

export const getAuthorPoemCount = limit => get(`author/getAuthorPoemCount/${limit}`)
