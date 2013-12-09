all: build

.PHONY: build

build: bin/sitemap-stats

clean:
	-rm -rf bin/

bin/sitemap-stats: main.go
	go build -o $@ main.go sitemap.go sitemap_index.go
