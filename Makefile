CC = go build
CFLAGS = 

# src := $(wildcard *.go)
# dir := $(notdir $(src))


%:
# 	通配匹配所有规则
	@echo "$@ < $^"
	@$(CC) -o $(basename $@) $(CFLAGS) $^
clean:
	@find . -type f | grep -v vscode | grep -v go | grep -v Makefile | grep -v cpp | grep -v git | xargs rm -f
