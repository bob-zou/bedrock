TARGET = bedrock
VERSION := $(shell cat VERSION)

all: $(TARGET)

.PHONY: $(TARGET)
$(TARGET): $(GOFILES)
	packr2 build -ldflags "-X main.commit=`git rev-parse --short HEAD` -X main.version=$(VERSION)" -o $(TARGET)

clean:
	@rm -rf ./$(TARGET)

.PHONY: clean all