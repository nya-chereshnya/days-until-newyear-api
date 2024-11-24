package service

import "time"

type Service struct{}

func New() *Service {
	return &Service{}
}

func (*Service) DaysLeft() int64 {
	date := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	duration := time.Until(date)
	return int64(duration.Hours() / 24)
}
