all: install

build:
	go build
install: build
	-rm "$(ADPCLI_BIN)/adpcli"
	cp -v adpcli "$(ADPCLI_BIN)"