package main

import (
	"fmt"
	"projectgo/internal/pkd/storage"
)

func main() {
	s := storage.NewStorage()
	s.Set("key1", "val1")
	res := s.Get("key1")
	fmt.Println(res)

}
