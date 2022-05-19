# Golang fake Google BigQuery

## Build and Run

```bash
$ curl https://www.googleapis.com/discovery/v1/apis/bigquery/v2/rest > discovery.json
```
```bash
$ go build -o golang-fake-bigquery cmd/golang-fake-bigquery/main.go
```
```bash
$ ./golang-fake-bigquery -discovery-json-path discovery.json -port 9000
```

## How to use it in your Go code
```Go
cli, err := bigquery.NewClient(context.Background(), "ProjectID", option.WithoutAuthentication(), option.WithEndpoint("http://localhost:9000"))
if err != nil {
    panic(err)
}

err = c.Dataset("my_dataset").Create(context.Background(), &bigquery.DatasetMetadata{})
if err != nil {
    panic(err)
}

err = c.Dataset("my_dataset").Table("my_table").Create(context.Background(), &bigquery.TableMetadata{})
if err != nil {
    panic(err)
}

ins := c.Dataset("my_dataset").Table("my_table").Inserter()
err = ins.Put(context.Background(), []struct{ Name string }{{Name: "Joselito"}, {Name: "Viveiros"}})
if err != nil {
    panic(err)
}

query := c.Query("SELECT * FROM my_dataset.my_table")
job, err := query.Run(context.Background())
if err != nil {
    panic(err)
}

it, err := job.Read(context.Background())
if err != nil {
    panic(err)
}
// it.TotalRows == 1
```