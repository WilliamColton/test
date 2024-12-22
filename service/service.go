package service

import "test/dao"

type Service struct {
	Dao *dao.Dao
}

func NewService(dao *dao.Dao) *Service {
	return &Service{Dao: dao}
}
