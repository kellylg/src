package iremote

import (
	"github.com/GoCollaborate/src/artifacts/card"
)

type IDigest interface {
	Cards() map[string]card.Card
	TimeStamp() int64
	SetCards(map[string]card.Card)
	SetTimeStamp(int64)
}
