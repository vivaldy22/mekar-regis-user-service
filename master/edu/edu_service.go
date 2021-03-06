package edu

import (
	"context"
	"database/sql"

	userproto "github.com/vivaldy22/mekar-regis-user-service/proto"
	"github.com/vivaldy22/mekar-regis-user-service/tools/queries"
)

type service struct {
	db *sql.DB
}

func (s *service) GetAll(ctx context.Context, empty *userproto.Empty) (*userproto.EduList, error) {
	var edus = new(userproto.EduList)
	rows, err := s.db.Query(queries.GET_ALL_EDU)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var each = new(userproto.Edu)
		if err := rows.Scan(&each.EduId, &each.EduName); err != nil {
			return nil, err
		}
		edus.List = append(edus.List, each)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return edus, nil
}

func (s *service) GetByID(ctx context.Context, id *userproto.ID) (*userproto.Edu, error) {
	var edu = new(userproto.Edu)
	row := s.db.QueryRow(queries.GET_EDU_BY_ID, id.Id)

	err := row.Scan(&edu.EduId, &edu.EduName)
	if err != nil {
		return nil, err
	}
	return edu, nil
}

func NewService(db *sql.DB) userproto.EduCRUDServer {
	return &service{db}
}
