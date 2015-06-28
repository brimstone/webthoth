package main

import "time"

type DescriptionState int

const (
	Unknown DescriptionState = iota
	Offer
	Answer
)

type Description struct {
	Age     time.Time
	State   DescriptionState
	Payload string
}
