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

install: main.go
	mkdir -vp $(DIR)
	go build -o $(DIR)/terraform-provider-cloudknox

init_credentials:
	mkdir -vp $(DEFAULT_CREDENTIALS_FOLDER)

testacc:
	TF_ACC=1 go test terraform-provider-cloudknox/cloudknox $(TEST) -v $(TESTARGS) -timeout 120m

test: $(DIR)/terraform-provider-cloudknox main.tf $(CONFIGURATION_DEST)terraform-provider-cloudknox-config.yaml
	find ./ -name "*.tf*" -not -name "main.tf" -exec rm {} \;
	terraform init
	terraform apply -input=false -auto-approve -parallelism=10

clean:
	find ./ -name "*.tf" -not -name "main.tf" -exec rm {} \;
	rm -f info.log crash.log terraform-provider-cloudknox.exe terraform.tfstate

