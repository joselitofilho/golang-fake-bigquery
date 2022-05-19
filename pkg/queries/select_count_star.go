package queries

import (
	"fmt"
	"log"

	"github.com/joselitofilho/golang-fake-bigquery/pkg/data"
)

func executeSelectCountStar(projectName, datasetName, tableName string,
	projects map[string]data.Project) *data.Result {

	project, projectOk := projects[projectName]
	if !projectOk {
		log.Fatalf("Unknown project: %s", projectName)
	}

	dataset, datasetOk := project.Datasets[datasetName]
	if !datasetOk {
		log.Fatalf("Unknown dataset: %s", dataset)
	}

	table, tableOk := dataset.Tables[tableName]
	if !tableOk {
		log.Fatalf("Unknown table: %s", tableName)
	}

	numRows := len(table.Rows)
	numRowsString := fmt.Sprintf("%d", numRows)
	return &data.Result{
		Fields: []data.Field{{
			Name: "f0_",
			Type: "INTEGER",
			Mode: "NULLABLE",
		}},
		Rows: []data.ResultRow{{
			Values: []data.ResultValue{
				{
					Value: &numRowsString,
				},
			},
		}},
	}
}
