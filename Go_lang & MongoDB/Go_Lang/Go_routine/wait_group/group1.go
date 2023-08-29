package main
import (
	"fmt"
	"sync"
)

	func routine(i int, wg *sync.WaitGroup){
		fmt.Println("started routine", i)
		fmt.Printf("Routine ended %d \n",i)
		wg.Done()
	}


	func main(){
		no_of_routines := 3
		var wg sync.WaitGroup
		for i:=0;i<no_of_routines;i++ {
			wg.Add(1)
			go routine(i, &wg)
		}
		wg.Wait();
		fmt.Println("all routines are finished")

	}
