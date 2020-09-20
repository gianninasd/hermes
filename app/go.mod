module dg/app

go 1.15

replace dg/utils => ../utils

replace dg/client => ../client

require (
	dg/client v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.1.2
	github.com/magiconair/properties v1.8.3
	github.com/satori/go.uuid v1.2.0 // indirect
)
