package main

import (
    "fmt"
    "os"
)

type state string

const (
    q0 = state("q0") // Initial state
    q1 = state("q1")
    q2 = state("q2")
    q3 = state("q3")
    qf = state("qf") // Final state
)

type move int
const (
    L move = iota
    R
)

type head struct {
    idx int
    state state
}

type TuringMachine struct {
    rules []rule
    head *head
    tape []Symbol
}

// Add blank symbol to the beginning and end of the tape.
func encloseBlankSymbol(tape []Symbol) []Symbol {
    enclosed := make([]Symbol, len(tape) + 2)
    enclosed[0], enclosed[len(tape)+1] = 'B', 'B'
    copy(enclosed[1:], tape)
    return enclosed
}

func New(tape []Symbol, rule []rule) *TuringMachine {
    tape = encloseBlankSymbol(tape)

    head := &head {
        idx: 1, // Since the tape is enclosed blank symbols, head index starts at 1.
        state: q0,
    }

    return &TuringMachine{
        rules: rule,
        head: head,
        tape: tape,
    }
}

func (tm *TuringMachine) transit() {
    if tm.isFinish() {
        fmt.Fprintln(os.Stderr, "Machine falls into finish state in transition")
        os.Exit(1)
    }

    s := tm.head.state
    c := tm.tape[tm.head.idx]

    for _, r := range tm.rules {
        if r.inState == s && r.inSymbol == c {
            tm.head.state = r.outState
            tm.tape[tm.head.idx] = r.outSymbol
            switch r.nextMove {
            case L:
                tm.head.idx -= 1
            case R:
                tm.head.idx += 1
            }

            return
        }
    }

    fmt.Fprintln(os.Stderr, "It doesn't match any rules in transition.")
    os.Exit(1)
}

func (tm TuringMachine) isFinish() bool {
    return tm.head.state == qf
}

func (tm TuringMachine) print() {
    fmt.Printf("%s", tm.tape[:tm.head.idx])
    fmt.Printf("[%s]", tm.head.state)
    fmt.Printf("%s", tm.tape[tm.head.idx:])
    fmt.Printf("\n")
}

func (tm *TuringMachine) Run() {
    fmt.Println("Start")
    for {
        tm.print()
        if tm.isFinish() {
            fmt.Println("Halt")
            break
        }
        tm.transit()
    }
}
