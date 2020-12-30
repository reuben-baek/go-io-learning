package model

import (
	"crypto/rand"
	"crypto/sha1"
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

func TestObject(t *testing.T) {
	reader := rand.Reader

	p := make([]byte, 100)
	reader.Read(p)
	fmt.Printf("%v", p)
	fmt.Println()
	hash := sha1.New()
	hash.Write(p)

	reader.Read(p)
	fmt.Printf("%v", p)
	fmt.Println()
	hash.Write(p)

	sum := hash.Sum(nil)
	fmt.Printf("%v", sum)
	fmt.Println()

	hash2 := sha1.New()
	io.CopyN(hash2, reader, 200000)
	sum2 := hash2.Sum(nil)
	fmt.Printf("%v", sum2)
	fmt.Println()
}
