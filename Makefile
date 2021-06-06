CC = go build
CFLAGS = 

# src := $(wildcard *.go)
# dir := $(notdir $(src))


%:
# 	通配匹配所有规则
	@echo "$@ < $^"
	@$(CC) -o $(basename $@) $(CFLAGS) $^
clean:
	@ls . | grep -v go | grep -v Makefile | xargs rm -f