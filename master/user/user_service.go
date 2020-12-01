package user

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"github.com/vivaldy22/mekar-regis-user-service/tools/queries"

	userproto "github.com/vivaldy22/mekar-regis-user-service/proto"
)

type service struct {
	db *sql.DB
}

func (s *service) GetAll(ctx context.Context, empty *userproto.Empty) (*userproto.UserList, error) {
	var users = new(userproto.UserList)
	rows, err := s.db.Query(queries.GET_ALL_USERS)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var each = new(userproto.User)
		if err := rows.Scan(&each.UserId, &each.UserName, &each.UserBday, &each.UserKtp, &each.UserJob,
			&each.UserEdu, &each.UserStatus); err != nil {
			return nil, err
		}
		users.List = append(users.List, each)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (s *service) GetByID(ctx context.Context, id *userproto.ID) (*userproto.User, error) {
	var user = new(userproto.User)
	row := s.db.QueryRow(queries.GET_USER_BY_ID, id.Id)

	err := row.Scan(&user.UserId, &user.UserName, &user.UserBday, &user.UserKtp, &user.UserJob,
		&user.UserEdu, &user.UserStatus)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *service) Create(ctx context.Context, user *userproto.User) (*userproto.User, error) {
	tx, err := s.db.Begin()

	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(queries.CREATE_USER)

	if err != nil {
		return nil, err
	}

	id := uuid.New().String()
	_, err = stmt.Exec(id, user.UserName, user.UserBday, user.UserKtp, user.UserJob,
		user.UserEdu)
	if err != nil {
		return nil, tx.Rollback()
	}

	user.UserId = id
	stmt.Close()
	return user, tx.Commit()
}

func (s *service) Update(ctx context.Context, request *userproto.UserUpdateRequest) (*userproto.User, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(queries.UPDATE_USER)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(request.User.UserName, request.User.UserBday, request.User.UserKtp, request.User.UserJob,
		request.User.UserEdu, request.Id)
	if err != nil {
		return nil, tx.Rollback()
	}

	stmt.Close()
	request.User.UserId = request.Id
	return request.User, tx.Commit()
}

func (s *service) Delete(ctx context.Context, id *userproto.ID) (*userproto.Empty, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return new(userproto.Empty), err
	}

	stmt, err := tx.Prepare(queries.DELETE_USER)
	if err != nil {
		return new(userproto.Empty), err
	}

	_, err = stmt.Exec(id.Id)
	if err != nil {
		return new(userproto.Empty), tx.Rollback()
	}

	stmt.Close()
	return new(userproto.Empty), tx.Commit()
}

func NewService(db *sql.DB) userproto.UserCRUDServer {
	return &service{db}
}
