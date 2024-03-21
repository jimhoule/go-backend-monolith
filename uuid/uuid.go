package uuid

import "app/uuid/services"

func GetService() services.UuidService {
	return &services.NativeUuidService{}
}