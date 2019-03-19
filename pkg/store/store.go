package store

import "time"

type Record struct {
	Key          string
	Value        string
	ContentType  string
	CreationTime time.Time
	ModifiedTime time.Time
	Metadata     map[string]string
}

type Store struct {
	records map[string]Record
	get     <-chan getRequest
	put     <-chan putRequest
	delete  <-chan deleteRequest
}

type deleteRequest struct {
	key string
}

type getRequest struct {
	key    string
	result chan<- *Record
}

type putRequest struct {
	record   *Record
	ifAbsent bool
	result   chan<- bool
}

func (s *Store) Process() {
	for {
		select {
		case req := <-s.get:
			record := s.records[req.key]
			req.result <- &record
		case req := <-s.delete:
			delete(s.records, req.key)
		case req := <-s.put:
			if _, exists := s.records[req.record.Key]; exists && req.ifAbsent {
				req.result <- false
				continue
			}

			s.records[req.record.Key] = *req.record
			req.result <- true
		}
	}
}

func (s *Store) Put(record *Record) {

}

func (s *Store) PutIfAbsent(record *Record) (*Record, bool) {
	return nil, false
}

func (s *Store) Get(key string) (*Record, bool) {
	return nil, false
}

func (s *Store) Delete(key string) {

}
