#TEST?=$$(go list ./... |grep -v 'vendor')
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

PKG_NAME=datadog
DIR=~/.terraform.d/plugins

CONFIGURATION_FILE= ./cloudknox/config/terraform-provider-cloudknox-config.yaml
CONFIGURATION_DEST= /opt/cloudknox/

LOG_DEST = /var/log/cloudknox/

DEFAULT_CREDENTIALS_FOLDER= ~/.cnx/

default: build

build: install init_credentials
	mkdir -p $(CONFIGURATION_DEST) && cp $(CONFIGURATION_FILE)  $(CONFIGURATION_DEST)
	mkdir -p $(LOG_DEST)
	@printf "\nSet Credentials -> $(DEFAULT_CREDENTIALS_FOLDER)creds.conf"
	@printf "\nLogs -> $(LOG_DEST)application.log"
	@printf "\nBackend Config -> $(CONFIGURATION_DEST)terraform-provider-cloudknox-config.yaml\n"

install: fmtcheck
	@printf "\nInstalling provider to $(DIR)\n"
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

test: fmtcheck
	find ./ -name "*.tf*" -not -name "main.tf" -exec rm {} \;
	terraform init
	terraform apply -input=false -auto-approve -parallelism=10

clean:
	find ./ -name "*.tf" -not -name "main.tf" -exec rm {} \;
	rm -f info.log crash.log terraform-provider-cloudknox.exe terraform.tfstate

