package job

import (
	"context"
	"database/sql"

	"github.com/vivaldy22/mekar-regis-user-service/tools/queries"

	userproto "github.com/vivaldy22/mekar-regis-user-service/proto"
)

type service struct {
	db *sql.DB
}

func (s *service) GetAll(ctx context.Context, empty *userproto.Empty) (*userproto.JobList, error) {
	var jobs = new(userproto.JobList)
	rows, err := s.db.Query(queries.GET_ALL_JOBS)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var each = new(userproto.Job)
		if err := rows.Scan(&each.JobId, &each.JobName); err != nil {
			return nil, err
		}
		jobs.List = append(jobs.List, each)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return jobs, nil
}

func (s *service) GetByID(ctx context.Context, id *userproto.ID) (*userproto.Job, error) {
	var job = new(userproto.Job)
	row := s.db.QueryRow(queries.GET_JOB_BY_ID, id.Id)

	err := row.Scan(&job.JobId, &job.JobName)
	if err != nil {
		return nil, err
	}
	return job, nil
}

func NewService(db *sql.DB) userproto.JobCRUDServer {
	return &service{db}
}
