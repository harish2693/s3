package internal

import (
	"log"
	"sync"
	"github.com/shirou/gopsutil/mem"
	)

type BufferPool struct {
	mu sync.Mutex
	cond *sync.Cond

	numBuffers uint64
	computedMaxbuffers uint64

	pool *sync.Pool
}

const BUF_SIZE = 5 * 1024 * 1024

func maxMemToUse(buffersNow uint64) uint64 {
	m, err := mem.VirtualMemory()
	if err != nil {
		panic(err)
	}

	log.d
}