package main

import (
	"flag"
	"fmt"
	"github.com/perillaroc/ecflow-client-go"
	"log"
)

func main() {
	ecflowHost := flag.String("host", "localhost", "ecflow server host")
	ecflowPort := flag.String("port", "3141", "ecflow server port")
	flag.Parse()

	client := ecflow_client.CreateEcflowClient(*ecflowHost, *ecflowPort)
	defer client.Close()

	ret := client.Sync()
	if ret != 0 {
		log.Fatal("sync has error")
	}

	records := client.StatusRecords()
	for _, record := range records {
		fmt.Printf("%s: [%s]\n", record.Path, record.Status)
	}
	fmt.Printf("%d nodes\n", len(records))
}
