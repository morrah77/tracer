package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
)

type Conf struct {
	listen string
}

var conf Conf
var logger *log.Logger

func init() {
	conf = Conf{}
	logger = log.New(os.Stdout, `Tracer `, log.Flags())
	flag.StringVar(&(conf.listen), `listen`, `8080`, `Port to listen`)
	flag.Parse()
}

func main() {
	println(`Listen at`, conf.listen)
	var (
		ln   net.Listener
		err  error
		conn net.Conn
	)
	addr := prepareAddress()
	logger.Println(addr)
	ln, err = net.Listen(`tcp`, addr)
	if err != nil {
		logger.Println(err)
		return
	}
	defer func() {
		ln.Close()
	}()
	for {
		conn, err = ln.Accept()
		if err != nil {
			logger.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	var (
		n   int
		buf []byte
		err error
	)
	buf = make([]byte, 32*1024)
	for {
		n, err = conn.Read(buf)
		if err == io.EOF {
			println(buf, string(buf), n)
			break
		}
		println(buf, string(buf), n)
	}
}
