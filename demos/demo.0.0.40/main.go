package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	slurmrestapi "github.com/heromicro/slurmrestapi.go"
)

func main() {

	cfg := slurmrestapi.NewConfiguration()
	cfg.HTTPClient = &http.Client{Timeout: time.Second * 3600}
	cfg.Scheme = "http"
	cfg.Host = "10.10.30.2:6820"

	// scontrol token username=root
	jwt := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjIwMTkyOTY0MjksImlhdCI6MTcwMzkzNjQyOSwic3VuIjoicm9vdCJ9.xCC6SLhOPmG8XxUS5MHyjcmGKGde2uSCA4P5Mw9CoyM"

	cfg.AddDefaultHeader("X-SLURM-USER-TOKEN", jwt)
	cfg.AddDefaultHeader("X-SLURM-USER-NAME", "root")

	client := slurmrestapi.NewAPIClient(cfg)

	// jreq := client.SlurmApi.SlurmV0040Ping(context.Background())
	jreq := client.SlurmAPI.SlurmV0040GetPing(context.Background())
	pings, resp, err := client.SlurmAPI.SlurmV0040GetPingExecute(jreq)

	if err != nil {
		log.Fatalf("FAIL: %s", err)
	} else if resp.StatusCode != 200 {
		log.Fatalf("Invalid status code: %d\n", resp.StatusCode)
	}

	fmt.Println(" ------ ", resp.Body)
	for _, ping := range pings.GetPings() {
		fmt.Printf("Job %s - %s\n", ping.GetPinged(), ping.GetHostname())
	}

}
