package main

import (
	"context"

	"golang.org/x/sync/semaphore"
)

type SemaphoreWG struct {
	sem *semaphore.Weighted
}

func NewSemaphoreWG() *SemaphoreWG {
	return &SemaphoreWG{
		sem: semaphore.NewWeighted(1 << 30),
	}
}

func (wg *SemaphoreWG) Add(n int64) {
	wg.sem.Acquire(context.Background(), n)
}

func (wg *SemaphoreWG) Done() {
	wg.sem.Release(1)
}

func (wg *SemaphoreWG) Wait() {
	wg.sem.Acquire(context.Background(), 1<<30)
}
