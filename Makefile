build:
	go build -o terraform-provider-fleetmanager
	terraform init

clean:
	rm terraform-provider-fleetmanager

install: build
	mkdir -p ~/.terraform.d/plugins
	cp terraform-provider-fleetmanager \
		~/.terraform.d/plugins/terraform-provider-fleetmanager
