#######################################################################
## 
##
## 
#######################################################################
.PHONY: all build


all: build


build: fixer

fixer:
	@echo "start building..."
	@go build
	@echo "done."

clean:
	@rm -f fixer

#######################################################################