# 启动依赖加载
.PHONY: setup
setup:
	@sh ./.scripts/setup.sh

# 依赖加载
.PHONY: tidy
tidy:
	@go mod tidy

# 运行测试程序
.PHONY: ut
ut:
	@go test -race ./...

# 格式化代码
.PHONY: fmt
fmt:
	@sh ./.scripts/fmt.sh

# 项目检查命令
.PHONY: check
check:
	@$(MAKE) --no-print-directory setup
	@$(MAKE) --no-print-directory tidy
	@$(MAKE) --no-print-directory fmt
	@$(MAKE) --no-print-directory ut
