package entity

import "time"

type Block struct {
	Index        int
	TimeStamp    time.Time
	Data         any
	PreviousHash []byte
	Hash         []byte
}
