package view

import (
	"sync"
)

type View interface {
	Start(wg sync.WaitGroup) error
}
