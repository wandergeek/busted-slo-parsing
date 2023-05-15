.PHONY: generate-slo-client
generate-slo-client: ## generate Kibana slo client
	@ docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
		-i https://raw.githubusercontent.com/elastic/kibana/master/x-pack/plugins/observability/docs/openapi/slo/bundled.yaml \
		--skip-validate-spec \
		--git-repo-id terraform-provider-elasticstack \
		--git-user-id elastic \
		-p isGoSubmodule=true \
		-p packageName=slo \
		-p generateInterfaces=true \
		-g go \
		-o /local/generated/slo
	@ rm -rf generated/slo/go.mod generated/slo/go.sum generated/slo/test
	@ go fmt ./generated/...