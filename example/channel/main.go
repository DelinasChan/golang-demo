package main

import(
	"fmt"
)

func test(i int){
	fmt.Printf("@test i is %d \n",i);

}

func main(){

	ch := make(chan int,5);

	for i:=0 ; i<10; i++{

		if(i > 3 && i < 6){
			ch <- i
		}
	}
	for {
		data, ok := <-ch
		if !ok {
			break
		}

		fmt.Printf("Data: %T %d\n",data,data)
		go test(data)
	}

	fmt.Println("Process is End\n");
}

