package front

import (
	"sync"

	"github.com/puper/ppim/hash"
)

var (
	storage = NewStorage()
)

type Storage struct {
	mu      sync.RWMutex
	buckets map[uint32]*Bucket
}

func NewStorage() *Storage {
	return &Storage{
		buckets: make(map[uint32]*Bucket),
	}
}

func (this *Storage) Put(session *Session) {
	var (
		ok     bool
		bucket *Bucket
	)
	bucketId := hash.HashBucketId(session.UserId, config.GetBucketNum())
	this.mu.Lock()
	if bucket, ok = this.buckets[bucketId]; !ok {
		bucket = NewBucket()
		this.buckets[bucketId] = bucket
	}
	this.mu.Unlock()
	bucket.Put(session)
}

func (this *Storage) Del(session *Session) {
	var (
		ok     bool
		bucket *Bucket
	)
	bucketId := hash.HashBucketId(session.UserId, config.GetBucketNum())
	this.mu.RLock()
	bucket, ok = this.buckets[bucketId]
	this.mu.RUnlock()
	if ok {
		bucket.Del(session)
	}
}

type Bucket struct {
	mu       sync.RWMutex
	sessions map[string][]*Session
}

func NewBucket() *Bucket {
	return &Bucket{
		sessions: make(map[string][]*Session),
	}
}

func (this *Bucket) Put(session *Session) {
	this.mu.Lock()
	defer this.mu.Unlock()
	this.sessions[session.UserId] = append(this.sessions[session.UserId], session)
}

func (this *Bucket) Del(session *Session) {
	var (
		sessions []*Session
		ok bool
	)
	this.mu.Lock()
	if session.Platform == "" {
		delete(this.sessions[session.UserId])
	} else {
		if sessions, ok = this.sessions[session.UserId]; ok {
			for k, v := range sessions {
				if session.Platform == "" {
				}
				if v.Platform == session.Platform {
					if session.ConnectionId == "" && session.ConnectionId == 
				}
			}
		}
	}
	this.mu.Unlock()
}

type Session struct {
	UserId       string
	Platform     string
	ConnectionId string
}
