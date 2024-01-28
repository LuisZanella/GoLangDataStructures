package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}
func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("Not Implemented")
}
func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)
	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("Spy Store got cancelled")
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

//	func (s *SpyStore) Cancel() {
//		time.Sleep(100 * time.Millisecond)
//		s.cancelled = true
//	}
//
//	func (s *SpyStore) assertWasCancelled() {
//		s.t.Helper()
//		if !s.cancelled {
//			s.t.Error("store was not told to cancel")
//		}
//	}
//
//	func (s *SpyStore) assertWasNotCancelled() {
//		s.t.Helper()
//		if s.cancelled {
//			s.t.Error("store was told to cancel")
//		}
//	}
func TestServer(t *testing.T) {

	t.Run("Test the reponse", func(t *testing.T) {
		data := "Hello, world"
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)
		if response.Body.String() != data {
			t.Errorf("Got: %q Want: %q", response.Body.String(), data)
		}
	})
	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		if response.written {
			t.Error("a response should not have been written")
		}
	})
	//t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
	//	data := "Hello, world"
	//	store := &SpyStore{response: data, t: t}
	//	svr := Server(store)
	//	request := httptest.NewRequest(http.MethodGet, "/", nil)
	//
	//	cancellingCtx, cancel := context.WithCancel(request.Context())
	//	time.AfterFunc(5*time.Millisecond, cancel)
	//	request = request.WithContext(cancellingCtx)
	//
	//	response := httptest.NewRecorder()
	//
	//	svr.ServeHTTP(response, request)
	//	if !store.cancelled {
	//		t.Errorf("Store was not told to cancel")
	//	}
	//})
	//t.Run("Returns data from store", func(t *testing.T) {
	//	data := "Hello, World"
	//	store := &SpyStore{response: data, t: t}
	//	svr := Server(store)
	//	request := httptest.NewRequest(http.MethodGet, "/", nil)
	//	response := httptest.NewRecorder()
	//
	//	svr.ServeHTTP(response, request)
	//
	//	if response.Body.String() != data {
	//		t.Errorf("Got: %q Want: %q", response.Body.String(), data)
	//	}
	//	if store.cancelled {
	//		t.Errorf("It shouldn't have cancelled the store")
	//	}
	//})
}
