package cidr

// ParseIPV4GatewayNetmaskError is used to indicate any error with parsing the
// IPV4 address and the netmask of the ENI
type ParseIPV4GatewayNetmaskError struct {
	operation string
	origin    string
	message   string
}

func (err *ParseIPV4GatewayNetmaskError) Error() string {
	return err.operation + " " + err.origin + ": " + err.message
}

func newParseIPV4GatewayNetmaskError(operation string, origin string, message string) error {
	return &ParseIPV4GatewayNetmaskError{
		operation: operation,
		origin:    origin,
		message:   message,
	}
}
