API_PACKAGE_PATH := ./cmd/api
API_BINARY_NAME := api
WEB_PACKAGE_PATH := ./cmd/web
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
test: export APP_ENV=test
test: 
	go test -v -race -buildvcs ./...

.PHONY: build_api
build_api:
	go build -o tmp/bin/${API_BINARY_NAME} ${API_PACKAGE_PATH}

.PHONY: build_web
build_web:
	go build -o tmp/bin/${WEB_BINARY_NAME} ${WEB_PACKAGE_PATH}

.PHONY: build
build: build_api build_web

.PHONY: dev_api
dev_api: export APP_ENV=development
dev_api: 
	go run github.com/cosmtrek/air@v1.43.0 \
		--build.cmd "make build_api" \
		--build.bin "tmp/bin/${API_BINARY_NAME}" \
		--build.delay "100" \
		--build.exclude_dir "" \
		--build.include_ext "go" \
		--misc.clean_on_exit "true"

.PHONY: dev_web
dev_web: export APP_ENV=development
dev_web: 
	go run github.com/cosmtrek/air@v1.43.0 \
		--build.cmd "make build_web" \
		--build.bin "tmp/bin/${WEB_BINARY_NAME}" \
		--build.delay "100" \
		--build.exclude_dir "" \
		--build.include_ext "go, tpl, templ, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
		--misc.clean_on_exit "true"

clean:
	rm -rf bin tmp 
	rm -rf ${API_PACKAGE_PATH}/tmp ${API_PACKAGE_PATH}/bin
	rm -rf ${WEB_PACKAGE_PATH}/tmp ${WEB_PACKAGE_PATH}/bin

.PHONY: dev
dev: export APP_ENV=development
dev:
	make -j2 dev_web dev_api
