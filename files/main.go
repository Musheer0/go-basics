package main

import (
	"fmt"
	"os"
)

func handleError(e error) {
	if e != nil {
		if e.Error() != "EOF" {
			panic(e)
		}
	
	}
}
func main() {
	// name:="example.txt"
	//read a file mannually 
	//1 open
	//create a buffer of same size as file
	//file.Read(buffer_array) this will add buffer to buffer arrray
	//read the buffer using string(buffer) this will convert buffur to string
	// f, err := os.Open(name)
	// handleError(err)
	// info, err :=f.Stat()
	// handleError(err)
	// fmt.Println(info.ModTime())
	// fmt.Println(info.Name())
	// //read file
	// r,err :=os.Open(name)
	// handleError(err)
	// defer r.Close()
	// buf:=make([]byte,info.Size())
	// d,err:=r.Read(buf)
	// handleError(err)
	// for i:=0;i<len(buf);i++{
	// 	println("data",d,string(buf[i]))
	// }
	//normal way to read small file cuz these load it to mem
	// f,err:=os.ReadFile(name)
	// handleError(err)
	// fmt.Println(string(f))
	//read folders
	// dir,err :=os.Open("../")
	// handleError(err)
	// defer dir.Close()
	// f,err :=dir.ReadDir(-1)
	// handleError(err)
	// for _,fi:=range f{
	// 	fmt.Println(fi.Name(),fi.IsDir())
	// }
	//create file
	// f,err :=os.Create("golang.json")
	// handleError(err)
	// defer f.Close()
	// f.WriteString(`[{"name":"golang"}]`)
	//read from one file and update it to another file
	// s,err:=os.Open("golang.json")
	// handleError(err)
	// defer s.Close()
	// dest ,err:=os.Create("dest.json")
	// handleError(err)
	// defer dest.Close()
	// reader :=bufio.NewReader(s)
	// writer :=bufio.NewWriter(dest)
	// for {
	// 	b,err :=reader.ReadByte()
	// 	if err!=nil{
	// 		if err.Error() != "EOF" {
	// 		panic(err)
	// 	}
	// 	break
	// 	}
	// 	writer.WriteByte(b)
	// }
	// writer.Flush()
	//delete file
	os.Remove("dest.json")
	fmt.Println("program ended")
	
}