package main

type Symbol byte

type rule struct {
    inState  state
    inSymbol Symbol

    outState state
    outSymbol Symbol
    nextMove move
}

// Rules for determining that open brackets are balanced with closed
// brackets.
var BalancedBracketRules []rule = []rule {
    {q0, '<', q0, '<', R},
    {q0, '>', q1, '-', L},
    {q1, '<', q0, '-', R},
    {q0, '-', q0, '-', R},
    {q1, '-', q1, '-', L},
    {q0, 'B', q2, 'B', L},
    {q2, '-', q2, '-', L},
    {q2, 'B', qf, 'B', R},
}
