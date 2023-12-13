
import { read } from '../shared/read'

const input_path = `${__dirname}/input`
const lines = read(input_path)

interface Range {
  source_start: number
  source_end: number
  transformation: number
}

interface Mapping {
  source_type: string
  destination_type: string
  ranges: Range[]
}

interface Seed {
  start: number
  range: number
}

let seeds: number[] = []
let all_seeds: Seed[] = []
const mappings: Mapping[] = []
const mappingsByDestination: Record<string, Mapping> = {}

const parse = () => {
  const splitted_first_line = lines[0]
    .split(':')[1]
    .trim()
    .split(' ')
  seeds = splitted_first_line.map(seed => Number.parseInt(seed.trim(), 10))
  splitted_first_line.forEach(seed_value => {
    if (all_seeds[all_seeds.length - 1] && !all_seeds[all_seeds.length - 1]?.range) {
      all_seeds[all_seeds.length - 1].range = Number.parseInt(seed_value.trim(), 10)
    } else {
      all_seeds.push({
        start: Number.parseInt(seed_value.trim(), 10),
        range: 0
      })
    }
  })
  lines.splice(0, 1)
  let source_type = ''
  let destination_type = ''
  let ranges: Range[] = []

  lines.forEach(line => {
    if (line === '') {
      if (source_type) {
        const mapping = {
          source_type,
          destination_type,
          ranges: ranges.sort((a, b) => a.source_start - b.source_start)
        }

        mappings.push(mapping)
        mappingsByDestination[destination_type] = mapping

        source_type = ''
        destination_type = ''
        ranges = []
      }
      return
    }

    if (line.includes('map')) {
      [source_type, destination_type] = line.split(' ')[0].split('-to-')
      return
    }

    const [destination_start, source_start, length] = line.split(' ')
    const new_range = {
      source_start: Number.parseInt(source_start, 10),
      transformation: Number.parseInt(destination_start, 10) - Number.parseInt(source_start, 10),
      source_end: Number.parseInt(source_start, 10) + Number.parseInt(length, 10)
    }
    const after_range_index = ranges.findIndex(r => r.source_start === new_range.source_end && r.transformation === new_range.transformation)
    if (after_range_index > -1) {
      ranges[after_range_index].source_start = new_range.source_start
    } else {
      const before_range_index = ranges.findIndex(r => r.source_end === new_range.source_start && r.transformation === new_range.transformation)
      if (before_range_index > -1) {
        ranges[before_range_index].source_end = new_range.source_end
      } else {
        ranges.push({
          source_start: Number.parseInt(source_start, 10),
          transformation: Number.parseInt(destination_start, 10) - Number.parseInt(source_start, 10),
          source_end: Number.parseInt(source_start, 10) + Number.parseInt(length, 10)
        })
      }
    }
  })
}


parse()

const binarySearch = (ranges: Range[], value: number): Range | undefined => {
  if (!ranges.length) {
    return undefined
  }

  const middle = Math.floor(ranges.length / 2)
  console.log(ranges.length, middle, ranges[middle])
  const range = ranges[middle]
  if (range.source_start + range.transformation <= value && value <= range.source_end + range.transformation) {
    return range
  }

  if (range.source_end + range.transformation < value) {
    return binarySearch(ranges.slice(middle + 1), value)
  }
  return binarySearch(ranges.slice(0, middle), value)
}

const findValue = (type: string, value: number): number => {
  const map = mappings.find(m => m.source_type === type)
  if (!map) {
    return value
  }

  const found_range = map.ranges.find(range => {
    return range.source_start <= value && value < range.source_end
  })

  if (found_range) {
    return findValue(map.destination_type, value + found_range.transformation)
  }

  return findValue(map.destination_type, value)
}

// const location_ids = seeds.map(seed => {
//   const value = findValue('seed', seed)
//   return value
// })

// console.log(Math.min(...location_ids))

const findReversedSeed = (destination_type: string, destination_number: number): number => {
  const map = mappingsByDestination[destination_type]
  if (!map) {
    return destination_number
  }

  const found_range = map.ranges.find(range => range.source_start + range.transformation <= destination_number && destination_number < range.source_end + range.transformation)
  if (found_range) {
    return findReversedSeed(map.source_type, destination_number - found_range.transformation)
  }

  return findReversedSeed(map.source_type, destination_number)
}

const isInSeedRange = (value: number): boolean => {
  return all_seeds.some(seed => seed.start <= value && value <= seed.start + seed.range)
}

console.time('findValues')
let location_number = 0
for (let i = 0; i < Infinity; i++) {
  const seed = findReversedSeed('location', i)
  if (isInSeedRange(seed)) {
    location_number = i
    break
  }
}

console.timeEnd('findValues')
console.log(location_number)
