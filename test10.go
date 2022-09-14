package main
import "fmt"
import "sync"

type Semaphore interface{
    Acquire()
	Release()
}

type semaphore struct{
	semc chan struct{}
}

func New(maxconcurrency int) Semaphore{
	return &semaphore{
		semc:make(chan struct{},maxconcurrency)
	}
}