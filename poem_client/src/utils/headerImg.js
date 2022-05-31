import { generateFromString } from 'generate-avatar'

const getHeaderImg = (userName, headerImg) => {
  if (headerImg === '') {
    const str = generateFromString(userName)

    return `data:image/svg+xml;utf8,${str}`
  }

  return headerImg
}

export default getHeaderImg
