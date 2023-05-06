package main

import (
	"bytes"
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

func main() {
	tmp := []byte("test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~test string~")
	N := 1000000
	var start time.Time
	
	start = time.Now()
	for i := 0; i < N; i++ {
		_ = string(tmp)
	}
	fmt.Println("string:", time.Since(start))

	start = time.Now()
	for i := 0; i < N; i++ {
		_ = fmt.Sprintf("%s", tmp)
	}
	fmt.Println("Sprintf:", time.Since(start))

	start = time.Now()
	for i := 0; i < N; i++ {
		_ = bytes.NewBuffer(tmp).String()
	}
	fmt.Println("NewBuffer:", time.Since(start))

	start = time.Now()
	for i := 0; i < N; i++ {
		_ = BytesToString1(tmp)
	}
	fmt.Println("reflect:", time.Since(start))
}

func BytesToString1(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}
