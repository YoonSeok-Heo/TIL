package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"
	"unsafe"
)

var (
	openFile, _ = os.Open("./testfile.txt")
	N           = 40000000
)

func main() {
	var start time.Time

	openFile, _ = os.Open("./testfile.txt")
	start = time.Now()
	ReadLine()
	fmt.Println("NewReader.ReadLine:", time.Since(start))
	
	openFile, _ = os.Open("./testfile.txt")
	start = time.Now()
	ReadBytes()
	fmt.Println("NewReader.ReadBytes:", time.Since(start))
	
	openFile, _ = os.Open("./testfile.txt")
	start = time.Now()
	ReadString()
	fmt.Println("NewReader.ReadString:", time.Since(start))
	
	openFile, _ = os.Open("./testfile.txt")
	start = time.Now()
	ScannerText()
	fmt.Println("NewScanner.ScannerText:", time.Since(start))

	openFile, _ = os.Open("./testfile.txt")
	start = time.Now()
	ScannerBytes()
	fmt.Println("NewScanner.ScannerBytes:", time.Since(start))

}

func ReadLine() {
	reader := bufio.NewReader(openFile)
	for i := 0; i < N; i++ {
		t, _, _ := reader.ReadLine()
		str := BytesToString(t)
		str = strings.TrimSuffix(str, "\r\n")
	}
}

func ReadBytes() {
	reader := bufio.NewReader(openFile)
	for i := 0; i < N; i++ {
		t, _ := reader.ReadBytes(10)
		str := BytesToString(t)
		str = strings.TrimSuffix(str, "\r\n")
	}
}

func ReadString() {
	reader := bufio.NewReader(openFile)
	for i := 0; i < N; i++ {
		_, _ = reader.ReadString(10)
	}
}

func ScannerText() {
	scanner := bufio.NewScanner(openFile)
	//scanner.Buffer(make([]byte, 0, 2000000), 2000000)
	for i := 0; i < N; i++ {
		scanner.Scan()
		_ = scanner.Text()
	}
}

func ScannerBytes() {
	scanner := bufio.NewScanner(openFile)
	//scanner.Buffer(make([]byte, 0, 2000000), 2000000)
	for i := 0; i < N; i++ {
		scanner.Scan()
		_ = BytesToString(scanner.Bytes())
	}
}

func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}
