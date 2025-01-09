### Creating a changelog

#### Tools
We manually leverage the goreleaser dependency [chglog](# https://github.com/goreleaser/chglog/tree/main) until we can automate the process fully via automatic changelog generation by goreleaser

#### Steps
- (optional) Adjust the number of backtracked tags used by `make gather-relevant-commits`. This can be necessary when tags were created that where not officially released as we'll then have to fold commits from multiple tags into one release.
- `make most-recent-release-notes`
- Decide the appropriate category for each commit
- Paste the final notes into the GitHub prerelease form