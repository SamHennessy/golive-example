package domain

import "sync"

type Service struct {
	todoLists     map[string][]ToDoTask
	todoListMutex sync.RWMutex
}

func NewService() *Service {
	var s Service
	s.todoLists = map[string][]ToDoTask{}

	return &s
}

func (s *Service) GetToDoList(key string) []ToDoTask {
	s.todoListMutex.RLock()

	list := s.todoLists[key]

	s.todoListMutex.RUnlock()

	return list
}

func (s *Service) SetToDoList(key string, list []ToDoTask) {
	s.todoListMutex.Lock()

	s.todoLists[key] = list

	s.todoListMutex.Unlock()
}
