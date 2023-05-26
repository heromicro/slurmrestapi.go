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
	cfg.Host = "localhost:6820"

	// scontrol token username=root
	jwt := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTYzNTY0NjEsImlhdCI6MTY4NDgyMDQ2MSwic3VuIjoicm9vdCJ9.1Ex5Z4j5swtzeTbV6lWfBw9N4g4sTEVKVCk49LYW5XY"
	cfg.AddDefaultHeader("X-SLURM-USER-TOKEN", jwt)
	cfg.AddDefaultHeader("X-SLURM-USER-NAME", "root")

	client := slurmrestapi.NewAPIClient(cfg)

	jreq := client.SlurmApi.SlurmV0039Ping(context.Background())
	pings, resp, err := client.SlurmApi.SlurmV0039PingExecute(jreq)

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
