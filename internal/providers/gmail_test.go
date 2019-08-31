package providers

import (
	"context"
	"testing"

	"net/http"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMakeChallenge(t *testing.T) {
	s := "somestring"
	ret := makeChallenge(&s)
	assert.Equal(t, ret, "Y_b-eXAm15Tg3D4r0nmu4Z3S-NtnSIFypkS7aHkqVww")
}

// Ensure that code is being sent on the channel
func TestCreateResponseListenerDataIsSent(t *testing.T) {
	srv, c := createResponseListener()
	defer srv.Shutdown(context.TODO())
	_, err := http.Get("http://" + srv.Addr + oauth2Endpoint + "?code=1234")
	if err != nil {
		assert.Fail(t, "Get request failed: "+err.Error())
	}
	select {
	case r := <-*c:
		assert.Equal(t, r, "1234")
	case <-time.After(time.Second * 10):
		assert.Fail(t, "Timed out waiting for code")
	}
}

// Ensure that the channel is being closed after the server shuts down
func TestCreateResponseListenerChannelIsClosed(t *testing.T) {
	srv, c := createResponseListener()
	srv.Shutdown(context.TODO())
	select {
	case d, ok := <-*c:
		if ok {
			assert.Fail(t, "Channel is still open and had data in the queue: "+d)
		}
	case <-time.After(time.Second * 10):
		assert.Fail(t, "Channel is still open!")
	}
}
