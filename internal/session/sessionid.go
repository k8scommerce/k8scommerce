package session

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"io"
	"sync/atomic"
	"time"
)

func NewSessionId() string {
	timestamp := time.Now()
	length := 12
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	binary.BigEndian.PutUint32(b[4:8], uint32(timestamp.Unix()))
	randID := randUint32()
	uint24(b[9:12], atomic.AddUint32(&randID, 1))
	return hex.EncodeToString(b)
}

func randUint32() uint32 {
	var b [4]byte
	_, err := io.ReadFull(rand.Reader, b[:])
	if err != nil {
		panic(err)
	}
	return (uint32(b[0]) << 0) | (uint32(b[1]) << 8) | (uint32(b[2]) << 16) | (uint32(b[3]) << 24)
}

func uint24(b []byte, v uint32) {
	b[0] = byte(v >> 16)
	b[1] = byte(v >> 8)
	b[2] = byte(v)
}
