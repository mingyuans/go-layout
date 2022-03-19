# select the cmd which contains 'apiserver'
API_COMMANDS ?= $(filter %apiserver, $(wildcard ${ROOT_DIR}/cmd/*))
API_BINS ?= $(foreach cmd,${API_COMMANDS},$(notdir ${cmd}))

CONFIGS_DIR := $(ROOT_DIR)/configs
LOCAL_CONFIG := $(addsuffix .yaml,${CONFIGS_DIR}/local-${API_BINS})

.PHONY: local.run
local.run: go.build.darwin_amd64.$(API_BINS)
	$(OUTPUT_DIR)/platforms/$(OS)/$(ARCH)/$(API_BINS) -c $(LOCAL_CONFIG)


