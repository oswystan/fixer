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

install: fixer db/pg.sql
	@echo "installing ..."	
	@./install.sh
	@echo "done."

clean:
	@echo "start cleaning..."
	@./install.sh rm_db
	@rm -f fixer
	@echo "done."

#######################################################################
