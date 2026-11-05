.PHONY: all config config-check build-server server dev client manifest release clean

all: build-server

config:
	./scripts/render-config.sh

config-check:
	./scripts/render-config.sh --check

build-server:
	cd server && CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -o paperplane-server ./cmd/server

server:
	cd server && go run ./cmd/server

dev: config
	docker compose -f docker-compose.dev.yml up -d --build
	@echo "Go 服务端已起；客户端： cd client && flutter pub get && flutter run -d chrome"

client:
	cd client && flutter pub get && flutter run -d chrome

manifest:
	python3 scripts/gen-manifest.py --root .

release: config-check build-server manifest
	@mkdir -p dist
	@echo "产物：server/paperplane-server, FILE_MANIFEST.md"

clean:
	rm -f server/paperplane-server dist/*.zip
