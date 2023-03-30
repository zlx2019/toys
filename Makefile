# example: make tag=v1.0.x

# 要操作的tag标签
tag := $(firstword $(MAKECMDGOALS))

# 推送tag
push: create
	git push origin $(tag)
	@echo "push tag $(tag)"

# 创建tag
.PHONY: create
create:
	git tag $(tag)
	@echo "create tagName: $(tag)"


# 删除tag,推送到远程
del: del_local
	git push origin :refs/tags/$(tag)

# 先删除本地tag
.PHONY: del_local
del_local:
	git tag -d $(tag)
	@echo "delete tag $(tag)"

# 获取到tag中最后一个tag标签
LAST_TAG := $(shell git describe --tags --abbrev=0)
