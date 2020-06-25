default: build

build: main.go
	go build -o terraform-provider-cloudknox.exe

test: terraform-provider-cloudknox.exe main.tf ./cloudknox/config/resources.yaml
	find ./ -name "*.tf*" -not -name "main.tf" -exec rm {} \;
	terraform init
	terraform apply -input=false -auto-approve -parallelism=10

clean:
	find ./ -name "*.tf" -not -name "main.tf" -exec rm {} \;
	rm -f info.log crash.log terraform-provider-cloudknox.exe terraform.tfstate