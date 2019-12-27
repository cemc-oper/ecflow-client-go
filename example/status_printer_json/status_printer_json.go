package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/nwpc-oper/ecflow-client-go"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ecflowHost := flag.String("host", "localhost", "ecflow server host")
	ecflowPort := flag.String("port", "3141", "ecflow server port")
	outputFile := flag.String("output", "", "output file, default is os.Stdout")
	isProfiling := flag.Bool("enable-profiling", false, "enable profiling")
	profilingAddress := flag.String("profiling-address", "127.0.0.1:30485", "profiling address")
	flag.Parse()

	if *isProfiling {
		fmt.Println("enable profiling...")
		go func() {
			log.Println(http.ListenAndServe(*profilingAddress, nil))
		}()
	}

	go func() {
		for {
			client := ecflow_client.CreateEcflowClient(*ecflowHost, *ecflowPort)

			var err error
			target := os.Stdout
			if *outputFile != "" {
				target, err = os.Create(*outputFile)
				if err != nil {
					client.Close()
					panic(err)
				}
			}

			ret := client.Sync()
			if ret != 0 {
				log.Fatal("sync has error")
			}

			recordsJson := client.StatusRecordsJsonChar()
			var records []ecflow_client.StatusRecord
			err = json.Unmarshal([]byte(recordsJson), &records)
			if err != nil {
				fmt.Printf("This is some error: %v\n", err)
				return
			}
			//for _, record := range records {
			//	fmt.Fprintf(target, "%s %s\n", record.Path, record.Status)
			//}
			fmt.Fprintf(target, "%d nodes\n", len(records))
			client.Close()
			time.Sleep(10 * time.Second)
		}
	}()

	select {}

}
