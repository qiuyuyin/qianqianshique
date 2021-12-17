const parseTime = str => {
  const date = new Date(str)
  const now = new Date()
  if (now.toLocaleDateString() !== date.toLocaleDateString()) {
    return `${date.getMonth() + 1}-${date.getDate()}`
  } if (now.getHours() !== date.getHours()) {
    return `${now.getHours() - date.getHours()}小时前`
  } if (now.getMinutes() !== date.getMinutes()) {
    return `${now.getMinutes() - date.getMinutes()}分钟前`
  }

  return `${now.getSeconds() - date.getSeconds()}秒前`
}

export default parseTime
