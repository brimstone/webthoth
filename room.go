package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

type Room struct {
	Descriptions map[string]Description
}

func NewRoom() Room {
	r := Room{}
	r.Descriptions = make(map[string]Description)
	return r
}

func (r Room) ListDescriptions() []string {
	offers := make([]string, 0)
	for k, _ := range r.Descriptions {
		if time.Since(r.Descriptions[k].Age) > time.Minute {
			delete(r.Descriptions, k)
			continue
		}
		offers = append(offers, k)
	}
	return offers
}

func (r Room) NewDescription(payload string) (string, error) {
	u := make([]byte, 8)
	_, err := rand.Read(u)
	if err != nil {
		return "", err
	}
	key := hex.EncodeToString(u)
	r.Descriptions[key] = Description{
		Age:     time.Now(),
		State:   0,
		Payload: payload,
	}
	return key, nil
}

func (r Room) GetDescription(key string) (string, error) {
	var offer Description
	var ok bool
	if offer, ok = r.Descriptions[key]; !ok {
		return "", fmt.Errorf("Key not found")
	}
	if time.Since(offer.Age) > time.Minute {
		delete(r.Descriptions, key)
		return "", fmt.Errorf("Key not found")
	}
	if offer.State == Answer {
		delete(r.Descriptions, key)
	}
	return offer.Payload, nil
}

func (r Room) AnswerDescription(key string, payload string) error {
	var offer Description
	var ok bool
	if offer, ok = r.Descriptions[key]; !ok {
		return fmt.Errorf("Key not found")
	}
	offer.Payload = payload
	offer.State = Answer
	r.Descriptions[key] = offer
	return nil
}
