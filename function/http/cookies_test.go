package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var c = &fnRewriteCookies{}

func TestRewriteCookies(t *testing.T) {
	cookies := []interface{}{"lo=; Path=/; Expires=Wed, 01 Jan 1800 00:00:00 GMT; Secure", "tsc=UcRKVG68jC...; Domain=cloud.tibco.com; Path=/; Max-Age=86399; HttpOnly; Secure"}
	final, err := c.Eval(cookies, "tsc", "mashery.com", "/mashery-path/")
	assert.Nil(t, err)
	assert.NotNil(t, final)
	expectedResult := []interface{}{"lo=; Path=/; Expires=Wed, 01 Jan 1800 00:00:00 GMT; Secure", "tsc=UcRKVG68jC...; Domain=mashery.com; Path=/mashery-path/; Max-Age=86399; HttpOnly; Secure"};
	assert.Equal(t, expectedResult, final)
}
