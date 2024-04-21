// Package task provides functionality for interacting with tasks in the Gtracker application.
package task

import (
	"context"
	"fmt"
	"strings"

	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/dto"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/model"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/pkg/db"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/pkg/logger"
)

// Repository represents the repository for interacting with tasks in the database.
type Repository struct {
	db     *db.DB
	logger logger.Logger
}

// NewRepository creates a new instance of the Repository.
func NewRepository(db *db.DB, logger logger.Logger) *Repository {
	return &Repository{db: db, logger: logger}
}

// Create inserts a new task into the database.
func (t Repository) Create(ctx context.Context, goalId int, task dto.CreateTask) (int, error) {
	var id int
	query := `INSERT INTO tasks(name, description, frequency, goal_id) 
			  VALUES ($1, $2, $3, $4) RETURNING id
	`

	err := t.db.QueryRow(ctx,
		query,
		task.Name,
		task.Description,
		task.Frequency,
		goalId,
	).Scan(&id)
	if err != nil {
		t.logger.Error(err)
		return 0, t.db.ParseError(err)
	}
	return id, nil
}

// Get retrieves a list of tasks for a specific goal from the database.
func (t Repository) Get(ctx context.Context, goalId int, listParams dto.ListParams) ([]model.Task, error) {
	var tasks = []model.Task{}
	query := `
		SELECT t.id, t.name, t.description, t.frequency, t.created_at, t.updated_at, t.goal_id
		FROM tasks t
		JOIN goals g on g.id = t.goal_id
		WHERE g.id = $1
		ORDER BY t.created_at DESC 
		LIMIT $2 OFFSET $3
	`
	var limit = listParams.PageSize
	var offset = (listParams.PageID - 1) * listParams.PageSize
	rows, err := t.db.Query(ctx, query, goalId, limit, offset)
	if err != nil {
		t.logger.Error(err)
		return nil, t.db.ParseError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var r model.Task

		err := rows.Scan(
			&r.ID,
			&r.Name,
			&r.Description,
			&r.Frequency,
			&r.CreatedAt,
			&r.UpdatedAt,
			&r.GoalID,
		)
		if err != nil {
			t.logger.Error(err)
			return nil, t.db.ParseError(err)
		}

		tasks = append(tasks, r)
	}

	if err := rows.Err(); err != nil {
		t.logger.Error(err)
		return nil, t.db.ParseError(err)
	}
	return tasks, err
}

// GetByID retrieves a task by its ID from the database.
func (t Repository) GetByID(ctx context.Context, taskId int) (model.Task, error) {
	var task model.Task
	query := `
			SELECT t.id, name, description, frequency, created_at, updated_at, goal_id 
			FROM tasks t
			WHERE t.id = $1
			`
	err := t.db.QueryRow(ctx, query, taskId).Scan(
		&task.ID,
		&task.Name,
		&task.Description,
		&task.Frequency,
		&task.CreatedAt,
		&task.UpdatedAt,
		&task.GoalID,
	)
	if err != nil {
		t.logger.Error(err)
		return model.Task{}, t.db.ParseError(err)
	}
	return task, nil
}

// UpdateByID updates a task by its ID in the database.
func (t Repository) UpdateByID(ctx context.Context, taskId int, task dto.UpdateTask) error {
	var setValues []string
	var args []any
	var argID = 1

	if task.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argID))
		args = append(args, *task.Name)
		argID++
	}

	if task.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argID))
		args = append(args, *task.Description)
		argID++
	}

	if task.Frequency != nil {
		setValues = append(setValues, fmt.Sprintf("frequency=$%d", argID))
		args = append(args, *task.Frequency)
		argID++
	}

	setValues = append(setValues, fmt.Sprintf("updated_at=now()"))
	updatedFields := strings.Join(setValues, ", ")
	updatedFields = fmt.Sprintf("%s WHERE id = %d", updatedFields, taskId)
	query := fmt.Sprintf("UPDATE tasks SET %s", updatedFields)

	t.logger.Printf("query: %s, args: %s", query, args)
	_, err := t.db.Exec(ctx, query, args...)
	if err != nil {
		t.logger.Error(err)
		return t.db.ParseError(err)
	}
	return nil
}

// DeleteByID deletes a task by its ID from the database.
func (t Repository) DeleteByID(ctx context.Context, taskId int) error {
	query := `
		DELETE FROM tasks t
		WHERE t.id = $1
	`
	_, err := t.db.Exec(ctx, query, taskId)
	if err != nil {
		t.logger.Error(err)
		return t.db.ParseError(err)
	}
	return nil
}
