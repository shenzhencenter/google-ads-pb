VERSION = 14
INPUT_PATH = googleapis
OUTPUT_PATH = protogen
ADSLIB_PATH = $(INPUT_PATH)/google/ads/googleads/v${VERSION}
API_CONFIG_PATH = $(INPUT_PATH)/google/ads/googleads/v${VERSION}/googleads_v${VERSION}.yaml
GRPC_CONFIG_PATH = $(INPUT_PATH)/google/ads/googleads/v${VERSION}/googleads_grpc_service_config.json
BUILD_PATH = $(OUTPUT_PATH)/google/ads/googleads/v${VERSION}
BUILD_PKG = google.golang.org/genproto/googleapis/ads/googleads/v${VERSION}
PUBLISH_PKG = github.com/Optable/google-ads-pb/protogen
PATH := $(shell go env GOPATH)/bin:$(PATH)

.PHONY: generate
generate:
	@if [ `git branch --show-current` = "main" ] ; then \
        echo "You are on the main branch. Please checkout to another branch and retry." && exit 1; \
    fi
	@while [ -z "$$version" ]; do \
		read -p 'Enter the version number (like ${VERSION}): ' version; \
	done ;\
	$(MAKE) VERSION=$$version go-install-deps clone prepare-folders build fix-files cleanup
	@echo "Done!"
	@echo "Please do the following steps:"
	@echo "1. Update the version in the README.md"
	@echo "2. Create a PR and merge it to main"
	@echo "3. Create a tag with the version number (Ex: git tag v1.0.0)"
	@echo "4. Push the tag to the remote (Ex: git push origin v1.0.0)"
	@echo "5. Create a release in GitHub with the tag"
	@echo "6. Publish the release on golang.org (Ex: GOPROXY=proxy.golang.org go list -m github.com/Optable/google-ads-pb@v1.0.0)"
	@echo "7. Update the version to use the new one in the main project"

.PHONY: go-install-deps
go-install-deps:
	@echo "\nInstalling dependencies..."
	go get google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go get github.com/googleapis/gapic-generator-go/cmd/protoc-gen-go_gapic@latest

.PHONY: clone
clone:
	@echo "\nCloning googleapis repo..."
	rm -rf ./${INPUT_PATH}
	git clone --depth=1 --branch=master https://github.com/googleapis/googleapis.git  ./${INPUT_PATH}

.PHONY: prepare-folders
prepare-folders:
	@echo "\nPreparing folders..."
	rm -rf ./clients/* ./${OUTPUT_PATH}/*

.PHONY: build
build:
	@echo "\nBuilding the protos and clients..."
	protoc --proto_path=${INPUT_PATH} \
	--go_out=${OUTPUT_PATH} --go_opt=paths=source_relative \
	--go-grpc_out=${OUTPUT_PATH} --go-grpc_opt=paths=source_relative \
	--go_gapic_out=${OUTPUT_PATH} \
	--go_gapic_opt 'go-gapic-package=clients;clients' \
	--go_gapic_opt "api-service-config=./${API_CONFIG_PATH}" \
	--go_gapic_opt "grpc-service-config=./${GRPC_CONFIG_PATH}" \
	`find ./${ADSLIB_PATH} -name \*.proto`


.PHONY: fix-files
fix-files:
	@echo "\nFixing the files"
	@echo "\nMoving the proto files..."
	mv ${BUILD_PATH}/* ./${OUTPUT_PATH}
	@echo "\nRenaming the package..."
	find ./${OUTPUT_PATH} -name '*.go' -exec sed -i.bak "s|${BUILD_PKG}|${PUBLISH_PKG}|" '{}' \; -exec rm {}.bak \; \
	&& find ./${OUTPUT_PATH}/clients/internal/snippets -name '*.go' -exec sed -i.bak "s|\"clients\"|\"${PUBLISH_PKG}/clients\"|g" '{}' \; -exec rm {}.bak \;
	@echo "\nMoving the clients..."
	mv ./${OUTPUT_PATH}/clients/* ./clients/
	@if [ -f ${OUTPUT_PATH}/resources/experiment_arm.pb.go ] ; then \
		echo "\nRename experiment_arm.pb.go to experimentarm.pb.go"; \
		mv -f ./${OUTPUT_PATH}/resources/experiment_arm.pb.go ./${OUTPUT_PATH}/resources/experimentarm.pb.go; \
	fi

.PHONY: cleanup
cleanup:
	@echo "\nCleaning up..."
	rm -rf ./${INPUT_PATH}	\
	&& rm -rf ./${OUTPUT_PATH}/google  \
	&& rm -rf ./${OUTPUT_PATH}/clients
