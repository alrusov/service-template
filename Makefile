SHELL := /bin/bash 

APP  := $(shell basename `pwd`)

OS := linux
ARCH := $(shell arch)

TS   := $(shell date -u +'%Y-%m-%d %H:%M:%S')
YEAR := $(shell date -u +'%Y')

COPYRIGHT = (C)_Alexey_Rusov_(rolic402@mail.ru),_2017-$(YEAR)

SPACE =
SPACE +=

TS        := $(subst $(SPACE),_,$(strip $(TS)))
COPYRIGHT := $(subst $(SPACE),_,$(strip $(COPYRIGHT)))

CGO := 0
EXTRA_LD := -extldflags -static

ifneq (,$(wildcard VERSION))
	VERSION := $(shell cat VERSION)
	BUILD_NUMBER_FILE := BUILD_NUMBER

	ifneq (,$(wildcard $(BUILD_NUMBER_FILE)))
		BUILD_NUMBER := $(shell cat $(BUILD_NUMBER_FILE))
	else
		BUILD_NUMBER := 1
	endif

	ifneq ($(strip $(BUILD_NUMBER)),)
		VERSION := $(VERSION).$(BUILD_NUMBER)
	endif
endif

.DEFAULT_GOAL := build

build:
	if [ -d "vendor/github.com" ]; then \
		echo 'Please delete the vendor/github.com directory'; \
		echo 'See https://github.com/golang/go/issues/19000 for details'; \
	else \
		((n = $(BUILD_NUMBER) + 1));\
		echo $${n} >$(BUILD_NUMBER_FILE); \
		CGO_ENABLED=$(CGO) \
                 go build -o "cmd/$(OS)/$(ARCH)/$(APP)" \
                   --ldflags "$(EXTRA_LD) \
                   -X github.com/alrusov/misc.appVersion=$(VERSION) -X github.com/alrusov/misc.buildTime=$(TS) -X github.com/alrusov/misc.copyright=$(COPYRIGHT)"; \
	fi;
