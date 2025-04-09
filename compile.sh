cmake -B build
cmake --build build
go get github.com/newrelic/go-agent/v3/newrelic
go get github.com/newrelic/go-agent/v3/integrations/nrgin
go build .
