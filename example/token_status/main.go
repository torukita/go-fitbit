package main

import (
	"fmt"
	"os"
	"time"

	"github.com/torukita/go-fitbit/fitbit"
)

func main() {
	access_token := os.Getenv("TEST_TOKEN")
	if access_token == "" {
		os.Exit(1)
	}
	c := fitbit.New(access_token)
	v, _, err := c.GetTokenState(access_token)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("==== token state ====")
	if v.Active {
		fmt.Println("Active: true")
		fmt.Printf("IssuedAt :%s\n", time.Unix(v.Iat/1000, 0))
		fmt.Printf("ExpiresAt:%s\n", time.Unix(v.Exp/1000, 0))
	} else {
		fmt.Println("token is not active")
	}
	//	resp.OutputWithIndent(os.Stdout)
	/*
		var buf, out bytes.Buffer
		resp.Output(&buf)
		if err := json.Indent(&out, buf.Bytes(), "", "  "); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		_, err = out.WriteTo(os.Stdout)
	*/
}
