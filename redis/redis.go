package main

import (
	redis "CodeLab/redis/string"
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	cmd := redis.GetSetValue("numbers", "1234567", ctx)
	fmt.Println(cmd.Result())
}
