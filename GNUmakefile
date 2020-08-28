GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

DIR=~/.terraform.d/plugins

DEFAULT_CONFIG_FOLDER= ~/.cloudknox/

default: build

build: install init_config

install: fmt
	@printf "\n==> Installing provider to $(DIR)\n"
	mkdir -vp $(DIR)
	go build -o $(DIR)/terraform-provider-cloudknox
	@printf "\n    ==> Provider installed\n"

uninstall:
	@printf "\n==> Uninstalling provider from $(DIR)\n"
	@rm -vf $(DIR)/terraform-provider-cloudknox
	@rm -rf $(DEFAULT_CONFIG_FOLDER)
	@printf "\n    ==> Provider uninstalled\n"

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"


init_config:
	@printf "\n==> Creating config folder at $(DEFAULT_CONFIG_FOLDER)\n"
	mkdir -vp $(DEFAULT_CONFIG_FOLDER)
	@printf "\n    ==> Config folder created\n"

testacc: fmtcheck
	TF_ACC=1 go test terraform-provider-cloudknox/cloudknox -v -timeout 120m

test: 
	provider "google" {
  	credentials = "/Users/sravani/Downloads/devopstest-218421-3c2821f32c55.json"
  	project     = "devopstest-218421"
  	region      = "us-west2"
	}	
	data "cloudknox_role_policy" "user-activity-gcp-role" {
    		name = "role-activity-gcp-role"
    		output_path = "./"
    		auth_system_info = {
         	id = "devopstest-218421",
         	type = "GCP"
     	}
	identity_type = "USER"
    		identity_ids = [
        	"sravani@cloudknox.io"
	]

    	filter_history_days = 90
	}
	
