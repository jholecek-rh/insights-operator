package record

import "time"

// Represents records stored in memory
type MemoryRecord struct {
	Name        string
	Fingerprint string
	At          time.Time
	Data        []byte
}

type MemoryRecords []MemoryRecord

func (r MemoryRecords) Less(i, j int) bool { return r[i].At.After(r[j].At) }
func (r MemoryRecords) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r MemoryRecords) Len() int           { return len(r) }
