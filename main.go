package main

import (
	"fmt"
	"context"
	"github.com/Omkar-Kamat/go-pro/application"
)

func main(){
	app := application.New()

	err := app.Start(context.TODO())

	if err != nil{
		fmt.Print("failed to start app:",err)
	}
}
