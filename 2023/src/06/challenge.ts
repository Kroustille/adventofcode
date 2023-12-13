
import { read } from '../shared/read'

const input_path = `${__dirname}/input`
const lines = read(input_path)

interface Race {
  duration: number
  record: number
}

const races: Race[] = []
const time_line = lines[0]
const distance_line = lines[1]
time_line.split(':')[1].trim().split(' ').filter(Boolean).forEach(time => {
  races.push({
    duration: Number.parseInt(time.trim(), 10),
    record: 0
  })
})

distance_line.split(':')[1].trim().split(' ').filter(Boolean).forEach((record, index) => {
  races[index].record = Number.parseInt(record, 10)
})

console.log(races)

const findWays = (race: Race): number => {
  let possible_ways = 0
  for (let i = 0; i < race.duration; i++) {
    const way = i * (race.duration - i)
    if (way > race.record) {
      possible_ways++
    }
  }

  return possible_ways
}

const all_possible_ways = races.reduce((acc, race) => {
  return acc * findWays(race)
}, 1)

console.log(all_possible_ways)