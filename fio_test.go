package main

import (
	"os"
	"testing"
)

func _read(path string, bufSize int) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := make([]byte, bufSize)
	for {
		n, err := f.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func BenchmarkReadContainerFile_Buf128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := _read("/tmp/gophercolor.png", 128)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkReadContainerFile_Buf1024(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := _read("/tmp/gophercolor.png", 1024)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkReadVolumeFile_Buf128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := _read("./gophercolor.png", 128)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkReadVolumeFile_Buf1024(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := _read("./gophercolor.png", 1024)
		if err != nil {
			panic(err)
		}
	}
}
