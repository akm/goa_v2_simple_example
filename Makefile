BASE_PACKAGE_PATH=github.com/akm/goa_v2_simple_example

.PHONY: backup_vendor restore_vendor
backup_vendor:
	@mv vendor vendor.bak || echo "WARNING: vendor not found"
restore_vendor:
	@mv vendor.bak vendor || echo "WARNING: vendor.bak not found"

.PHONY: goa_gen
goa_gen: backup_vendor
	goa gen $(BASE_PACKAGE_PATH)/design; \
	$(MAKE) restore_vendor

.PHONY: goa_example
goa_example: backup_vendor
	goa example $(BASE_PACKAGE_PATH)/design; \
	$(MAKE) restore_vendor
