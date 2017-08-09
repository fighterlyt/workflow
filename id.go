package workflow

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"sync/atomic"
	"time"
)

type IdGenerator interface {
	Generate() string
}

var (
	idGenerator IdGenerator = newObjectIdGenerator()
)

func SetIdGenerator(generator IdGenerator) {
	idGenerator = generator
}

type objectIdGenerator struct {
	machineId       []byte
	processId       int
	objectIdCounter uint32
}

func newObjectIdGenerator() *objectIdGenerator {
	return &objectIdGenerator{
		machineId:       readMachineId(),
		processId:       os.Getpid(),
		objectIdCounter: readRandomUint32(),
	}

}
func (o *objectIdGenerator) Generate() string {
	var b [12]byte
	// Timestamp, 4 bytes, big endian
	binary.BigEndian.PutUint32(b[:], uint32(time.Now().Unix()))
	// Machine, first 3 bytes of md5(hostname)
	b[4] = o.machineId[0]
	b[5] = o.machineId[1]
	b[6] = o.machineId[2]
	// Pid, 2 bytes, specs don't specify endianness, but we use big endian.
	b[7] = byte(o.processId >> 8)
	b[8] = byte(o.processId)
	// Increment, 3 bytes, big endian
	i := atomic.AddUint32(&o.objectIdCounter, 1)
	b[9] = byte(i >> 16)
	b[10] = byte(i >> 8)
	b[11] = byte(i)
	return string(b[:])
}

func readMachineId() []byte {
	var sum [3]byte
	id := sum[:]
	hostname, err1 := os.Hostname()
	if err1 != nil {
		_, err2 := io.ReadFull(rand.Reader, id)
		if err2 != nil {
			panic(fmt.Errorf("cannot get hostname: %v; %v", err1, err2))
		}
		return id
	}
	hw := md5.New()
	hw.Write([]byte(hostname))
	copy(id, hw.Sum(nil))
	return id
}
func readRandomUint32() uint32 {
	var b [4]byte
	_, err := io.ReadFull(rand.Reader, b[:])
	if err != nil {
		panic(fmt.Errorf("cannot read random object id: %v", err))
	}
	return uint32((uint32(b[0]) << 0) | (uint32(b[1]) << 8) | (uint32(b[2]) << 16) | (uint32(b[3]) << 24))
}
