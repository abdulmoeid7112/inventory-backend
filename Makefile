include Makefile.variables

.PHONY: format todo test check prepare

local: | quiet
	@$(eval DOCKRUN= )
	@mkdir -p tmp
	@touch tmp/dev_image_id
quiet:
	@:

prepare: tmp/dev_image_id
tmp/dev_image_id:
	@mkdir -p tmp
	@docker rmi -f ${DEV_IMAGE} > /dev/null 2>&1 || true
	@docker build -t ${DEV_IMAGE} -f Dockerfile.dev .
	@docker inspect -f "{{ .ID }}" ${DEV_IMAGE} > tmp/dev_image_id

test: check db_start
	${DOCKTEST} bash ./scripts/test.sh

check:prepare format
	${DOCKRUN} bash ./scripts/check.sh

format:
	${DOCKRUN} bash ./scripts/format.sh

db_start: db_stop
	@docker run --name inventory-backend-mongo-db -p 27015-27017:27015-27017 -d mongo:4.2.0

db_stop:
	bash ./scripts/db_stop.sh

protogen:
	${DOCKRUN} bash ./scripts/proto_gen.sh