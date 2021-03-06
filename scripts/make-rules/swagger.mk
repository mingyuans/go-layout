# Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file.

# ==============================================================================
# Makefile helper functions for swagger
#

.PHONY: swagger.run
swagger.run: tools.verify.swagger
	@echo "===========> Generating swagger API docs"
	@swagger generate spec --scan-models -w $(ROOT_DIR)/cmd/genswaggerdocs -o $(ROOT_DIR)/docs/api/swagger.json
	@echo "<=========== Please check docs at docs/api"

.PHONY: swagger.serve
swagger.serve: tools.verify.swagger
	@swagger serve -F=redoc --no-open --port 36666 $(ROOT_DIR)/docs/api/swagger.json
