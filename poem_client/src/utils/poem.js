export const handleLongContent = content => {
  let newLine
  const newContent = []
  for (let i = 0; i < content.length; i += 1) {
    if (i % 2 === 0) {
      newLine = content[i]
      if (i === content.length - 1) newContent.push(newLine)
    } else {
      newLine += content[i]
      newContent.push(newLine)
    }
  }

  return newContent
}
export const poem = () => {

}
