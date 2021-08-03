package status_controller_test

import (
	"context"
	"fmt"
	"github.com/fahani/asia-car/src/infrastructure/status/controller"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"testing"
)

func setup(port string) *http.Server {
	srv := http.Server{Addr: fmt.Sprintf(":%s", port)}
	http.HandleFunc("/status", status_controller.Status)
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()
	time.Sleep(500 * time.Millisecond)
	return &srv
}

func teardown(s *http.Server) error {
	err := s.Close()
	if err != nil {
		return err
	}
	err = s.Shutdown(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

func getBodyString(resp *http.Response) (string, error) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}

func TestStatusRequest(t *testing.T) {
	serv := setup("9000")

	resp, err := http.Get("http://localhost:9000/status")

	assert.Nil(t, err)

	body, err := getBodyString(resp)
	assert.Nil(t, err)

	assert.Equal(t, `{"status": "ok"}`, body)

	err = teardown(serv)
	assert.Nil(t, err)
}
