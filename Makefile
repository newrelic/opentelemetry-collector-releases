PROJECT_WORKSPACE	?= $(CURDIR)
INCLUDE_TEST_DIR	?= $(PROJECT_WORKSPACE)/test

include $(INCLUDE_TEST_DIR)/test.mk
include $(INCLUDE_TEST_DIR)/terraform/alerts/alerts-terraform.mk
include $(PROJECT_WORKSPACE)/build.mk

# common utils
include ./Makefile.Common

