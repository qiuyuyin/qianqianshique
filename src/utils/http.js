import axios from 'axios'

// 引入axios
axios.defaults.timeout = 4000 // 超时的时间设置
axios.defaults.withCredentials = true // 允许跨域
axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded'

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

export function get(url, params = {}) {
  return new Promise((resolve, reject) => {
    axios.get(`/token/${url}`, {
      params,
    }).then(response => {
      resolve(response.data)
    }).catch(err => {
      reject(err)
    })
  })
}

export function post(url, data = {}) {
  return new Promise((resolve, reject) => {
    axios.post(url, data)
      .then(response => {
        resolve(response.data)
      }, err => {
        reject(err)
      })
  })
}
