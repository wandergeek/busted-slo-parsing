.PHONY: generate-slo-client
generate-slo-client: 
	@rm -rf generated
	@ docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:latest generate \
		-i /local/bundled.yaml \
		--git-repo-id terraform-provider-elasticstack \
		--git-user-id elastic \
		-p isGoSubmodule=true \
		-p packageName=slo \
		-p generateInterfaces=true \
		-g go \
		-o /local/generated/slo
	rm -rf generated/slo/go.mod generated/slo/go.sum generated/slo/test
	go fmt ./generated/...
