package task

import (
	"sync"
	"time"
)

type Store struct {
	mu     sync.Mutex
	tasks  []Task
	nextID int
}

func NewStore() *Store {
	return &Store{nextID: 1}
}

func (s *Store) Create(title string) Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	t := Task{ID: s.nextID, Title: title, CreatedAt: time.Now()}
	s.tasks = append(s.tasks, t)
	s.nextID++

	return t
}

func (s *Store) List() []Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]Task, len(s.tasks))
	copy(out, s.tasks)
	return out
}

func (s *Store) Get(id int) (Task, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, t := range s.tasks {
		if t.ID == id {
			return t, true
		}
	}
	return Task{}, false
}