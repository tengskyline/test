package main


import (
"fmt"

)

type test struct {
	aa *test2
}
type test2 struct {
	c uint32
}
 var mapT  map[uint32][]*test2
func main() {
	mapT = make(map[uint32][]*test2)
	for i:=uint32(0);i<uint32(10);i++{
		if value,exist := mapT[i];exist{
			value = append(value,&test2{c:i})
		}else{
			tmp := make([]*test2,0,1)
			tmp =append(tmp,&test2{c:i})
			mapT[i] = tmp
		}
		//mapT[i] =append(mapT[i],i)
	}
	for i:=uint32(0);i<uint32(10);i++{
		if value,exist := mapT[i];exist{
			value = append(value,&test2{c:i})
		}else{
			tmp := make([]*test2,0,1)
			tmp =append(tmp,&test2{c:i})
			mapT[i] = tmp
		}
		//mapT[i] =append(mapT[i],i)
	}
	for _,v := range mapT{
		fmt.Print(len(v))
	}
	return
}
