import { read } from '../shared/read'

const input_path = `${__dirname}/input`
const lines = read(input_path)

enum Play {
  ROCK = 1,
  PAPER = 2,
  SCISSOR = 3
}

const getPlays = (round: string): Play[] => {
  const [opponent_play_string, my_play_string] = round.split(' ')
  return [getOpponentPlay(opponent_play_string), getMyPlay(my_play_string)]
}

const getMyPlay = (play: string): Play => {
  switch(play) {
    case 'X':
      return Play.ROCK
    case 'Y':
      return Play.PAPER
    case 'Z':
      return Play.SCISSOR
  }

  throw new Error('unrecognized opponent play')
}

const getOpponentPlay = (play: string): Play => {
  switch(play) {
    case 'A':
      return Play.ROCK
    case 'B':
      return Play.PAPER
    case 'C':
      return Play.SCISSOR
  }

  throw new Error('unrecognized opponent play')
}

const getLosingPlay = (opponent_play: Play): Play => {
  switch(opponent_play) {
    case Play.ROCK:
      return Play.SCISSOR
    case Play.PAPER:
      return Play.ROCK
    case Play.SCISSOR:
      return Play.PAPER
  }
}

const getWinningPlay = (opponent_play: Play): Play => {
  switch(opponent_play) {
    case Play.SCISSOR:
      return Play.ROCK
    case Play.ROCK:
      return Play.PAPER
    case Play.PAPER:
      return Play.SCISSOR
  }
}

const getMyPlayStrategy = (opponent_play: Play, my_play: string): Play => {
  switch(my_play) {
    case 'X':
      return getLosingPlay(opponent_play)
    case 'Y':
      return opponent_play
    case 'Z':
    return getWinningPlay(opponent_play)
  }

  throw new Error('unrecognized strategy')
}

const getRoundScore = (opponent_play: Play, my_play: Play): number => {
  if (opponent_play === my_play) {
    return 3
  }

  if (
    opponent_play === Play.ROCK && my_play === Play.SCISSOR ||
    opponent_play === Play.SCISSOR && my_play === Play.PAPER ||
    opponent_play === Play.PAPER && my_play === Play.ROCK
  ) {
    return 0
  }

  return 6
}

const total_score_part1 = lines.reduce((score, round) => {
  if (!round) {
    return score
  }

  
  const [opponent_play, my_play] = getPlays(round)

  return score + my_play + getRoundScore(opponent_play, my_play)
}, 0)

const total_score_part2 = lines.reduce((score, round) => {
  if (!round) {
    return score
  }
  
  const [opponent_play_string, my_play_string ] = round.split(' ')
  const opponent_play = getOpponentPlay(opponent_play_string)
  const my_play = getMyPlayStrategy(opponent_play, my_play_string)

  return score + my_play + getRoundScore(opponent_play, my_play)
}, 0)

// Part 1
console.log('Part 1 =', total_score_part1)

// Part 2
console.log('Part 2 =', total_score_part2)