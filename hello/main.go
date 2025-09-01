package main
import "fmt"
const (
	PORT=3000
	host="localhost"
)
func main(){
	// var hello = "hello world"
	// name:="goland"
	// var extra string
	// extra="extra"
	// const per  = 1
	// fmt.Print(per)
	// fmt.Print(PORT)
	// fmt.Printf(host)
	// fmt.Printf((extra))
	// fmt.Printf(hello,name)

	whoAmI :=func (i interface{})  {
		switch i.(type){
		case string:
			fmt.Printf("its a string")
		default:
			fmt.Printf("other type")
		}
		
	}
	whoAmI("1")
	
}