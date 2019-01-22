package main

import "fmt"

func main() {
    exp := Multiply{Add{Number{1}, Number{2}}, Number{3}}
    mach := Machine{exp}
    mach.run()
}

type Operand interface {
    inspect() string
    reducible() bool
    reduce() Operand
}

type Number struct {
    value int
}

func (n Number) inspect() string {
    return fmt.Sprint(n.value)
}

func (n Number) reducible() bool {
    return false
}

func (n Number) reduce() Operand {
    return n
}

type Add struct {
    left, right Operand
}

func (a Add) inspect() string {
    return fmt.Sprintf("%s + %s", a.left.inspect(), a.right.inspect())
}

func (a Add) reducible() bool {
    return true
}

func (a Add) reduce() Operand {
    if a.left.reducible() {
        return Add{a.left.reduce(), a.right}
    } else if a.right.reducible() {
        return Add{a.left, a.right.reduce()}
    } else {
        var left interface{} = a.left
        var right interface{} = a.right
        return Number{left.(Number).value + right.(Number).value}
    }
}

type Multiply struct {
    left, right Operand
}

func (m Multiply) inspect() string {
    return fmt.Sprintf("%s * %s", m.left.inspect(), m.right.inspect())
}

func (m Multiply) reducible() bool {
    return true
}

func (m Multiply) reduce() Operand {
    if m.left.reducible() {
        return Multiply{m.left.reduce(), m.right}
    } else if m.right.reducible() {
        return Multiply{m.left, m.right.reduce()}
    } else {
        var left interface{} = m.left
        var right interface{} = m.right
        return Number{left.(Number).value * right.(Number).value}
    }
}

type Machine struct {
    expression Operand
}

func (m *Machine) step() {
    m.expression = m.expression.reduce()
}

func (m *Machine) run() {
    for m.expression.reducible() {
        fmt.Println(m.expression.inspect())
        m.step()
    }
}
