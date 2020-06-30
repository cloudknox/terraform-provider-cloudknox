CONFIGURATION_FILE= ./cloudknox/config/terraform-provider-cloudknox-config.yaml
CONFIGURATION_DEST= ~/opt/cloudknox/
LOG_DEST = ~/var/log/cloudknox/

default: build

build: main.go
	mkdir -p $(CONFIGURATION_DEST) && cp $(CONFIGURATION_FILE)  $(CONFIGURATION_DEST)
	mkdir -p $(LOG_DEST)
	go build -o terraform-provider-cloudknox.exe

test: terraform-provider-cloudknox.exe main.tf $(CONFIGURATION_DEST)terraform-provider-cloudknox-config.yaml
	find ./ -name "*.tf*" -not -name "main.tf" -exec rm {} \;
	terraform init
	terraform apply -input=false -auto-approve -parallelism=10

clean:
	find ./ -name "*.tf" -not -name "main.tf" -exec rm {} \;
	rm -f info.log crash.log terraform-provider-cloudknox.exe terraform.tfstate

testacc:
	TF_ACC=1 go test terraform-provider-cloudknox/cloudknox $(TEST) -v $(TESTARGS) -timeout 120m