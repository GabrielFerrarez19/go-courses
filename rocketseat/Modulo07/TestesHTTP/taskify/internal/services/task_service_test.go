package services

import (
	"testing"
	"time"

	"taskify/internal/store"

	"github.com/stretchr/testify/assert"
)

type MockTaskStore struct{}

func (mocktaskstore *MockTaskStore) CreateTask(title string, description string, priority int32) (_ store.Task, _ error) {
	return store.Task{
		Id:          1,
		Title:       title,
		Description: description,
		Priority:    priority,
		Created_at:  time.Now(),
		Updated_at:  time.Now(),
	}, nil
}

func (mocktaskstore *MockTaskStore) GetTaskById(id int32) (_ store.Task, _ error) {
	return store.Task{
		Id:          id,
		Title:       "Mock Test Task",
		Description: "Mock Test Description",
		Priority:    1,
		Created_at:  time.Now(),
		Updated_at:  time.Now(),
	}, nil
}

func (mocktaskstore *MockTaskStore) ListTask() (_ []store.Task, _ error) {
	return []store.Task{
		{
			Id:          1,
			Title:       "Mock Test Task",
			Description: "Mock Test Description",
			Priority:    1,
			Created_at:  time.Now(),
			Updated_at:  time.Now(),
		},

		{
			Id:          2,
			Title:       "Mock Test Task2",
			Description: "Mock Test Description2",
			Priority:    1,
			Created_at:  time.Now(),
			Updated_at:  time.Now(),
		},
	}, nil
}

func (mocktaskstore *MockTaskStore) UpdateTask(id int32, title string, description string, priority int32) (_ store.Task, _ error) {
	return store.Task{
		Id:          1,
		Title:       title,
		Description: description,
		Priority:    priority,
		Created_at:  time.Now(),
		Updated_at:  time.Now(),
	}, nil
}

func (mocktaskstore *MockTaskStore) DeleteTask(id int32) (_ error) {
	return nil
}

func TestCreateTask(t *testing.T) {
	mockStore := MockTaskStore{}
	TaskServices := NewTaskServices(&mockStore)
	task, err := TaskServices.Store.CreateTask("Mock Test Task", "Mock Test Description", 1)
	if err != nil {
		t.Fatal(err)
	}

	assert.NoError(t, err)
	assert.Equal(t, "Mock Test Task", task.Title)
	assert.Equal(t, "Mock Test Description", task.Description)
	assert.Equal(t, int32(1), task.Priority)
}

func TestGetTaskById(t *testing.T) {
	mockStore := MockTaskStore{}
	TaskServices := NewTaskServices(&mockStore)
	task, err := TaskServices.Store.GetTaskById(1)
	if err != nil {
		t.Fatal(err)
	}

	assert.NoError(t, err)
	assert.Equal(t, int32(1), task.Id)
	assert.Equal(t, "Mock Test Task", task.Title)
}

func TestListTasks(t *testing.T) {
	mockStore := MockTaskStore{}
	TaskServices := NewTaskServices(&mockStore)

	tasks, err := TaskServices.Store.ListTask()
	if err != nil {
		t.Fatal(err)
	}

	assert.NoError(t, err)
	assert.Len(t, tasks, 2)
	assert.Equal(t, "Mock Test Task", tasks[0].Title)
	assert.Equal(t, "Mock Test Task2", tasks[1].Title)
}
