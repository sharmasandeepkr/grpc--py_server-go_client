package main

import (
	"context"
	"fmt"
	"grpc/models"
	"log"
	"time"

	"google.golang.org/grpc"
)

const address = "localhost:50051"

func main() {
	for i := 0; i < 9; i++ {
		func() {
			conn, err := grpc.Dial(address, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()

			client := models.NewCalculatorClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			// fmt.Println(client.SquareRoot(context.Background(), &models.Number{Value: 36}))
			number, err := client.SquareRoot(ctx, &models.Number{Value: float32(16)})
			if err != nil {
				log.Fatalf("%v", err)
			}
			fmt.Println(number.Value)
		}()
	}

}
