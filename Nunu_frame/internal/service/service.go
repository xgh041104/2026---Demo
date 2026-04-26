package service

import (
	"Nunu_frame/internal/repository"
	"Nunu_frame/pkg/jwt"
	"Nunu_frame/pkg/log"
)

type Service struct {
	logger *log.Logger
	jwt    *jwt.JWT
	tm     repository.Transaction
}

func NewService(
	tm repository.Transaction,
	logger *log.Logger,
	jwt *jwt.JWT,
) *Service {
	return &Service{
		logger: logger,
		jwt:    jwt,
		tm:     tm,
	}
}
