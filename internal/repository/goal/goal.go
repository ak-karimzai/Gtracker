package goal

import (
	"context"
	"fmt"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/dto"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/model"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/pkg/db"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/pkg/logger"
	"strings"
)

type Repository struct {
	db     *db.DB
	logger logger.Logger
}

func NewRepository(db *db.DB, logger logger.Logger) *Repository {
	return &Repository{db: db, logger: logger}
}

func (g Repository) Create(ctx context.Context, userId int, goal dto.CreateGoal) (int, error) {
	var id int
	query := `
		INSERT INTO goals(name, description, start_date, end_date, user_id)
		VALUES ($1, $2, $3, $4, $5) RETURNING id
	`
	err := g.db.QueryRow(ctx,
		query,
		goal.Name,
		goal.Description,
		goal.StartDate.ToStdDate(),
		goal.TargetDate.ToStdDate(),
		userId,
	).Scan(&id)
	if err != nil {
		g.logger.Error(err)
		return 0, g.db.ParseError(err)
	}

	return id, nil
}

func (g Repository) Get(ctx context.Context, userId int, listParams dto.ListParams) ([]model.Goal, error) {
	var goals = []model.Goal{}
	query := `
		SELECT g.id, g.name, g.description, g.completion_status, g.start_date, g.end_date, g.created_at, g.updated_at, g.user_id
		FROM goals g
		JOIN users u on u.id = g.user_id
		WHERE u.id = $1
		ORDER BY g.created_at DESC 
		LIMIT $2 OFFSET $3
	`
	var limit = listParams.PageSize
	var offset = (listParams.PageID - 1) * listParams.PageSize
	rows, err := g.db.Query(ctx, query, userId, limit, offset)
	if err != nil {
		g.logger.Error(err)
		return nil, g.db.ParseError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var goal model.Goal
		err = rows.Scan(
			&goal.ID,
			&goal.Name,
			&goal.Description,
			&goal.CompletionStatus,
			&goal.StartDate,
			&goal.TargetDate,
			&goal.CreatedAt,
			&goal.UpdatedAt,
			&goal.UserID,
		)
		if err != nil {
			g.logger.Error(err)
			return nil, g.db.ParseError(err)
		}
		goals = append(goals, goal)
	}
	if err != nil {
		g.logger.Error(err)
		return nil, g.db.ParseError(err)
	}
	return goals, nil
}

func (g Repository) GetByID(ctx context.Context, goalId int) (model.Goal, error) {
	var goal model.Goal
	query := `
		SELECT g.id, name, description, completion_status, start_date, end_date, g.created_at, g.updated_at, user_id
		FROM goals g 
		WHERE g.id = $1
	`

	err := g.db.QueryRow(ctx, query, goalId).Scan(
		&goal.ID,
		&goal.Name,
		&goal.Description,
		&goal.CompletionStatus,
		&goal.StartDate,
		&goal.TargetDate,
		&goal.CreatedAt,
		&goal.UpdatedAt,
		&goal.UserID,
	)
	if err != nil {
		g.logger.Error(err)
		return model.Goal{}, g.db.ParseError(err)
	}
	return goal, nil
}

func (g Repository) UpdateByID(ctx context.Context, goalId int, update dto.UpdateGoal) error {
	var setValues []string
	var args []any
	var argID = 1

	if update.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argID))
		args = append(args, *update.Name)
		argID++
	}

	if update.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argID))
		args = append(args, *update.Description)
		argID++
	}

	if update.CompletionStatus != nil {
		setValues = append(setValues, fmt.Sprintf("completion_status=$%d", argID))
		args = append(args, *update.CompletionStatus)
		argID++
	}

	if update.TargetDate != nil {
		setValues = append(setValues, fmt.Sprintf("start_date=$%d", argID))
		args = append(args, update.StartDate.ToStdDate())
		argID++
	}

	if update.TargetDate != nil {
		setValues = append(setValues, fmt.Sprintf("end_date=$%d", argID))
		args = append(args, update.TargetDate.ToStdDate())
		argID++
	}

	setValues = append(setValues, fmt.Sprintf("updated_at=now()"))
	updatedFields := strings.Join(setValues, ", ")
	updatedFields = fmt.Sprintf("%s WHERE id = %d", updatedFields, goalId)

	query := fmt.Sprintf("UPDATE goals SET %s ", updatedFields)
	g.logger.Printf("query: %s, args: %s", query, args)
	_, err := g.db.Exec(ctx, query, args...)
	if err != nil {
		g.logger.Error(err)
		return g.db.ParseError(err)
	}
	return nil
}

func (g Repository) DeleteByID(ctx context.Context, goalId int) error {
	query := `
		DELETE FROM goals
		WHERE id = $1
	`
	_, err := g.db.Exec(ctx, query, goalId)
	if err != nil {
		g.logger.Error(err)
		return g.db.ParseError(err)
	}
	return nil
}
