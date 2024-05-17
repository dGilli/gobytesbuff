package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

type BytesClosingBuffer struct {
    *bytes.Buffer
    io.Closer
}

func NewBytesClosingBuffer() *BytesClosingBuffer {
    return &BytesClosingBuffer{
        Buffer: new(bytes.Buffer),
    }
}

func (b *BytesClosingBuffer) Close() error {
    fmt.Println("closing")
    return nil
}

func writeTo(w io.Writer, msg []byte) error {
    _, err := w.Write(msg)
    return err
}

func main() {
    buf := NewBytesClosingBuffer()
    defer buf.Close()

    if err := writeTo(buf, []byte("Happy newyear")); err != nil {
        log.Fatal(err)

    }

    fmt.Println(buf.String())
}
