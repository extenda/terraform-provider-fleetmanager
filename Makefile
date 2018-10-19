TEST?=$$(go list ./... |grep -v 'vendor')

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

test:
	go test -i $(TEST) || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m
