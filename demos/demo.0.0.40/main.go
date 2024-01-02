package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	slurmrestapi "github.com/heromicro/slurmrestapi.go"
)

type SlurmRestConfig struct {
	Scheme string `json:"scheme,omitempty"`
	Host   string `json:"host,omitempty"`
	Port   int32  `json:"port"`
	Jwt    string `json:"jwt,omitempty"`
	User   string `json:"user,omitempty"`
}

// DSN 数据库连接串
func (a SlurmRestConfig) DSN() string {
	return fmt.Sprintf("%s://%s:%d", a.Scheme, a.Host, a.Port)
}

func (a SlurmRestConfig) Addr() string {
	return fmt.Sprintf("%s:%d", a.Host, a.Port)
}

func MakeUpSlurmApiConfig(username string, slurmConfig *SlurmRestConfig) *slurmrestapi.Configuration {

	slurmConf := slurmrestapi.NewConfiguration()
	slurmConf.HTTPClient = &http.Client{Timeout: time.Second * 3600}
	slurmConf.Scheme = slurmConfig.Scheme
	slurmConf.Host = slurmConfig.Addr()

	slurmConf.AddDefaultHeader("X-SLURM-USER-TOKEN", slurmConfig.Jwt)
	slurmConf.AddDefaultHeader("X-SLURM-USER-NAME", username)

	return slurmConf
}

func MakeUpSlurmApiClient(username string, slurmConfig *SlurmRestConfig) *slurmrestapi.APIClient {

	slurmConf := slurmrestapi.NewConfiguration()
	slurmConf.HTTPClient = &http.Client{Timeout: time.Second * 3600}
	slurmConf.Scheme = slurmConfig.Scheme
	slurmConf.Host = slurmConfig.Addr()

	slurmConf.AddDefaultHeader("X-SLURM-USER-TOKEN", slurmConfig.Jwt)
	slurmConf.AddDefaultHeader("X-SLURM-USER-NAME", username)

	client := slurmrestapi.NewAPIClient(slurmConf)

	return client
}

func slurm_ping(client *slurmrestapi.APIClient) {

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

func slurm_nodes(client *slurmrestapi.APIClient) {

	nodes, resp, err := client.SlurmAPI.SlurmV0040GetNodes(context.Background()).Execute()

	if err != nil {
		log.Fatalf("FAIL: %s", err)
	} else if resp.StatusCode != 200 {
		log.Fatalf("Invalid status code: %d\n", resp.StatusCode)
	}

	fmt.Println(" ------ ======= ")
	for _, node := range nodes.Nodes {

		fmt.Println(" ----- ======= Name: ", *node.Name, " Hostname: ", *node.Hostname, " Address: ", *node.Address, " Cpus: ", *node.Cpus, " Extra: ", node.GetExtra(), " State: ", node.State, " ")

	}

}

func slurm_partitions_get(client *slurmrestapi.APIClient) {

	partitionsResp, resp, err := client.SlurmAPI.SlurmV0040GetPartitions(context.Background()).Execute()
	if err != nil {
		log.Fatalf("FAIL: %s", err)
	} else if resp.StatusCode != 200 {

		log.Fatalf("Invalid status code: %d\n", resp.StatusCode)

	}

	fmt.Println(" ------ ======= ")
	for _, pt := range partitionsResp.Partitions {

		fmt.Println(" ----- ======= Name: ", *pt.Name, " pt.NodeSets: ", *pt.NodeSets, "  pt.Alternate: ", *pt.Alternate, "  pt.Tres.Configured: ", *pt.Tres.Configured, " *pt.Cluster: ", *pt.Cluster, " *pt.Qos.Assigned ", *pt.Qos.Assigned)

		//
	}

}

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

	// slurm_ping(client)
	// slurm_nodes(client)
	slurm_partitions_get(client)

}
