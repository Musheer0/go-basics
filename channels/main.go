package main

import (
	"fmt"
)

// func processNum(numChan chan int){
//     for i :=range numChan{
// 		fmt.Println(i)
// 		time.Sleep(time.Second*2)
// 	}
// }
// func main() {
// 	mychan := make(chan int)
// 	go processNum(mychan)
// 	mychan<-5
// 	for{
// 		mychan <- rand.Intn(100)
// 	}

// }
// func add(numChan chan int , n1 int,n2 int){
// 	numChan<-n1+n2
// }
// func main(){
// 	mychan :=make(chan int)
// 	go add(mychan,4,5)
// 	res :=<-mychan
// 	fmt.Println(res)

// }

//can be used inplace of sync.WaitGroup if only singl task
// func sendMail(malichan chan bool){
// 	defer func(){malichan <- true}()
// 	fmt.Println("email sent")
// }

// func main(){
// 	emailChan :=make(chan bool)

// 	go sendMail(emailChan)
// 	<-emailChan

// }

// func sendMail(malichan chan string,done chan bool){
// 	defer func(){done<-true}()
// 	for email :=range malichan{
// 		fmt.Println(email)
// 		time.Sleep(time.Second)
// 	}
// }

// func main(){
// 	emailChan :=make(chan string,10)
// 	done:=make(chan bool)
// 	go sendMail(emailChan,done)
// 	for i:=0;i<10;i++{
// 		emailChan<- fmt.Sprintf("%d@mail.com",i+1)
// 	}
// 	fmt.Println("done sending")
// 	//this is important 
// 	close(emailChan)
// 	<-done
// }	


func main(){
	chan1 :=make(chan int)
	chan2:=make(chan string)
	go func ()  {
		chan1<-10
	}()
	go func ()  {
		chan2<-"thee"	
	}()

	for i:=0;i<2;i++{
		select{
		case val1:=<-chan1:
			fmt.Println(val1)
		case val2:=<-chan2:
			fmt.Println(val2)
		}
	}
}
//you can also set permission for channels within funtion like this

func reciveOny(em <-chan string){
	//you call only recive no send 
	//this will throw error
	em<-"eeee"
}

func sendOny(em chan<- string){
	//you call only send no recive
	//this will throw error
	<-em
}