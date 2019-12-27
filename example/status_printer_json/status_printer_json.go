package main

import (
	"flag"
	"fmt"
	"github.com/nwpc-oper/ecflow-client-go"
	"log"
	"os"
)

func main() {
	ecflowHost := flag.String("host", "localhost", "ecflow server host")
	ecflowPort := flag.String("port", "3141", "ecflow server port")
	outputFile := flag.String("output", "", "output file, default is os.Stdout")
	flag.Parse()

	client := ecflow_client.CreateEcflowClient(*ecflowHost, *ecflowPort)
	defer client.Close()

	var err error
	target := os.Stdout
	if *outputFile != "" {
		target, err = os.Create(*outputFile)
		if err != nil {
			panic(err)
		}
	}

	ret := client.Sync()
	if ret != 0 {
		log.Fatal("sync has error")
	}

	records := client.StatusRecordsJson()
	for _, record := range records {
		fmt.Fprintf(target, "%s %s\n", record.Path, record.Status)
	}
	fmt.Fprintf(target, "%d nodes\n", len(records))
}
