build:
	go build -o terraform-provider-fleetmanager
	terraform init

clean:
	go clean
	rm terraform.tfstate
	rm terraform.tfstate.backup

install: build
	mkdir -p ~/.terraform.d/plugins
	cp terraform-provider-fleetmanager \
		~/.terraform.d/plugins/terraform-provider-fleetmanager
