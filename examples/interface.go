package main

var t interface {
    talk() string
}

type martian struct{}

func (m martian) talk() string {
    return "nack nack"
}

type laser int

func (l laser) talk() string {
    return strings.Repeat("pew ", int(l))
}