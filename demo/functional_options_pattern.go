package main

import (
	"fmt"
	"github.com/pkg/errors"
	"net"
	"time"
)

const defaultRate = 100 * time.Millisecond

type Subscriber struct {
	conn net.Conn
	rate time.Duration
}

func New(options ...Option) (*Subscriber, error) {
	// set defaults
	p := &Subscriber{
		rate: defaultRate,
	}
	// apply options
	for _, opt := range options {
		err := opt(p)
		if err != nil {
			return nil, errors.Wrap(err, "unable to apply option")
		}
	}
	// validate
	if p.rate < 0 {
		return nil, errors.Errorf("negative rate given")
	}

	return p, nil
}

type Option func(*Subscriber) error

func WithConnection(conn net.Conn) Option {
	return func(p *Subscriber) error {
		p.conn = conn
		return nil
	}
}

func WithAddress(addr string) Option {
	return func(p *Subscriber) error {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			return errors.Wrap(err, "unable to connect")
		}
		p.conn = conn
		return nil
	}
}

func WithRate(rate time.Duration) Option {
	return func(p *Subscriber) error {
		p.rate = rate
		return nil
	}
}

func main() {
	_, err := New(WithAddress("127.0.0.1:5000"), WithRate(5*time.Second))
	fmt.Println(err)
	_, err = New()
	fmt.Println(err)
	_, err = New(WithRate(-1))
	fmt.Println(err)
}
