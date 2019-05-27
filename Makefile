VERSION=0.0.2
PATH_BUILD=build/
FILE_COMMAND=kitago
FILE_ARCH=darwin_amd64

clean:
	@rm -rf ./build

build: clean
	@$(GOPATH)/bin/goxc \
	  -bc="darwin,amd64" \
	  -pv=$(VERSION) \
	  -d=$(PATH_BUILD) 

version:
	@echo $(VERSION)

install:
	install -d -m 755 '$(HOME)/bin/'
	install $(PATH_BUILD)$(VERSION)/$(FILE_ARCH)/$(FILE_COMMAND) '$(HOME)/bin/$(FILE_COMMAND)'