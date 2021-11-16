import axios from 'axios'
import { get } from '../utils/http'

// 在这里注明说声明的id和密钥
const clientId = '87339f343af311eca50dbee3'
const clientSecret = 'h4FPNif7M6wqju2cfKyB9j3iDNt218VP'

export const getTokenOfPoem = () => get('writing_helper/helper/oauth/gen_token', {
  client_id: clientId,
  client_secret: clientSecret,
})

export const getAIPoem = (token, data) => {
  axios.interceptors.response.use(
    response => {
      // 如果返回的状态码为200，说明接口请求成功，可以正常拿到数据
      // 否则的话抛出错误
      if (response.status === 200) {
        return Promise.resolve(response)
      }

      return Promise.reject(response)
    },
  )
  axios.defaults.headers.post['Content-Type'] = 'application/json'
  axios.defaults.headers.post['X-Token'] = token
  axios.defaults.headers.post['X-Auth-Type'] = 'xtoken'

  return new Promise((resolve, reject) => {
    axios.post('/token/auth-proxy/ancient-poem/poem_tang', data)
      .then(response => {
        resolve(response)
      }, err => {
        reject(err)
      })
  })
}
