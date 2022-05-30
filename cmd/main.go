package main

import (
	"fmt"
	"sync"

	chatsrv "github.com/devpablocristo/demo/chat"
	inventorysrv "github.com/devpablocristo/demo/inventory"
)

func main() {
	fmt.Println("Hello World")
	wg := sync.WaitGroup{}

	wg.Add(2)
	go inventorysrv.InventoryService(&wg)
	go chatsrv.ChatService(&wg)
	wg.Wait()
}
