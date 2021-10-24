package main

func main() {
    input := "<<><>>"
    tm := New([]Symbol(input), BalancedBracketRules)
    tm.Run()
}
