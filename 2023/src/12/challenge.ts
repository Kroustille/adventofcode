import { read } from '../shared/read'

const input_path = `${__dirname}/input`
const lines = read(input_path)

interface PlayRecord {
  springs: string[]
  values: number[]
}

enum SpringType {
  Operational = '.',
  Damaged = '#',
  Unknown = '?'
}

enum SpringState {
  Valid = 'valid',
  Invalid = 'invalid',
  Unknown = 'unknown'
}

const records: PlayRecord[] = lines.map(line => {
  const [springs, values] = line.split(' ')
  return {
    springs: springs.split(''),
    values: values.split(',').map(v => Number.parseInt(v, 10))
  }
})

const getSpringState = (record: PlayRecord): SpringState => {
  let contiguous_spring = 0
  let values_index = 0

  for (let i = 0 ; i < record.springs.length ; i++) {
    const spring = record.springs[i]
    switch(spring) {
      case SpringType.Damaged:
        contiguous_spring++
        break

      case SpringType.Operational:
        if (contiguous_spring) {
          if (contiguous_spring !== record.values[values_index]) {
            return SpringState.Invalid
          }

          contiguous_spring = 0
          values_index++
        }
        
        break

      case SpringType.Unknown:
        if (contiguous_spring > record.values[values_index]) {
          return SpringState.Invalid
        }
  
        return SpringState.Unknown
    }
  }

  if (contiguous_spring) {
    if (contiguous_spring !== record.values[values_index]) {
      return SpringState.Invalid
    }

    contiguous_spring = 0
    values_index++
  }

  if (values_index !== record.values.length) {
    return SpringState.Invalid
  }

  return SpringState.Valid
}

const getNewSpring = (spring: string[], index: number, type: SpringType): string[] => {
  const array_spring = spring
  return [...array_spring.slice(0, index), type, ...array_spring.slice(index + 1)]
}

const countPossibleArrangements = (record: PlayRecord, possible_arrangements: number): number => {
  const spring_state = getSpringState(record)
  if (spring_state === SpringState.Invalid) {
    return possible_arrangements
  }

  const unknown_spring_index = record.springs.findIndex(spring => spring === SpringType.Unknown)
  if (unknown_spring_index === -1) {
    return possible_arrangements + 1
  }

  const damaged_arrangements = countPossibleArrangements({
    springs: getNewSpring(record.springs, unknown_spring_index, SpringType.Damaged),
    values: record.values
  }, possible_arrangements)

  const operationnal_arrangements = countPossibleArrangements({
    springs: getNewSpring(record.springs, unknown_spring_index, SpringType.Operational),
    values: record.values
  }, possible_arrangements)

  return possible_arrangements + damaged_arrangements + operationnal_arrangements
}


const launchFindPossibleArrangement = (current_record: PlayRecord): number => {
  const cache: Map<string, number> = new Map()

  const findPossibleArrangement = (spring_index: number, value_index: number, contiguous_spring: number, possible_arrangements: number): number => {
    const is_end_of_spring = spring_index === current_record.springs.length
    if (is_end_of_spring) {
      const is_valid = value_index === current_record.values.length && contiguous_spring === 0 || 
        value_index === current_record.values.length - 1 && contiguous_spring === current_record.values[value_index]
      if (!is_valid) {
        return possible_arrangements
      }

      return possible_arrangements + 1
    }

    const current_spring = current_record.springs[spring_index]
    const handleOperational = (): number => {
      if (contiguous_spring && contiguous_spring !== current_record.values[value_index]) {
        return possible_arrangements
      }

      const next_spring_index = spring_index + 1
      const next_value_index = contiguous_spring ? value_index + 1 : value_index
      const key = `${next_spring_index};${next_value_index};0`
      if (cache.has(key)) {
        return cache.get(key)!
      }

      const arrangements = findPossibleArrangement(next_spring_index, next_value_index, 0, possible_arrangements)

      cache.set(key, arrangements)

      return arrangements
    }

    const handleDamaged = (): number => {
      return findPossibleArrangement(spring_index + 1, value_index, contiguous_spring + 1, possible_arrangements)
    }

    if (current_spring === SpringType.Damaged) {
      return handleDamaged()
    }

    if (current_spring === SpringType.Operational) {
      return handleOperational()
    }

    if (current_spring === SpringType.Unknown) {
      return handleOperational() + handleDamaged()
    }

    return possible_arrangements
  }

  return findPossibleArrangement(0, 0, 0, 0)
}


// console.time('countAllArrangements')
// const part_1_total = records.reduce((acc, record) => {
//   return acc + countPossibleArrangements(record, 0)
// }, 0)
// console.timeEnd('countAllArrangements')

// console.log(part_1_total)


const unfold = (record: PlayRecord): PlayRecord => {
  let unfolded_springs: string[] = []
  const values: number[] = []
  for (let i = 0 ; i < 5 ; i++) {
    unfolded_springs = unfolded_springs.concat(...record.springs, '?')
    values.push(...record.values)
  }
  
  return {
    springs: unfolded_springs.slice(0, unfolded_springs.length - 1),
    values
  }
}

// const record = records[0]
// console.log(record)
// // console.log(unfold(record))
// console.time('find arrangements')
// console.log(unfold(record).springs.join(''))
// const arrangements = findPossibleArrangement('', unfold(record), 0, 0)
// console.timeEnd('find arrangements')
// console.log(arrangements)

// console.time('count arrangements')
// const count = countPossibleArrangements(unfold(record), 0)
// console.timeEnd('count arrangements')
// console.log(count)


const unfolded_records = records.map(record => unfold(record))
console.time('findAllArrangements')
console.log('launch')
const part_2_total = unfolded_records.reduce((acc, record) => {
  const possible_arrangements = launchFindPossibleArrangement(record)
  // const possible_arrangements = countPossibleArrangements(record, 0)

  const record_springs = record.springs.join('')
  console.log(record_springs, possible_arrangements)

  return acc + possible_arrangements
}, 0)
console.timeEnd('findAllArrangements')

console.log(part_2_total)