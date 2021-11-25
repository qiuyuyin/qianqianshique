import { post, get } from './http'

export const getSentenceList = params => post('sentence/getSentenceList', params)

export const getSentenceByID = id => get(`sentence/getSentenceByID/${id}`)
