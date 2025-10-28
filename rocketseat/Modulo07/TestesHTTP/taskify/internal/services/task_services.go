package services

import "taskify/internal/store"

type TaskServices struct {
	Store store.TaskStore
}

func NewTaskServices(store store.TaskStore) *TaskServices {
	return &TaskServices{
		Store: store,
	}
}

func (s *TaskServices) CreateTask(title, description string, priority int32) (store.Task, error) {
	task, err := s.Store.CreateTask(title, description, priority)
	if err != nil {
		return store.Task{}, err
	}
	return task, nil
}

func (s *TaskServices) GetTaskById(id int32) (store.Task, error) {
	task, err := s.Store.GetTaskById(id)
	if err != nil {
		return store.Task{}, err
	}
	return task, nil
}

func (s *TaskServices) ListTask() ([]store.Task, error) {
	task, err := s.Store.ListTask()
	if err != nil {
		return []store.Task{}, err
	}

	return task, err
}

func (s *TaskServices) UpdatedTask(id int32, title, description string, priority int32) (store.Task, error) {
	task, err := s.Store.UpdateTask(id, title, description, priority)
	if err != nil {
		return store.Task{}, err
	}

	return task, nil
}

func (s *TaskServices) DeleteTask(id int32) error {
	err := s.Store.DeleteTask(id)
	if err != nil {
		return err
	}

	return nil
}
