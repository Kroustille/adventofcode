import { read } from '../shared/read'

const input_path = `${__dirname}/input`
const lines = read(input_path)

enum RockType {
  Cube = '#',
  Round = 'O'
}

interface Rock {
  type: 'cube' | 'round'
  x: number
  y: number
}

const rocks: Rock[] = []
lines.forEach((line, line_index) => {
  line.split('').forEach((character, character_index) => {
    switch(character) {
      case RockType.Cube:
        rocks.push({
          type: 'cube',
          x: character_index,
          y: line_index
        })
        break

      case RockType.Round:
        rocks.push({
          type: 'round',
          x: character_index,
          y: line_index
        })
        break
    }
  })
})

const getFinalValue = (final_rocks: Rock[]) => {
  return final_rocks.reduce((acc, rock) => {
    if (rock.type === 'cube') {
      return acc
    }

    return acc + lines.length - rock.y
  }, 0)
}

const cache: Map<string, number[]> = new Map()

const spinNorth = () => {
  for(let column = 0 ; column < lines[0].length ; column++) {
    const column_rocks = rocks.filter(rock => rock.x === column).sort((a, b) => a.y - b.y)
    // const coords = column_rocks.map(rock => rock.y).join(';')
    // const key = `north-${coords}`
    // if (cache.has(key)) {
    //   const new_columns = cache.get(key)
    //   new_columns!.forEach((y, index) => {
    //     column_rocks[index].y = y
    //   })

    //   continue
    // }
    
    for(let i = 0 ; i < column_rocks.length ; i++) {
      const rock = column_rocks[i]
      if (rock.type === 'cube') {
        continue
      }
  
      const previous_rock = column_rocks[i - 1]
      if (previous_rock) {
        rock.y = previous_rock.y + 1
      } else {
        rock.y = 0
      }
    }

    // cache.set(key, column_rocks.map(rock => rock.y))
  }
}

const spinSouth = () => {
  for(let column = 0 ; column < lines[0].length ; column++) {
    const column_rocks = rocks.filter(rock => rock.x === column).sort((a, b) => a.y - b.y)
    // const coords = column_rocks.map(rock => rock.y).join(';')
    // const key = `south-${coords}`
    // if (cache.has(key)) {
    //   const new_columns = cache.get(key)
    //   new_columns!.forEach((y, index) => {
    //     column_rocks[index].y = y
    //   })

    //   continue
    // }

    for(let i = column_rocks.length - 1 ; i >= 0 ; i--) {
      const rock = column_rocks[i]
      if (rock.type === 'cube') {
        continue
      }
  
      const previous_rock = column_rocks[i + 1]
      if (previous_rock) {
        rock.y = previous_rock.y - 1
      } else {
        rock.y = lines.length - 1
      }
    }

    // cache.set(key, column_rocks.map(rock => rock.y))
  }
}

const spinWest = () => {
  for (let row = 0 ; row < lines.length ; row ++) {
    const row_rocks = rocks.filter(rock => rock.y === row).sort((a, b) => a.x - b.x)
    // const coords = row_rocks.map(rock => rock.x).join(';')
    // const key = `west-${coords}`
    // if (cache.has(key)) {
    //   const new_columns = cache.get(key)
    //   new_columns!.forEach((x, index) => {
    //     row_rocks[index].x = x
    //   })

    //   continue
    // }

    for (let i = 0 ; i < row_rocks.length ; i++) {
      const rock = row_rocks[i]
      if (rock.type === 'cube') {
        continue
      }

      const previous_rock = row_rocks[i - 1]
      if (previous_rock) {
        rock.x = previous_rock.x + 1
      } else {
        rock.x = 0
      }
    }

    // cache.set(key, row_rocks.map(rock => rock.x))
  }
}

let retrieve_cache = 0
const spinEast = () => {
  for (let row = 0 ; row < lines.length ; row ++) {
    const row_rocks = rocks.filter(rock => rock.y === row).sort((a, b) => a.x - b.x)
    // const coords = row_rocks.map(rock => rock.x).join(';')
    // const key = `east-${coords}`
    // if (cache.has(key)) {
    //   retrieve_cache++
    //   const new_columns = cache.get(key)
    //   new_columns!.forEach((x, index) => {
    //     row_rocks[index].x = x
    //   })

    //   continue
    // }

    for (let i = row_rocks.length - 1 ; i >= 0 ; i--) {
      const rock = row_rocks[i]
      if (rock.type === 'cube') {
        continue
      }

      const previous_rock = row_rocks[i + 1]
      if (previous_rock) {
        rock.x = previous_rock.x - 1
      } else {
        rock.x = lines[0].length - 1
      }
    }

    // cache.set(key, row_rocks.map(rock => rock.x))
  }
}

const print = (rocks_to_print: Rock[]) => {
  for (let y = 0 ; y < lines.length ; y++) {
    let line = ''
    for(let x = 0 ; x < lines[0].length ; x++) {
      const rock = rocks_to_print.find(rock => rock.x === x && rock.y === y)
      if (rock) {
        if (rock.type === 'cube') {
          line += '#'
        } else {
          line += 'O'
        }
      } else {
        line += '.'
      }
    }

    console.log(line)
  }
  console.log()
}

const launchCycle = () => {
  spinNorth()
  spinWest()
  spinSouth()
  spinEast()
}

const cycles: Rock[][] = []

const areSameCycles = (a_rocks: Rock[], b_rocks: Rock[]): boolean => {
  const a_round_rocks = a_rocks.filter(rock => rock.type === 'round')
  const b_round_rocks = b_rocks.filter(rock => rock.type === 'round')
  return a_round_rocks.every(a_rock => 
    b_round_rocks.some(b_rock => a_rock.x === b_rock.x && a_rock.y === b_rock.y)
  )
}

let first_index = 0
let last_index = 0
for (let i = 0 ; i < 50000 ; i++) {
  launchCycle()
  const current_cycle = rocks.map(rock => ({...rock}))
  const saved_cycle_index = cycles.findIndex(cycle => areSameCycles(cycle, current_cycle))
  
  if (saved_cycle_index > -1) {
    first_index = saved_cycle_index
    last_index = i
    cycles.push(current_cycle)
    break
  }
  
  cycles.push(current_cycle)
} 

const getIndex = (value: number): number =>  {
  const difference = last_index - first_index
  return ((value - first_index) % difference) + first_index - 1
}

const final_index = getIndex(1000000000)
const final_cycle = cycles[final_index]

// final should be equal to 5
console.log('first', first_index, 'last', last_index, 'final',  final_index)

console.log()
for(let i = first_index ; i <= last_index; i++) {
  console.log(i, getFinalValue(cycles[i]))
}

console.log()
console.log(getFinalValue(final_cycle))