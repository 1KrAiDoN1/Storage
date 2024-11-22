package storage

type Storage struct {
	inner map[string]string
}

func NewStorage() Storage {
	return Storage{
		inner: make(map[string]string),
	}
}

func (r Storage) Set(key string, value string) {
	r.inner[key] = value
}

func (r Storage) Get(key string) string {
	return r.inner[key]
}
