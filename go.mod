module github.com/hezof/clients

go 1.21

require (
	github.com/hezof/core v0.0.0
	github.com/hezof/protojson v0.0.0
)

replace (
	github.com/hezof/log v0.0.0 => ../log
	github.com/hezof/core v0.0.0 => ../core
	github.com/hezof/protojson v0.0.0 => ../protojson
)
