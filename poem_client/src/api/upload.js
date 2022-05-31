import axios from './http'

const uploadFile = (file, id) => {
  const data = new FormData()
  data.append('file', file)
  console.log(file)
  console.log(data)

  return axios({
    method: 'post',
    url: `/server/upload/uploadFile/${id}`,
    data,
    headers: { 'Content-Type': 'multipart/form-data' },
  })
}
export default uploadFile
