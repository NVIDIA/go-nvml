GEN_DIR = $(PWD)/gen
PKG_DIR = $(PWD)/pkg
GEN_BINDINGS_DIR = $(GEN_DIR)/nvml
PKG_BINDINGS_DIR = $(PKG_DIR)/nvml

SOURCES = $(shell find $(GEN_BINDINGS_DIR) -type f)

.PHONY: all test clean
.PHONY: bindings test-bindings clean-bindings

all: bindings
test: test-bindings
clean: clean-bindings

$(PKG_BINDINGS_DIR): $(SOURCES)
	c-for-go -out $(PKG_DIR) $(GEN_BINDINGS_DIR)/nvml.yml
	cp $(GEN_BINDINGS_DIR)/*.go $(PKG_BINDINGS_DIR)
	cp $(GEN_BINDINGS_DIR)/nvml.h $(PKG_BINDINGS_DIR)
	patch $(PKG_BINDINGS_DIR)/nvml.h $(GEN_BINDINGS_DIR)/nvml.h.patch
	cd $(PKG_BINDINGS_DIR); \
		go tool cgo -godefs types.go > types_gen.go; \
		go fmt types_gen.go; \
	cd -> /dev/null
	rm -rf $(PKG_BINDINGS_DIR)/types.go $(PKG_BINDINGS_DIR)/_obj

bindings: $(PKG_BINDINGS_DIR)

test-bindings: $(PKG_BINDINGS_DIR)
	cd $(PKG_BINDINGS_DIR); \
		go test -v .; \
	cd -> /dev/null

clean-bindings:
	rm -rf $(PKG_BINDINGS_DIR)
