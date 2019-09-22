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

	client := ecflow_client.NewEcflowClientWrapper(*ecflowHost, *ecflowPort)

	ret := client.Sync()
	if ret != 0 {
		log.Fatal("sync has error")
	}

	records := client.StatusRecords()
	count := int(records.Size())
	for i := 0; i < count; i++ {
		record := records.Get(i)
		fmt.Printf("%s: [%s]\n", record.GetPath_(), record.GetStatus_())
	}
	fmt.Printf("%d nodes\n", count)
}
