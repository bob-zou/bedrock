TARGET = bedrock

all: $(TARGET)

.PHONY: $(TARGET)
$(TARGET): $(GOFILES)
	packr2 build -ldflags "-s -w" -o $(TARGET)

clean:
	@rm -rf ./$(TARGET)

.PHONY: clean all