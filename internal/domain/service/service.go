package service

import "db_cp_6_sem/pkg/logger"

type Client any

type IRepo interface {
	//OpenConn(request *entities.AuthRequest, ctx context.Context) (ConnDB, string, error) // Get conn, role in database and error
	IUserRepo
	ITypeRepo
	IAnalyzerRepo
	IEventRepo
	IGasRepo
	ISensorRepo
}

type Service struct {
	repo IRepo
	log  *logger.Logger
}

func NewService(repo IRepo, log *logger.Logger) *Service {
	return &Service{
		repo: repo,
		log:  log,
	}
}
