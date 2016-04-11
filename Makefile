#######################################################################
## 
##
## 
#######################################################################
.PHONY: all build install clean st ut


all: build


build: fixer

SRC_FILES := $(shell find . -name "*.go"|grep -v "_test.go")

fixer: $(SRC_FILES)
	@echo "start building..."
	@go build
	@echo "done."

install: fixer datastore/pg.sql
	@echo "installing ..."	
	@./install.sh
	@echo "done."
run: install
	@./fixer

st:
	@echo "start st..."
	@./fixer &
	@sleep 1
	@-go test .
	@killall fixer
	@echo "done."
ut:
	@go test `go list  ./* 2>/dev/null`

clean:
	@echo "start cleaning..."
	@-killall fixer
	@./install.sh rm_db
	@rm -f fixer
	@echo "done."

#######################################################################
