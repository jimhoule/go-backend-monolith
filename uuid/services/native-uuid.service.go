package services

import "github.com/google/uuid"

type NativeUuidService struct{}

func (nus *NativeUuidService) Generate() string {
	return uuid.NewString()
}