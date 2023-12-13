import { read } from '../shared/read'

const input_path = `${__dirname}/input`
const lines = read(input_path)

interface Game {
  id: number
  max_red: number
  max_green: number
  max_blue: number
}

const parseLine = (line: string): Game => {
  const splitted = line.split(':')
  const game_id = Number.parseInt(splitted[0].split(' ')[1], 10)

  let max_red = 0
  let max_green = 0
  let max_blue = 0

  splitted[1].split(';').forEach(record => {
    const values = record.trim().split(',')
    values.forEach(value => {
      const count = Number.parseInt(value.trim().split(' ')[0], 10)
      const type = value.trim().split(' ')[1]
      switch(type) {
        case 'red':
          if (count > max_red) {
            max_red = count
          }
          break
        case 'green':
          if (count > max_green) {
            max_green = count
          }
          break
        case 'blue':
          if (count > max_blue) {
            max_blue = count
          }
          break
      }
    })
  })

  return {
    id: game_id,
    max_red,
    max_green,
    max_blue
  }
}

const parsed = lines.map(line => parseLine(line))
const sum = parsed.reduce((acc, game) => {
    if (game.max_red > 12) {
      return acc
    }
    if (game.max_green > 13) {
      return acc
    }
    if (game.max_blue > 14) {
      return acc
    }
    return acc + game.id
  }, 0)

// console.log(sum)
const total = parsed.reduce((acc, game) => {
  return acc + game.max_blue * game.max_red * game.max_green
}, 0  )

console.log(total)