The ID for a user cert in the MSP is generated via the following code.

```go
// GetID returns a unique ID associated with the invoking identity.
func (c *ClientID) GetID() (string, error) {
	// When IdeMix, c.cert is nil for x509 type
	// Here will return "", as there is no x509 type cert for generate id value with logic below.
	if c.cert == nil {
		return "", fmt.Errorf("cannot determine identity")
	}
	// The leading "x509::" distinguishes this as an X509 certificate, and
	// the subject and issuer DNs uniquely identify the X509 certificate.
	// The resulting ID will remain the same if the certificate is renewed.
	id := fmt.Sprintf("x509::%s::%s", getDN(&c.cert.Subject), getDN(&c.cert.Issuer))
	return base64.StdEncoding.EncodeToString([]byte(id)), nil
}
```
