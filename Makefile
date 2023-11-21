API_PACKAGE_PATH := ./cmd/api/
API_BINARY_NAME := api
WEB_PACKAGE_PATH := ./cmd/web/
WEB_BINARY_NAME := web

.PHONY: tidy
tidy: 
	go fmt ./...
	go mod tidy -v

.PHONY: audit
audit:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	go test -race -buildvcs -vet=off ./...

.PHONY: test
test: 
	go test -v -race -buildvcs ./...

.PHONY: build
build:
	go build -o tmp/bin/${API_BINARY_NAME} ${API_PACKAGE_PATH}
	go build -o tmp/bin/${WEB_BINARY_NAME} ${WEB_PACKAGE_PATH}

.PHONY: run
run: build
	tmp/bin/${API_BINARY_NAME}
	tmp/bin/${WEB_BINARY_NAME}

.PHONY: dev_api
dev_api: 
	go run github.com/cosmtrek/air@v1.43.0 \
		--build.cmd "make build" \
		--build.bin "tmp/bin/${API_BINARY_NAME}" \
		--build.delay "100" \
		--build.exclude_dir "" \
		--build.include_ext "go, tpl, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
		--misc.clean_on_exit "true"

.PHONY: dev_web
dev_web: 
	go run github.com/cosmtrek/air@v1.43.0 \
		--build.cmd "make build" \
		--build.bin "tmp/bin/${WEB_BINARY_NAME}" \
		--build.delay "100" \
		--build.exclude_dir "" \
		--build.include_ext "go, tpl, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
		--misc.clean_on_exit "true"

clean:
	rm -rf bin tmp cmd/api/tmp cmd/web/tmp

