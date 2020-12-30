package model

import "io"

type Storage interface {
	Create(id string) Object
	Open(id string) Object
}

type Object interface {
	io.ReadWriteCloser
}

func NewObject(id string, io io.ReadWriter) Object {
	return &object{id: id, io: io}
}

type object struct {
	id string
	io io.ReadWriter
}

func (o *object) Read(p []byte) (n int, err error) {
	return o.io.Read(p)
}

func (o *object) Write(p []byte) (n int, err error) {
	return o.io.Write(p)
}

func (o *object) Close() error {
	if c, ok := o.io.(io.Closer); ok {
		return c.Close()
	}
	return nil
}
