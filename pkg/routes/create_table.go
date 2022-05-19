package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/joselitofilho/golang-fake-bigquery/pkg/data"
)

type CreateTableRequest struct {
	TableReference TableReference `json:"tableReference"`
	Schema         Schema         `json:"schema"`
}

type TableReference struct {
	ProjectId string `json:"projectId"`
	DatasetId string `json:"datasetId"`
	TableId   string `json:"tableId"`
}

type Schema struct {
	Fields []data.Field `json:"fields"`
}

func (app *App) createTable(w http.ResponseWriter, r *http.Request, projectName, datasetName string) {
	decoder := json.NewDecoder(r.Body)
	var body CreateTableRequest
	err := decoder.Decode(&body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	projectName2 := body.TableReference.ProjectId
	if projectName2 != projectName {
		log.Fatalf(`Expected project name to match: "%s" != "%s "`, projectName, projectName2)
	}
	datasetName2 := body.TableReference.DatasetId
	if datasetName2 != datasetName {
		log.Fatalf(`Expected dataset name to match: "%s" != "%s "`, datasetName, datasetName2)
	}
	tableName := body.TableReference.TableId

	project, projectOk := app.projects[projectName]
	if !projectOk {
		project = data.Project{
			Datasets: map[string]data.Dataset{},
		}
		app.projects[projectName] = project
	}

	dataset, datasetOk := project.Datasets[datasetName]
	if !datasetOk {
		log.Fatalf("Dataset doesn't exist: %s", datasetName)
	}

	fieldsCopy := make([]data.Field, len(body.Schema.Fields))
	copy(fieldsCopy, body.Schema.Fields)
	dataset.Tables[tableName] = data.Table{
		Fields: fieldsCopy,
		Rows:   []map[string]interface{}{},
	}

	// Just serve the input as output
	outputJson, err := json.Marshal(body)
	if err != nil {
		log.Fatalf("Error from Marshal: %v", err)
	}
	w.Write(outputJson)
}
