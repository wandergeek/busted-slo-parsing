package main

import (
	"encoding/json"
	"fmt"

	sloModel "busted-slo-parsing/generated/slo"
)

func main() {
	var slo sloModel.SloResponse

	// from https://github.com/elastic/kibana/blob/main/x-pack/plugins/observability/dev_docs/slo.md#availability
	data := []byte(`{
		"name": "My SLO Name",
		"description": "My SLO Description",
		"indicator": {
			"type": "sli.apm.transactionErrorRate",
			"params": {
				"environment": "production",
				"service": "o11y-app",
				"transactionType": "request",
				"transactionName": "GET /api",
				"index": "metrics-apm*"
			}
		},
		"timeWindow": {
			"duration": "30d",
			"isRolling": true
		},
		"budgetingMethod": "occurrences",
		"objective": {
			"target": 0.99
		}
	}`)

	// data := []byte(`{
	// 	"name": "My SLO Name",
	// 	"description": "My SLO Description",
	// 	"indicator": {
	// 		"type": "sli.kql.custom",
	// 		"params": {
	// 			"index": "high-cardinality-data-fake_logs*",
	// 			"good": "latency < 300",
	// 			"total": "",
	// 			"filter": "labels.groupId: group-0",
	// 			"timestampField": "custom_timestamp"
	// 		}
	// 	},
	// 	"timeWindow": {
	// 		"duration": "7d",
	// 		"isRolling": true
	// 	},
	// 	"budgetingMethod": "occurrences",
	// 	"objective": {
	// 		"target": 0.985
	// 	}
	// }`)

	// data := []byte(`{
	// 	"name": "My SLO Name",
	// 	"description": "My SLO Description",
	// 	"indicator": {
	// 		"type": "sli.apm.transactionDuration",
	// 		"params": {
	// 			"environment": "production",
	// 			"service": "o11y-app",
	// 			"transactionType": "request",
	// 			"transactionName": "GET /api",
	// 			"threshold": 500,
	// 			"index": "metrics-apm*"
	// 		}
	// 	},
	// 	"timeWindow": {
	// 		"duration": "7d",
	// 		"isRolling": true
	// 	},
	// 	"budgetingMethod": "occurrences",
	// 	"objective": {
	// 		"target": 0.99
	// 	}
	// }`)

	// data := []byte(`{
	// 	"name": "My SLO Name",
	// 	"description": "My SLO Description",
	// 	"indicator": {
	// 		"type": "sli.metric.custom",
	// 		"params": {
	// 			"index": "high-cardinality-data-fake_stack.message_processor-*",
	// 	  "good": {
	// 		"metrics": [
	// 		  {
	// 			"name": "A",
	// 			"aggregation": "sum",
	// 			"field": "processor.processed"
	// 		  }
	// 		],
	// 		"equation": "A"
	// 	  },
	// 			"total": {
	// 		"metrics": [
	// 		  {
	// 			"name": "A",
	// 			"aggregation": "sum",
	// 			"field": "processor.accepted"
	// 		  }
	// 		],
	// 		"equation": "A"
	// 	  },
	// 			"filter": "",
	// 			"timestampField": "@timestamp"
	// 		}
	// 	},
	// 	"timeWindow": {
	// 		"duration": "7d",
	// 		"isRolling": true
	// 	},
	// 	"budgetingMethod": "occurrences",
	// 	"objective": {
	// 		"target": 0.95
	// 	}
	// }`)

	if err := json.Unmarshal(data, &slo); err != nil {
		panic(fmt.Errorf("unable to unmarshal data into slo.SloResponse: %v", err))
	}

	fmt.Printf("slo: %+v\n", slo)

}
