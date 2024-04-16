package storage

type FakeStorge struct{}

func (fs *FakeStorge) Upload(filePath string, file []byte) (bool, error) {
	return true, nil
}

func (fs *FakeStorge) Download(filePath string) ([]byte, error) {
	return []byte{}, nil
}