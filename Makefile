all: interaction
.PHONY:all

## FFI

FFI_PATH:=extern/filecoin-ffi/
FFI_DEPS:=.install-filcrypto
FFI_DEPS:=$(addprefix $(FFI_PATH),$(FFI_DEPS))

$(FFI_DEPS): .filecoin-install ;

.filecoin-install: $(FFI_PATH)
	$(MAKE) -C $(FFI_PATH) $(FFI_DEPS:$(FFI_PATH)%=%)
	@touch $@
.PHONY:.filecoin-install

MODULES+=$(FFI_PATH)
BUILD_DEPS+=.filecoin-install
CLEAN+=.filecoin-install

$(MODULES): .update-modules ;

# dummy file that marks the last time modules were updated
.update-modules:
	git submodule update --init --recursive
	touch $@


interaction: ${BUILD_DEPS}
	rm -f lotus-interaction
	go build -o lotus-interaction ./cmd/lotus-interaction
.PHONY:interaction

clean:
	rm -rf $(CLEAN)
	rm -f interaction
