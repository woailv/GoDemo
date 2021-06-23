package main

import (
	"GoDemo/Err"
	"context"
	"fmt"
	"github.com/plutov/paypal/v4"
	"os"
)

func main() {
	// Create a client instance
	c, err := paypal.NewClient("clientID", "secretID", paypal.APIBaseSandBox)
	Err.IfPanic(err)
	c.SetLog(os.Stdout) // Set log to terminal stdout
	//accessToken, err := c.GetAccessToken(context.Background())

	refund, err := c.RefundSale(context.Background(), "", &paypal.Amount{
		Currency: "USD",
		Total:    "1",
	})
	Err.IfPanic(err)
	fmt.Println(refund)
}
