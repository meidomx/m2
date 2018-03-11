package logserver

import (
	"github.com/atnet/gof/log"
	"io"
	"net"
	"os"
	"path/filepath"
	"time"
)

import (
	gobytes "moetang.info/go/nekoq/util/bytes"
)

const BASE_DIR string = "./LogData/"

const BUF_SIZE int = 4096
const QUEUE_SIZE int = 1024

const (
	BIN_FILE_NAME string = "log.bin"
)

var (
	loopFlag = true
)

type LogEntry struct {
	Version byte
	LogTime int64
	Data    []byte

	underlyingBuf []byte
}

func (this *LogEntry) Reset() {
	this.Version = 0
	this.LogTime = 0
	this.Data = this.underlyingBuf
}

func (this *LogEntry) Write(w io.Writer) error {
	_, err := w.Write([]byte{this.Version, 0, 0, 0})
	if err != nil {
		return err
	}

	err = gobytes.WriteInt64BE(w, this.LogTime)
	if err != nil {
		return err
	}

	err = gobytes.WriteInt64BE(w, int64(len(this.Data))&int64(0xFFFFFFFF))
	if err != nil {
		return err
	}``

	_, err = w.Write(this.Data)
	if err != nil {
		return err
	}

	return nil
}

func Start() {
	baseDir, err := os.Open(BASE_DIR)
	if os.IsNotExist(err) {
		err = os.MkdirAll(BASE_DIR, os.FileMode(0755))
		if err != nil {
			panic(err)
		}
		baseDir, err = os.Open(BASE_DIR)
		if err != nil {
			panic(err)
		}
	}

	addr, err := net.ResolveUDPAddr("udp", ":10000")
	if err != nil {
		panic(err)
	}
	udpConn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}

	closeChan := make(chan bool, 16)

	logChan := make(chan *LogEntry, QUEUE_SIZE)
	bufChan := make(chan *LogEntry, QUEUE_SIZE)
	for i := 0; i < QUEUE_SIZE; i++ {
		entry := new(LogEntry)
		entry.underlyingBuf = make([]byte, BUF_SIZE)
		entry.Reset()
		bufChan <- entry
	}
	go saveToFile(baseDir, logChan, bufChan)
	go processUdpData(udpConn, logChan, bufChan)

	log.Println("start complete.")
	<-closeChan
}

func processUdpData(udpConn *net.UDPConn, logChan chan<- *LogEntry, bufChan <-chan *LogEntry) {
READ_LOOP:
	for loopFlag {
		entry := <-bufChan
		entry.Reset()
		n, _, err := udpConn.ReadFromUDP(entry.Data)
		if err != nil {
			log.Println(err)
			break READ_LOOP
		}
		entry.Data = entry.Data[:n]
		logChan <- entry
	}
}

func saveToFile(folder *os.File, logChan <-chan *LogEntry, bufChan chan<- *LogEntry) {
	parent, err := filepath.Abs(folder.Name())
	if err != nil {
		panic(err)
	}

	binfile, err := os.OpenFile(filepath.Join(parent, "logdata.bin"), os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}

	for loopFlag {
		entry := <-logChan

		//TODO src: \t分割, sA, sB, sC, s0, s1, ... , s7, s8，且数字类型为ASCII数字字符串
		//TODO dst: s0动作, s1地图, s2X, s3Y, s4人物名称, s5物品名称, s6物品ID, s7记录, s8交易对象, (时间)
		entry.Version = 1
		entry.LogTime = time.Now().UnixNano()
		entry.Write(binfile)

		bufChan <- entry
	}
}
