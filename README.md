# Golang fake Google BigQuery
A local service that emulate Google BigQuery for help us to create integration testing

## How to add it in your project

### Using docker-compose

Create a new `docker-compose.yaml` or just added the `golang-fake-bigquery` service to your existing one.
```
version: '3.8'

services:
  golang-fake-bigquery:
    image: joselitofilho/golang-fake-bigquery
    container_name: fake_bigquery
    ports:
      - 9000:9000
    restart: on-failure
    environment:
      FAKE_BQ_PORT: 9000
```

### Running locally

Clone repository

Get google API discovery file
```bash
$ curl https://www.googleapis.com/discovery/v1/apis/bigquery/v2/rest > discovery.json
```

Build the project
```bash
$ go build -o golang-fake-bigquery cmd/golang-fake-bigquery/main.go
```

Run the project
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