import { read } from '../shared/read'

const input_path = `${__dirname}/input`
const lines = read(input_path)

interface Hand {
  cards: string[]
  bid: number
}

const hands: Hand[] = lines.reduce((acc, line) => {
  const [card_values, bid_value] = line.split(' ')
  const hand: Hand = {
    cards: card_values.split(''),
    bid: Number.parseInt(bid_value, 10)
  }

  return [...acc, hand]
}, new Array<Hand>())

const buildOccurences = (hand: Hand): { other_cards: Record<string, number>, joker_count: number } => {
  const occurences: Record<string, number> = {}
  let joker_count = 0
  hand.cards.forEach(card => {
    if (card === 'J') {
      joker_count++
    } else {
      if (!occurences[card]) {
        occurences[card] = 1
      } else {
        occurences[card]++
      }
    }
  })

  return {
    other_cards: occurences,
    joker_count
  }
}

enum FigureLevel {
  FiveOfAKind = 6,
  FourOfAKind = 5,
  FullHouse = 4,
  ThreeOfAKind = 3,
  TwoPair = 2,
  OnePair = 1,
  HighCard = 0
}

const hasFiveOfAKind = (hand: Hand): boolean => {
  // KKKKK
  // KJJJJ
  const { other_cards, joker_count } = buildOccurences(hand)
  return Object.values(other_cards).length === 1 || joker_count === 5
}

const hasFourOfAKind = (hand: Hand): boolean => {
  // KKKK2
  // KJJJ2
  const { other_cards, joker_count } = buildOccurences(hand)
  const other_cards_count = Object.values(other_cards)
  return other_cards_count.length === 2 &&
    (other_cards_count[0] + joker_count === 4 || other_cards_count[1] + joker_count === 4)
}

const hasFullHouse = (hand: Hand): boolean => {
  const { other_cards } = buildOccurences(hand)
  return Object.values(other_cards).length === 2
}

const hasThreeOfAKind = (hand: Hand): boolean => {
  // KKK23
  // KKJ23
  // KJJ23
  // JJJ33
  // KJ1J2
  const { other_cards, joker_count } = buildOccurences(hand)
  const max_other_cards = Math.max(...Object.values(other_cards))
  return Object.values(other_cards).length === 3 && max_other_cards + joker_count === 3
}

const hasTwoPairs = (hand: Hand): boolean => {
  // KKQQ2
  // KJQQ2
  // KJJQ2
  // 32T3K
  const { other_cards } = buildOccurences(hand)
  const values = Object.values(other_cards)
  const pair_values = values.filter(value => value === 2)
  return values.length === 3 && pair_values.length === 2
}

const hasPair = (hand: Hand): boolean => {
  // KK234
  // KJ234
  const { other_cards, joker_count } = buildOccurences(hand)
  const max_other_cards = Math.max(...Object.values(other_cards))
  return max_other_cards + joker_count === 2
}

const getFigureLevel = (hand: Hand): FigureLevel => {
  if (hasFiveOfAKind(hand)) {
    return FigureLevel.FiveOfAKind
  }

  if (hasFourOfAKind(hand)) {
    return FigureLevel.FourOfAKind
  }

  if (hasFullHouse(hand)) {
    return FigureLevel.FullHouse
  }

  if (hasThreeOfAKind(hand)) {
    return FigureLevel.ThreeOfAKind
  }

  if (hasTwoPairs(hand)) {
    return FigureLevel.TwoPair
  }

  if (hasPair(hand)) {
    return FigureLevel.OnePair
  }

  return FigureLevel.HighCard
}

const getCardValue = (card: string): number => {
  switch (card) {
    case 'T':
      return 10
    case 'J':
      return 1
    case 'Q':
      return 12
    case 'K':
      return 13
    case 'A':
      return 14
    default:
      return Number.parseInt(card, 10)
  }
}

const compareCards = (left: string, right: string): number => {
  const [left_value, right_value] = [getCardValue(left), getCardValue(right)]
  if (left_value > right_value) {
    return 1
  } else if (right_value > left_value) {
    return -1
  }

  return 0
}

const compareHands = (left: Hand, right: Hand): number => {
  const [
    left_figure_level,
    right_figure_level
  ] = [
      getFigureLevel(left),
      getFigureLevel(right)
    ]

  if (left_figure_level > right_figure_level) {
    return 1
  }
  if (right_figure_level > left_figure_level) {
    return -1
  }

  for (let i = 0; i < left.cards.length; i++) {
    const card_comparison_value = compareCards(left.cards[i], right.cards[i])
    if (card_comparison_value) {
      return card_comparison_value
    }
  }

  return 0
}

const sorted_hands = hands.sort(compareHands)
const total_winnings = sorted_hands.reduce((acc, hand, index) => {
  return acc + hand.bid * (index + 1)
}, 0)

console.log(total_winnings)