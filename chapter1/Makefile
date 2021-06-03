CC = go build
FLAG = -x

.PHONY: ls getcputc

ls: ls.go
	$(CC) -o $@ $(FLAG) $^
getcputc: getcputc.go
	$(CC) -o $@ $(FLAG) $^
clean:
	@ls . | grep -v go | grep -v Makefile | xargs rm -f