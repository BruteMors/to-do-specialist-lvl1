package service

import "to-do-specialist-lvl1/internal/domain/entity"

type dayToDoStorage interface {
	Add(date string, task string) error
	DeleteOne(date string, task string) error
	DeleteAll(date string) (int, error)
	FindAllDate(date string) ([]string, error)
	FindAll() ([]entity.DayToDo, error)
}

type DayToDoService struct {
	storage dayToDoStorage
}

func NewDayToDoService(storage dayToDoStorage) *DayToDoService {
	return &DayToDoService{storage: storage}
}

func (s *DayToDoService) AddTask(date string, task string) error {
	if err := s.storage.Add(date, task); err != nil {
		return err
	}
	return nil
}

func (s *DayToDoService) DeleteTask(date string, task string) error {
	if err := s.storage.DeleteOne(date, task); err != nil {
		return err
	}
	return nil
}

func (s *DayToDoService) DeleteAllTasks(date string) (int, error) {
	count, err := s.storage.DeleteAll(date)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *DayToDoService) FindAllTaskDate(date string) ([]string, error) {
	event, err := s.storage.FindAllDate(date)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (s *DayToDoService) FindAllTasks() ([]entity.DayToDo, error) {
	events, err := s.storage.FindAll()
	if err != nil {
		return nil, err
	}
	return events, nil
}
