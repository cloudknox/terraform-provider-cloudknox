GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

DIR=~/.terraform.d/plugins

DEFAULT_CONFIG_FOLDER= ~/.cloudknox/

default: build

build: install init_config

install: fmt
	@printf "\n==> Installing provider to $(DIR)\n"
	mkdir -vp $(DIR)
	go build -o $(DIR)/terraform-provider-cloudknox

uninstall:
	@rm -vf $(DIR)/terraform-provider-cloudknox

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"


init_config:
	mkdir -vp $(DEFAULT_CREDENTIALS_FOLDER)

testacc: fmtcheck
	TF_ACC=1 go test terraform-provider-cloudknox/cloudknox -v -timeout 120m
