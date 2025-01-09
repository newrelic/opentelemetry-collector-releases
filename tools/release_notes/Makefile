THIS_MAKEFILE_DIR := $(realpath $(dir $(realpath $(lastword $(MAKEFILE_LIST)))))
OUTPUT_DIR=${THIS_MAKEFILE_DIR}/output
COMMIT_METADATA_FILE=${OUTPUT_DIR}/commit-metadata.yml

.PHONY: install-chglog
install-chglog:
	@command -v chglog &> /dev/null || go install github.com/goreleaser/chglog/cmd/chglog@v0.6.2

.PHONY: gather-relevant-commits
gather-relevant-commits: NUMBER_OF_BACKTRACKED_TAGS=2
gather-relevant-commits: install-chglog
	chglog init --output ${COMMIT_METADATA_FILE}
	yq '.[0:$(shell expr ${NUMBER_OF_BACKTRACKED_TAGS} + 1)]' -i ${COMMIT_METADATA_FILE}

.PHONY: most-recent-release-notes
most-recent-release-notes: gather-relevant-commits
	chglog format --template-file ${THIS_MAKEFILE_DIR}/release_notes.tmpl \
	--input ${COMMIT_METADATA_FILE} \
	--output ${OUTPUT_DIR}/most_recent_release_notes.md