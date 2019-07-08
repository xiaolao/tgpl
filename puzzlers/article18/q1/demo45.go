package main

import (
	"fmt"
	"os"
	"os/exec"
)

// underlyingError会返回已知的操作系统相关错误的潜在错误值
func underlyingError(err error) error {
	switch err := err.(type) {
	case *os.PathError:
		return err.Err
	case *os.LinkError:
		return err.Err
	case *os.SyscallError:
		return err.Err
	case *exec.Error:
		return err.Err
	}
	return err
}

func main() {
	// example 1:
	// Pipe returns a connected pair of Files; reads from r return bytes written to w.
	// It returns the files and an error, if any.
	r, w, err := os.Pipe()
	if err != nil {
		fmt.Printf("unexpected error: %s\n", err)
		return
	}
	// 人为制造 *os.PathError 类型错误。
	r.Close()
	_, err = w.Write([]byte("hi"))
	uError := underlyingError(err)
	fmt.Printf("underlying error: %s (type: %T)\n", uError, uError)
	fmt.Println()

	// example 2:
	paths := []string{
		os.Args[0],
		"/it/must/not/exist",
		os.DevNull,
	}
	printError := func(i int, err error) {
		if err == nil {
			fmt.Println("nil error")
			return
		}
		err = underlyingError(err)
		switch err {
		case os.ErrClosed:
			fmt.Printf("error(closed)[%d]: %s\n", i, err)
		case os.ErrInvalid:
			fmt.Printf("error(invalid)[%d]: %s\n", i, err)
		case os.ErrPermission:
			fmt.Printf("error(permission)[%d]: %s\n", i, err)
		}
	}

	var f *os.File
	var index int
	{
		index = 0
		f, err = os.Open(paths[index])
		if err != nil {
			fmt.Printf("unexpected error: %s\n", err)
		}
		f.Close()
		_, err = f.Read([]byte{})
		printError(index, err) // error(closed)[0]: file already closed.
	}
	{
		index = 1
		f, _ = os.Open(paths[index])
		_, err = f.Stat()
		printError(index, err) // error(invalied)[1]: invalid argument
	}
	{
		index = 2
		_, err = exec.LookPath(paths[index])
		printError(index, err)
	}
	if f != nil {
		f.Close()
	}
	fmt.Println()
}
