module github.com/newrelic/opentelemetry-collector-releases

go 1.18

require (
	github.com/goreleaser/goreleaser v1.13.1
	github.com/goreleaser/nfpm/v2 v2.22.1
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/aymanbagabas/go-osc52 v1.2.1 // indirect
	github.com/caarlos0/log v0.1.10 // indirect
	github.com/charmbracelet/lipgloss v0.6.1-0.20220911181249-6304a734e792 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/goreleaser/fileglob v1.3.0 // indirect
	github.com/iancoleman/orderedmap v0.2.0 // indirect
	github.com/invopop/jsonschema v0.7.0 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/muesli/reflow v0.3.0 // indirect
	github.com/muesli/termenv v0.13.0 // indirect
	github.com/rivo/uniseg v0.4.2 // indirect
	golang.org/x/sys v0.1.0 // indirect
)

//TODO REMOVE: Just For testing
replace github.com/newrelic/opentelemetry-collector-releases => ./
