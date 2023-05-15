package main

import (
	"encoding/json"
	"fmt"

	sloModel "busted-slo-parsing/generated/slo"
)

func main() {

	var slo sloModel.CreateSloRequest

	//from https://github.com/elastic/kibana/blob/main/x-pack/plugins/observability/dev_docs/slo.md#latency
	data := []byte(`{
		"name": "My SLO Name",
		"description": "My SLO Description",
		"indicator": {
			"type": "sli.apm.transactionDuration",
			"params": {
				"environment": "production",
				"service": "o11y-app",
				"transactionType": "request",
				"transactionName": "GET /api",
				"threshold": 500,
				"index": "metrics-apm*"
			}
		},
		"timeWindow": {
			"duration": "7d",
			"isRolling": true
		},
		"budgetingMethod": "occurrences",
		"objective": {
			"target": 0.99
		}
	}`)

	if err := json.Unmarshal(data, &slo); err != nil {
		panic(fmt.Errorf("unable to unmarshal data into slo.CreateSloRequest: %v", err))
	}

	fmt.Printf("slo: %+v\n", slo)

}
