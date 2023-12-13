import { read } from '../shared/read'

const input_path = `${__dirname}/input`
const lines = read(input_path)

interface Galaxy {
  id: number
  x: number
  y: number
}

let galaxies: Galaxy[] = []
let index = 1
lines.forEach((line, line_index) => {
  line.split('').forEach((character, character_index) => {
    if (character === '#') {
      galaxies.push({
        id: index,
        x: character_index,
        y: line_index,
      })
      index++
    }
  })
})

const empty_cols: number[] = []
for (let x = 0 ; x < lines[0].split('').length ; x++) {
  const has_galaxy = galaxies.some(galaxy => galaxy.x === x)
  if (!has_galaxy) {
    empty_cols.push(x)
  }
}

const empty_rows: number[] = []
for (let y = 0 ; y < lines.length ; y++) {
  const has_galaxy = galaxies.some(galaxy => galaxy.y === y)
  if (!has_galaxy) {
    empty_rows.push(y)
  }
}

const expansion_factor = 1000000

const expanded_galaxies: Galaxy[] = galaxies.map(galaxy => {
  const x_offset = empty_cols.reduce((acc, empty_col) => {
    if (galaxy.x > empty_col) {
      return acc + expansion_factor - 1
    }
    return acc
  }, 0)

  const y_offset = empty_rows.reduce((acc, empty_row) => {
    if (galaxy.y > empty_row) {
      return acc + expansion_factor - 1
    }

    return acc
  }, 0)

  return {
    id: galaxy.id,
    x: galaxy.x + x_offset,
    y: galaxy.y + y_offset
  }
})

const getCoupleKey = (a: Galaxy, b: Galaxy): string => {
  if (a.id < b.id) {
    return `${a.id};${b.id}`
  }

  return `${b.id};${a.id}`
}

const getDistance = (a: Galaxy, b: Galaxy): number => {
  return Math.abs(a.x - b.x) + Math.abs(a.y - b.y)
}

const distances: Record<string, number> = {}
expanded_galaxies.forEach((a) => {
  expanded_galaxies.forEach((b) => {
    if (a.id === b.id) {
      return 
    }

    const key = getCoupleKey(a, b)
    if (distances[key] !== undefined) {
      return
    }

    distances[key] = getDistance(a, b)
  })
})

const sum = Object.values(distances).reduce((acc, distance) => acc+distance, 0)
console.log(sum)