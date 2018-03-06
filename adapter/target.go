package apapter

import "fmt"

type target struct {
	
}

func(this *target)Request() {
	fmt.Println("普通请求！")
}
