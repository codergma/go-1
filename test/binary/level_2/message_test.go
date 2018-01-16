package test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/thrift-iterator/go"
	"github.com/thrift-iterator/go/protocol"
)

func Test_simple_message(t *testing.T) {
	should := require.New(t)
	buf := thrift.NewTMemoryBuffer()
	proto := thrift.NewTBinaryProtocol(buf, true, true)
	proto.WriteMessageBegin("hello", thrift.CALL, 17)
	proto.WriteStructBegin("args")
	proto.WriteFieldBegin("field1", thrift.I64, 1)
	proto.WriteI64(1)
	proto.WriteFieldBegin("field2", thrift.I64, 2)
	proto.WriteI64(2)
	proto.WriteFieldEnd()
	proto.WriteFieldStop()
	proto.WriteStructEnd()
	proto.WriteMessageEnd()
	iter := thrifter.NewIterator(buf.Bytes())
	msg := iter.ReadMessageHeader()
	should.Equal(protocol.VERSION_1, msg.Version)
	should.Equal("hello", msg.MessageName)
	should.Equal(protocol.CALL, msg.MessageType)
	should.Equal(protocol.SeqId(17), msg.SeqId)
}