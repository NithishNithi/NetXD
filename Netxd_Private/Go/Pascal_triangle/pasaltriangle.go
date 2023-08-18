package main

import "fmt"

func pascaltriangle(rows int){
	var ans int
	for i:=0;i<rows;i++{
		for j:=0;j<=i;j++{
			if(j==0 || i==0){
				ans = 1
			} else{
				ans = ans *(i-j+1)/j
			}
			fmt.Print(ans)
		}
		fmt.Println()
	}
}

func main(){
	var row int
	fmt.Scanln(&row)
	pascaltriangle(row)
}