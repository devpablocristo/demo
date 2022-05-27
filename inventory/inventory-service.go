package inventorysrv

import (
	"fmt"
	"sync"

	inventory "github.com/devpablocristo/interviews/b6/inventory/domain"
	mapdb "github.com/devpablocristo/interviews/b6/inventory/infrastructure/mapdb"
	muxrouter "github.com/devpablocristo/interviews/b6/inventory/infrastructure/muxrouter"
	slicedb "github.com/devpablocristo/interviews/b6/inventory/infrastructure/slicedb"
	http "github.com/devpablocristo/interviews/b6/inventory/interfaces/controllers/http"
	repository "github.com/devpablocristo/interviews/b6/inventory/interfaces/repository"
	usecases "github.com/devpablocristo/interviews/b6/inventory/usecases"
)

var (
	mapDB                             = mapdb.NewMapDB()
	sliceDB                           = slicedb.NewSliceDB()
	httpMuxRouter muxrouter.MuxRouter = *muxrouter.NewMuxRouter()
)

func init() {
	b1 := inventory.Book{
		Author: inventory.Person{
			Firstname: "J.K.",
			Lastname:  "Rowling",
		},
		Title: "Harry Potter and the Philosopher's Stone",
		Price: 45.00,
		ISBN:  "hpotter",
	}

	b2 := inventory.Book{
		Author: inventory.Person{
			Firstname: "Isaac",
			Lastname:  "Asimov",
		},
		Title: "Foundation",
		Price: 25.24,
		ISBN:  "fasimov",
	}

	mapDB.SaveBook(b1)
	mapDB.SaveBook(b2)

	sliceDB.SaveBook(b1)
	sliceDB.SaveBook(b2)

	//fmt.Println(*mapDB)

}
func InventoryService(wg *sync.WaitGroup) {
	defer wg.Done()

	inventoryControllers := getInventoryControllers()

	httpMuxRouter.POST("/inventory/add/book", inventoryControllers.Add)
	httpMuxRouter.GET("/inventory/get", inventoryControllers.GetAll)
	httpMuxRouter.SERVE(":9999")
	// fmt.Printf("\nMap\n")
	// // list
	// inventory, _ := mapDB.ListInventory()
	// for _, book := range inventory {
	// 	fmt.Println(book)
	// }

	// // get by isbn
	// b, _ := mapDB.GetBook("hpotter")
	// fmt.Println(b)

	// // delete
	// mapDB.DeleteBook("hpotter")

	// // list
	// inventory, _ = mapDB.ListInventory()
	// for _, book := range inventory {
	// 	fmt.Println(book)
	// }

	// fmt.Printf("\nSlice\n")
	// // list
	// i, _ := sliceDB.ListInventory()
	// fmt.Println(i)

	// // get by isbn
	// b, _ = sliceDB.GetBook("hpotter")
	// fmt.Println(b)

	// // delete
	// sliceDB.DeleteBook("fasimov")

	// // list
	// i, _ = sliceDB.ListInventory()
	// fmt.Println(i)

	wg.Wait()

}

func getInventoryControllers() http.HTTPInteractor {
	inventoryRepository := repository.NewRepositoryInteractor(mapDB)
	inventoryUseCases := usecases.MakeUseCasesInteractor(inventoryRepository)
	inventoryControllers := http.NewHTTPInteractor(inventoryUseCases)

	i, _ := inventoryUseCases.ListInventory()
	for _, book := range i {
		fmt.Println(book)
	}

	// inventoryRepository2 := repository.NewRepositoryInteractor(sliceDB)
	// inventoryUseCases2 := usecases.NewUseCasesInteractor(inventoryRepository2)
	// inventoryControllers2 := httpcrtls.NewHTTPInteractor(*inventoryUseCases2)

	// i2, _ := inventoryUseCases2.ListInventory()
	// for _, book := range i2 {
	// 	fmt.Println(book)
	// }

	return *inventoryControllers

}
