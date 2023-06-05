package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-database/entity"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	sqlQuery := "INSERT INTO comments(email,comment) VALUES(?,?)"
	result, err := repository.DB.ExecContext(ctx, sqlQuery, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)
	return comment, nil
}
func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	sqlQuery := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, sqlQuery, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		// exist
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		// does not exist
		return comment, errors.New("Comment with id " + strconv.Itoa(int(id)) + " does not exist!")
	}
}
func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	sqlQuery := "SELECT id, email, comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}
