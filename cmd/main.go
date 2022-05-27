package main

import (
	"fmt"
	"sync"

	inventorysrv "github.com/devpablocristo/interviews/b6/inventory"
)

func main() {
	fmt.Println("Hello World")
	wg := sync.WaitGroup{}

	wg.Add(1)
	go inventorysrv.InventoryService(&wg)
	wg.Wait()
}
