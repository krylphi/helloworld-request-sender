package handler

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Entry log entry
type Entry struct {
	// ContentId content id
	ContentId int `json:"content_id"`
	// Timestamp unix timestamp (with nano)
	Timestamp int64 `json:"timestamp"`
	//ClientId id of the sender
	ClientId int `json:"client_id"`
}

// Marshal marshals entry to json. We do not use json.Marshal because it's slower
func (e *Entry) Marshal() []byte {
	return []byte(fmt.Sprintf(`{"text": "hello world", "content_id": %s, "client_id":%s, "timestamp": %s}`,
		strconv.Itoa(e.ContentId),
		strconv.Itoa(e.ClientId),
		strconv.FormatInt(e.Timestamp, 10),
	))
}

// EntryGen generate new entry
func EntryGen(contentId int) *Entry {
	return &Entry{
		ContentId: contentId,
		Timestamp: time.Now().UnixNano() / 1000000,
		ClientId:  rand.Intn(10) + 1,
	}
}
