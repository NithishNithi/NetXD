package main
// ABDCEFHGIJLK
import "fmt"

func main() {
	str := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L'}
	tmp:=0
	for i := 0; i < len(str)-2; i++ {
		tmp++
		if(tmp==2){
			i+=1
			temp:=str[i]
			str[i] = str[i+1]
			str[i+1]=temp
			tmp=0
			i=i+1
		}
	}

	fmt.Println(string(str))
}
