default: build

build: main.go
	go build -o terraform-provider-cloudknox.exe

test: terraform-provider-cloudknox.exe main.tf ./cloudknox/config/resources.yaml
	find ./ -name "*.tf*"  -not -name "main.tf" -exec rm {} \;
	terraform init
	terraform apply -input=false -auto-approve

clean:
	find ./ -name "*.tf" -not -name "main.tf" -exec rm {} \;
	rm info.log crash.log terraform-provider-cloudknox.exe *.tfstate