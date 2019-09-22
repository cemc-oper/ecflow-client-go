# ecflow-client-go

An ecFlow client for Golang using ecFlow C++ API through SWIG.

## Install

`ecflow-client-go` needs ecFlow (and boost required by ecFlow). 

Set some environment variables before build the library.

```bash
export ECFLOW_BUILD_DIR=/some/path/to/ecflow/build
export ECFLOW_SOURCE_DIR=/some/path/to/ecflow/source
export BOOST_LIB_DIR=/some/path/to/boost/stage/lib
```

## Getting Started

```go
ecflowHost := "some host"
ecflowPort := "some port"

client := ecflow_client.CreateEcflowClient(ecflowHost, ecflowPort)
defer client.Close()

ret := client.Sync()
if ret != 0 {
	log.Fatal("sync has error")
}

records := client.StatusRecords()
for _, record := range records {
	fmt.Printf("%s: [%s]\n", record.Path, record.Status)
}
```

## License

Copyright 2019, perillaroc.

`ecflow-client-go` is licensed under [MIT License](./LICENSE.md).