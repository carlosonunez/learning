SHELL := /usr/bin/env bash
MAKEFLAGS += --silent

DOCKER_IMAGE_NAME ?= subfactorial_app
APP_VERSION := $(shell git rev-parse HEAD | head -c 8)
DOCKER_BASE_IMAGE_NAME_AND_VERSION := $(DOCKER_IMAGE_NAME)_base:$(APP_VERSION)
DOCKER_IMAGE_NAME_AND_VERSION := $(DOCKER_IMAGE_NAME):$(APP_VERSION)

.PHONY: \
	get_version_number \
	build_local_docker_image \
	remove_local_docker_image

get_version_number:
	git rev-parse HEAD | head -c 8

build_local_base_docker_image:
	if ! docker ps -a | grep -q "$(DOCKER_BASE_IMAGE_NAME_AND_VERSION)"; \
	then \
		docker build --target base_image -t "$(DOCKER_BASE_IMAGE_NAME_AND_VERSION)" $$PWD; \
	fi

remove_local_base_docker_image:
	if docker ps -a | grep -q "$(DOCKER_BASE_IMAGE_NAME_AND_VERSION)"; \
	then \
		docker rmi "$(DOCKER_BASE_IMAGE_NAME_AND_VERSION)"; \
	fi

.PHONY: \
	static_analysis \
	unit_tests \
	integration_tests \
	docker_image_security_scans

static_analysis:
	$(MAKE) test_setup && \
	if ! docker run --rm \
		--tty \
		--volume $(PWD):/app \
		$(DOCKER_BASE_IMAGE_NAME_AND_VERSION) rubocop /app; \
	then \
		$(MAKE) test_teardown_failure && \
		>&2 echo "ERROR: Static analysis failed."; \
	fi

.PHONY: \
	test_setup \
	test_teardown_failure \
	test_teardown_success

test_setup:
	$(MAKE) build_local_base_docker_image

test_teardown_failure:
	$(MAKE) remove_local_base_docker_image

test_teardown_success:
	$(MAKE) remove_local_base_docker_image
