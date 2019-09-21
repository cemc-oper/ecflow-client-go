package main

import (
	"fmt"
	"github.com/perillaroc/ecflow-client-go"
	"log"
)

func main() {
	// client := NewEcflowClient("10.40.143.18", "31071")
	//client := NewEcflowClient("10.40.143.16", "31067")
	client := ecflow_client.NewEcflowClient("10.40.143.18", "31067")

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
