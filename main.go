package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

func writeTo(w io.Writer, msg []byte) error {
    _, err := w.Write(msg)
    return err
}

func main() {
    buf := new(bytes.Buffer)
    if err := writeTo(buf, []byte("Happy newyear")); err != nil {
        log.Fatal(err)
    }

    fmt.Println(buf.String())
}
