.PHONY: all
all: fetch_docker build push

# BUILD
.PHONY: build
build:
	docker build -t jamrizzi/stripe-exporter:latest -f ./Dockerfile .

# PUSH
.PHONY: push
push:
	docker push jamrizzi/stripe-exporter:latest

# CLEAN
.PHONY: clean
clean: sweep bleach
	$(info cleaned)
.PHONY: sweep
sweep:
	$(info swept)
.PHONY: bleach
bleach:
	@rm -rf tmp
	$(info bleached)

# DEPENDANCIES
.PHONY: fetch_docker
fetch_docker:
ifeq ($(shell whereis docker), $(shell echo docker:))
	curl -L https://get.docker.com/ | bash
endif
	docker run hello-world
	$(info fetched docker)
