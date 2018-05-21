REPOSITORY=joaofnfernandes
IMAGE_DEPS=$(wildcard bin/api/*)

###############################################################################
# Main targets
###############################################################################

.PHONY: default
default: $(IMAGE_DEPS)
	docker-compose up

.PHONY: compose-validate
compose-validate:
	docker-compose -f docker-compose.yml config

.PHONY: clean
clean:
	@echo "Deleting artifacts"
	@rm -f bin/*

###############################################################################
# Utils
###############################################################################

.PHONY: images
images: bin/api

###############################################################################
# Docker images
###############################################################################

# Creates the image for the API service
API_SOUCES=$(shell find api -type f)
bin/api: $(API_SOURCES)
	@echo "Building $(REPOSITORY)/$(@F) image"
	@cd $(@F); docker build -t $(REPOSITORY)/$(@F) .; cd $(CWD)
	@mkdir -p bin; touch $@
