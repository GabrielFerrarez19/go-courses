package pgstore

import (
	"context"

	"taskify/internal/store"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PGTaskStore struct {
	Queries *Queries
	Pool    *pgxpool.Pool
}

func NewPGTaskStore(pool *pgxpool.Pool) PGTaskStore {
	return PGTaskStore{
		Queries: New(pool),
		Pool:    pool,
	}
}

func (pgs *PGTaskStore) CreateTask(ctx context.Context, title string, description string, priority int32) (store.Task, error) {
	task, err := pgs.Queries.CreateTask(ctx, CreateTaskParams{
		Title:       title,
		Description: description,
		Priority:    priority,
	})
	if err != nil {
		return store.Task{}, err
	}
	return store.Task{
		Id:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Priority:    task.Priority,
		Created_at:  task.CreatedAt.Time,
		Updated_at:  task.UpdateAt.Time,
	}, nil
}

func (pgs *PGTaskStore) GetTaskById(ctx context.Context, id int32) (store.Task, error) {
	panic("not implemented") // TODO: Implement
}

func (pgs *PGTaskStore) ListTask(ctx context.Context) ([]store.Task, error) {
	panic("not implemented") // TODO: Implement
}

func (pgs *PGTaskStore) UpdateTask(ctx context.Context, id int32, title string, description string, priority int32) (store.Task, error) {
	panic("not implemented") // TODO: Implement
}

func (pgs *PGTaskStore) DeleteTask(ctx context.Context, id int32) error {
	panic("not implemented") // TODO: Implement
}
