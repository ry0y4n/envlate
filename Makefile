# アプリ名
APP_NAME = my-app

# 出力ディレクトリ
OUTPUT_DIR = build

# ビルド対象ソース
SOURCE_FILE = ./src/main.go

# 対応プラットフォーム
PLATFORMS = linux/amd64 windows/amd64 darwin/amd64

# デフォルトターゲット: 全てのプラットフォームをビルド
all:
	@for platform in $(PLATFORMS); do \
		OS=$${platform%/*}; \
		ARCH=$${platform#*/}; \
		if [ "$${OS}" = "windows" ]; then \
			OUTPUT=$(OUTPUT_DIR)/$${OS}_$${ARCH}/$(APP_NAME).exe; \
		else \
			OUTPUT=$(OUTPUT_DIR)/$${OS}_$${ARCH}/$(APP_NAME); \
		fi; \
		mkdir -p $$(dirname $${OUTPUT}); \
		echo "Building for $${OS}/$${ARCH}..."; \
		GOOS=$${OS} GOARCH=$${ARCH} go build -o $${OUTPUT} $(SOURCE_FILE); \
	done

# Clean ターゲット
clean:
	@echo "Cleaning build directory..."
	@rm -rf $(OUTPUT_DIR)
