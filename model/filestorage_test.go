package model

import (
	"fmt"
	"io"
	"testing"
)

func TestFileStorage_Create(t *testing.T) {

	storage := NewFileStorage()

	object := storage.Create("a.bin")

	object.Write([]byte("hello"))

}
func TestFileStorage_Open(t *testing.T) {

	storage := NewFileStorage()

	object := storage.Open("a.bin")

	buffer := make([]byte, 5)
	object.Read(buffer)
	fmt.Printf("%s", buffer)

}

func TestFileStorage_Copy(t *testing.T) {
	fileStorage := NewFileStorage()
	memStorage := NewMemoryStorage()

	file := fileStorage.Open("a.bin")
	obj := memStorage.Create("a.bin")

	io.Copy(obj, file)

	obj2 := memStorage.Open("a.bin")
	buffer := make([]byte, 5)
	obj2.Read(buffer)
	fmt.Printf("%s", buffer)
}
