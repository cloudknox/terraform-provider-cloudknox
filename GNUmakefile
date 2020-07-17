GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

BASE_URL = "https://olympus.aws-staging.cloudknox.io"

DIR=~/.terraform.d/plugins
LOG_DEST = ./application.log

DEFAULT_CREDENTIALS_FOLDER= ~/.cnx/

default: build

build: install init_credentials
	set CNX_BASE_URL=$(BASE_URL)
	mkdir -p $(LOG_DEST)
	@printf "==> Directories Created"
	@printf "\n   ==>Set Credentials -> $(DEFAULT_CREDENTIALS_FOLDER)creds.conf"
	@printf "\n   ==>Logs -> $(LOG_DEST)application.log\n"

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


init_credentials:
	mkdir -vp $(DEFAULT_CREDENTIALS_FOLDER)

testacc: fmtcheck
	TF_ACC=1 go test terraform-provider-cloudknox/cloudknox -v -timeout 120m
