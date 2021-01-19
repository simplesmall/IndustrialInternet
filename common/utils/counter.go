package utils

import (
	"fmt"
	"time"
)

func CostCounter() {
	begin := time.Now()
	timeCounter:=time.Since(begin)
	fmt.Println("Run cost time : ",timeCounter)
}