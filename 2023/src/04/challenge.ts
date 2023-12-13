import { read } from '../shared/read'

const input_path = `${__dirname}/input`
const lines = read(input_path)

interface Card {
  winning_numbers: number[]
  player_numbers: number[]
  instance: number
}

const getWinningCount = (card: Card): number => {
  return card.player_numbers.reduce((acc, player_number) => {
    const is_winning = card.winning_numbers.some(winning_number => winning_number === player_number)
    if (is_winning) {
      return acc + 1
    }

    return acc
  }, 0)
}

const cards: Card[] = lines.map(line => {
  const [winning_line, player_line] = line.split(':')[1].split('|')
  const winning_numbers = winning_line.trim().split(' ').filter(Boolean).map(value => Number.parseInt(value.trim(), 10))
  const player_numbers = player_line.trim().split(' ').filter(Boolean).map(value => Number.parseInt(value.trim(), 10))
  return {
    winning_numbers,
    player_numbers,
    instance: 1
  }
})

const first_part_sum = cards.reduce((global_acc, card) => {
  const count = getWinningCount(card)
  if (count === 0) {
    return global_acc
  }
  if (count === 1) {
    return global_acc + 1
  }

  return global_acc + Math.pow(2, count - 1) 
}, 0)

// console.log(first_part_sum)

cards.forEach((card, index) => {
  const count = getWinningCount(card)
  for (let i = 1; i <= count ; i++) {
    if (cards[index + i]) {
      cards[index + i].instance += card.instance
    }
  }
})

console.log(cards.reduce((acc, card) => acc + card.instance, 0))