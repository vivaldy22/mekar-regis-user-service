package auth

import (
	"context"
	"database/sql"

	"github.com/vivaldy22/mekar-regis-user-service/tools/queries"

	"github.com/vivaldy22/mekar-regis-user-service/tools/consts"

	"github.com/vivaldy22/mekar-regis-user-service/tools/jwttoken"

	userproto "github.com/vivaldy22/mekar-regis-user-service/proto"
)

type service struct {
	db *sql.DB
}

func (s *service) GetPassword(ctx context.Context, request *userproto.LoginRequest) (*userproto.PassResponse, error) {
	var pass = new(userproto.PassResponse)
	row := s.db.QueryRow(queries.GET_PASSWORD, request.Username)

	err := row.Scan(&pass.HashedPassword)
	if err != nil {
		return nil, err
	}
	return pass, nil
}

func (s *service) GenerateToken(ctx context.Context, request *userproto.LoginRequest) (*userproto.LoginResponse, error) {
	token, err := jwttoken.JwtEncoder(request.Username, consts.CUSTOMKEY, consts.HMACADM)

	if err != nil {
		return nil, err
	}
	return &userproto.LoginResponse{
		Token: token,
	}, nil
}

func NewService(db *sql.DB) userproto.AuthRPCServer {
	return &service{db}
}
