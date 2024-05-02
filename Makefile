GOOS_LIST = linux darwin
GOARCH_LIST = amd64 arm64

all: clean release

build:
	@for GOOS in $(GOOS_LIST); do \
		for GOARCH in $(GOARCH_LIST); do \
			FILENAME=tut_$${GOOS}_$${GOARCH}; \
			echo "Building $$FILENAME"; \
			GOOS=$$GOOS GOARCH=$$GOARCH go build -o bin/$$FILENAME main.go; \
		done; \
	done

release: build
	@cd bin && for FILE in *; do \
		FILENAME=$${FILE}.tgz; \
		echo "Creating $$FILENAME"; \
		tar -czvf $$FILENAME $$FILE; \
	done

clean:
	@rm -rf bin/*