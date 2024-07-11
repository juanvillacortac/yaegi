GC=go

# Static linting of source files. See .golangci.toml for options
check:
	golangci-lint run

# Generate stdlib/syscall/syscall_GOOS_GOARCH.go for all platforms
gen_all_syscall: internal/cmd/extract/extract
	@for v in $$(${GC} tool dist list); do \
		echo syscall_$${v%/*}_$${v#*/}.go; \
		GOOS=$${v%/*} GOARCH=$${v#*/} ${GC} generate ./stdlib/syscall ./stdlib/unrestricted; \
	done

internal/cmd/extract/extract:
	rm -f internal/cmd/extract/extract
	${GC} generate ./internal/cmd/extract

generate: gen_all_syscall
	${GC} generate

install:
	GOFLAGS=-ldflags=-X=main.version=$$(git describe --tags) go install ./...

tests:
	${GC} test -v ./...
	${GC} test -race ./interp

# https://github.com/goreleaser/godownloader
install.sh: .goreleaser.yml
	godownloader --repo=juanvillacortac/yaegi -o install.sh .goreleaser.yml

generic_list = cmp/cmp.go slices/slices.go slices/sort.go slices/zsortanyfunc.go maps/maps.go \
			   sync/oncefunc.go sync/atomic/type.go

# get_generic_src imports stdlib files containing generic symbols definitions
get_generic_src:
	eval "`${GC} env`"; echo $$GOROOT; gov=$${GOVERSION#*.}; gov=$${gov%.*}; \
	for f in ${generic_list}; do \
		nf=stdlib/generic/go1_$${gov}_`echo $$f | tr / _`.txt; echo "nf: $$nf"; \
		if [[ -f "$$GOROOT/src/$$f" && -s "$$GOROOT/src/$$f" ]]; then \
			cat "$$GOROOT/src/$$f" > "$$nf"; \
		fi \
	done

gen_symbols: generate get_generic_src

.PHONY: check gen_all_syscall internal/cmd/extract/extract get_generic_src install
