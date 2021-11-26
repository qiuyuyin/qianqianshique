import { post, get } from './http'

// 根据参数得到诗词数据
export const getPoemList = params => post('poem/getPoemList', params)

// 根据id和poemType查找对应的诗词数据
export const getPoemByID = (id, poemType) => get(`poem/getPoemByID/${poemType}/${id}`)
