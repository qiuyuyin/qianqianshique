import { post } from './http'

export const login = params => post('base/login', params)

export const register = params => post('base/register', params)

export const postUserPoem = params => post('user/createUserPoem', params)

export const findUserPoemByPage = params => post('user/findUserPoemByPage', params)
