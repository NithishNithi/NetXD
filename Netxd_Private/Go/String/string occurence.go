package main

import "fmt"

func main(){
	var str string = "welcome to go language go"
	var inp string
	fmt.Scanln(&inp)
	tmp:=0
	count:=0
	ans:=0
	for i := 0;i < len(str);i++{
		if(str[i] == inp[tmp]){
			count++
			tmp++
			if(count == len(inp)){
				ans++
				count = 0
				tmp =0
			}
		}
	}

	fmt.Println(ans)
}