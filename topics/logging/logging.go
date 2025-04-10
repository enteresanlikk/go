package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func main() {
	log.Println("standart logger")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")

	mylog.SetPrefix("newmy:")
	mylog.Println("from mylog")

	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)

	buflog.Println("hello")

	fmt.Print("from buflog:", buf.String())

	jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi info")

	myslog.Info("hello info 2", "key", "value", "age", 25)
}
