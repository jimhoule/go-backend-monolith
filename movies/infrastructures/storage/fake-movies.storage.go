package storage

type FakeStorge struct{}

func (fs *FakeStorge) Upload(filePath string, file []byte) ([]byte, error) {
	return file, nil
}

func (fs *FakeStorge) UploadLarge(filePath string, file []byte) ([]byte, error) {
	return file, nil
}