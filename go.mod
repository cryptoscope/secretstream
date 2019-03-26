module go.cryptoscope.co/secretstream

require (
	github.com/cryptix/go v1.5.0
	github.com/go-kit/kit v0.9.0 // indirect
	github.com/hashicorp/go-multierror v1.0.0
	github.com/pkg/errors v0.8.1
	github.com/stretchr/testify v1.4.0
	go.cryptoscope.co/netwrap v0.1.0
	golang.org/x/crypto v0.0.0-20191002192127-34f69633bfdc
	golang.org/x/sys v0.0.0-20191007154456-ef33b2fb2c41 // indirect
)

go 1.13

replace golang.org/x/crypto => github.com/cryptix/golang_x_crypto v0.0.0-20200303113948-2939d6771b24
