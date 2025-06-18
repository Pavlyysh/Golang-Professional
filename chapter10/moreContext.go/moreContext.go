package main

import (
	"context"
	"fmt"
)

type aKey string

func main() {
	myKey := aKey("mySecretValue")
	ctx := context.WithValue(context.Background(), myKey, "mySecretValue")

	searchKey(ctx, myKey)

	searchKey(ctx, aKey("notThere"))
	emptyCtx := context.TODO()
	searchKey(emptyCtx, aKey("notThere"))
}

func searchKey(ctx context.Context, k aKey) {
	v := ctx.Value(k)
	if v != nil {
		fmt.Println("Found value:", v)
		return
	} else {
		fmt.Println("Key not found:", k)
	}
}
