
GOCMD = go
GOARCH = amd64
GOOS = linux
LDFLAGS = "-s -w"

HANDLERS = itemHandler userHandler
LAMBDAS = handlers
SERVERLESS_DEPLOY_CMD = deploy
SERVERLESS_REMOVE_CMD = remove

.PHONY: all build clean distclean deploy remove

build: $(HANDLERS)
all: $(HANDLERS)
	@echo
	@echo "Finish buld all...."
%:
	@echo
	@echo "build $@..."
	env GOARCH=$(GOARCH) GOOS=$(GOOS) $(GOCMD) build -ldflags=$(LDFLAGS) -o ./bin/$@ ./src/handlers/$@


clean:
	rm -rf ./bin ./vendor Gopkg.lock

distclean:
	rm -rf ./bin ./vendor Gopkg.lock


deploy: $(LAMBDAS)
	@echo 
	@echo "Finish deploy all...."
$(LAMBDAS):
	@echo
	@echo "$(CMD) $@..."
	sls $(SERVERLESS_DEPLOY_CMD) --verbose --config=./serverless/$@.yml

