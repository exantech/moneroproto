package moneroproto

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPackVarint(t *testing.T) {
	tests := []struct {
		value    uint64
		expected []byte
	}{
		{0, []byte{0x00}},
		{1, []byte{0x04}},
		{2, []byte{0x08}},
		{63, []byte{0xfc}},
		{64, []byte{0x01, 0x01}},
		{16383, []byte{0xfd, 0xff}},
		{16384, []byte{0x02, 0x00, 0x01, 0x00}},
		{31337, []byte{0xa6, 0xe9, 0x01, 0x00}},
		{1073741823, []byte{0xfe, 0xff, 0xff, 0xff}},
		{1073741824, []byte{0x03, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00}},
		{4611686018427387903, []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
	}

	for _, test := range tests {
		buffer := bytes.Buffer{}
		packVarint(&buffer, test.value)
		assert.Equal(t, test.expected, buffer.Bytes(), "serialized values must be equal")
	}
}

func TestPackVarintFail(t *testing.T) {
	buffer := bytes.Buffer{}
	_, err := packVarint(&buffer, 4611686018427387904)
	assert.Error(t, err, "must return an error")
}

func TestUnpackVarint(t *testing.T) {
	tests := []struct {
		value    []byte
		expected uint64
	}{
		{[]byte{0x00}, 0},
		{[]byte{0x04}, 1},
		{[]byte{0x08}, 2},
		{[]byte{0xfc}, 63},
		{[]byte{0x01, 0x01}, 64},
		{[]byte{0xfd, 0xff}, 16383},
		{[]byte{0x02, 0x00, 0x01, 0x00}, 16384},
		{[]byte{0xa6, 0xe9, 0x01, 0x00}, 31337},
		{[]byte{0xfe, 0xff, 0xff, 0xff}, 1073741823},
		{[]byte{0x03, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00}, 1073741824},
		{[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, 4611686018427387903},
	}

	for _, test := range tests {
		reader := bytes.NewReader(test.value)
		actual, err := unpackVarint(reader)
		assert.Nil(t, err)
		assert.Equal(t, test.expected, actual, "deserialized values must be equal")
	}
}

func TestWriteUint64(t *testing.T) {
	expected := []byte{0x05, 0xef, 0xbe, 0xad, 0xde, 0xef, 0xbe, 0xad, 0xde}
	buffer := bytes.Buffer{}
	_, err := writeUint64(&buffer, 0xdeadbeefdeadbeef)

	assert.Nil(t, err)
	assert.Equal(t, expected, buffer.Bytes())
}

func TestReadUint64(t *testing.T) {
	reader := bytes.NewReader([]byte{0xef, 0xbe, 0xad, 0xde, 0xef, 0xbe, 0xad, 0xde})
	val, err := readUint64(reader)

	assert.Nil(t, err)
	assert.Equal(t, uint64(0xdeadbeefdeadbeef), val)
}


func TestWriteUint32(t *testing.T) {
	expected := []byte{0x06, 0xbe, 0xba, 0xad, 0xab}
	buffer := bytes.Buffer{}
	_, err := writeUint32(&buffer, 0xabadbabe)

	assert.Nil(t, err)
	assert.Equal(t, expected, buffer.Bytes())
}

func TestReadUint32(t *testing.T) {
	reader := bytes.NewReader([]byte{0xbe, 0xba, 0xad, 0xab})
	val, err := readUint32(reader)

	assert.Nil(t, err)
	assert.Equal(t, uint32(0xabadbabe), val)
}

func TestWriteUint32Short(t *testing.T) {
	expected := []byte{0x06, 0xbe, 0xba, 0x00, 0x00}
	buffer := bytes.Buffer{}
	_, err := writeUint32(&buffer, 0xbabe)

	assert.Nil(t, err)
	assert.Equal(t, expected, buffer.Bytes())
}

func TestReadUint32Short(t *testing.T) {
	reader := bytes.NewReader([]byte{0xbe, 0xba, 0x00, 0x00})
	val, err := readUint32(reader)

	assert.Nil(t, err)
	assert.Equal(t, uint32(0xbabe), val)
}

func TestWriteUint16(t *testing.T) {
	expected := []byte{0x07, 0xad, 0xde}
	buffer := bytes.Buffer{}
	_, err := writeUint16(&buffer, 0xdead)

	assert.Nil(t, err)
	assert.Equal(t, expected, buffer.Bytes())
}

func TestReadUint16(t *testing.T) {
	reader := bytes.NewReader([]byte{0xad, 0xde})
	val, err := readUint16(reader)

	assert.Nil(t, err)
	assert.Equal(t, uint16(0xdead), val)
}

func TestWriteUint8(t *testing.T) {
	expected := []byte{0x08, 0xad}
	buffer := bytes.Buffer{}
	_, err := writeUint8(&buffer, 0xad)

	assert.Nil(t, err)
	assert.Equal(t, expected, buffer.Bytes())
}

func TestReadUint8(t *testing.T) {
	reader := bytes.NewReader([]byte{0xad})
	val, err := readUint8(reader)

	assert.Nil(t, err)
	assert.Equal(t, uint8(0xad), val)
}

func TestWriteInt64(t *testing.T) {
	expected := []byte{0x01, 0xef, 0xbe, 0xad, 0xde, 0xef, 0xbe, 0xad, 0x7e}
	buffer := bytes.Buffer{}
	_, err := writeInt64(&buffer, 0x7eadbeefdeadbeef)

	assert.Nil(t, err)
	assert.Equal(t, expected, buffer.Bytes())
}

func TestReadInt64(t *testing.T) {
	reader := bytes.NewReader([]byte{0xef, 0xbe, 0xad, 0xde, 0xef, 0xbe, 0xad, 0x7e})
	val, err := readInt64(reader)

	assert.Nil(t, err)
	assert.Equal(t, int64(0x7eadbeefdeadbeef), val)
}

func TestWriteInt32(t *testing.T) {
	expected := []byte{0x02, 0xbe, 0xba, 0xad, 0x7b}
	buffer := bytes.Buffer{}
	_, err := writeInt32(&buffer, 0x7badbabe)

	assert.Nil(t, err)
	assert.Equal(t, expected, buffer.Bytes())
}

func TestReadInt32(t *testing.T) {
	reader := bytes.NewReader([]byte{0xbe, 0xba, 0xad, 0x7b})
	val, err := readInt32(reader)

	assert.Nil(t, err)
	assert.Equal(t, int32(0x7badbabe), val)
}

func TestWriteInt32Negative(t *testing.T) {
	expected := []byte{0x02, 0xf0, 0xd8, 0xff, 0xff}
	buffer := bytes.Buffer{}
	_, err := writeInt32(&buffer, -10000)

	assert.Nil(t, err)
	assert.Equal(t, expected, buffer.Bytes())
}

func TestReadInt32Negative(t *testing.T) {
	reader := bytes.NewReader([]byte{0xf0, 0xd8, 0xff, 0xff})
	val, err := readInt32(reader)

	assert.Nil(t, err)
	assert.Equal(t, int32(-10000), val)
}

func TestWriteInt16(t *testing.T) {
	expected := []byte{0x03, 0xad, 0x7e}
	buffer := bytes.Buffer{}
	_, err := writeInt16(&buffer, 0x7ead)

	assert.Nil(t, err)
	assert.Equal(t, expected, buffer.Bytes())
}

func TestReadInt16(t *testing.T) {
	reader := bytes.NewReader([]byte{0xad, 0x7e})
	val, err := readInt16(reader)

	assert.Nil(t, err)
	assert.Equal(t, int16(0x7ead), val)
}

func TestWriteInt8(t *testing.T) {
	expected := []byte{0x04, 0x7d}
	buffer := bytes.Buffer{}
	_, err := writeInt8(&buffer, 0x7d)

	assert.Nil(t, err)
	assert.Equal(t, expected, buffer.Bytes())
}

func TestReadInt8(t *testing.T) {
	reader := bytes.NewReader([]byte{0x7d})
	val, err := readInt8(reader)

	assert.Nil(t, err)
	assert.Equal(t, int8(0x7d), val)
}

func TestWriteBoolTrue(t *testing.T) {
	expected := []byte{0x0b, 0x01}
	buffer := bytes.Buffer{}
	_, err := writeBool(&buffer, true)

	assert.Nil(t, err)
	assert.Equal(t, expected, buffer.Bytes())
}

func TestReadBoolTrue(t *testing.T) {
	reader := bytes.NewReader([]byte{0x01})
	val, err := readBool(reader)

	assert.Nil(t, err)
	assert.Equal(t, true, val)
}

func TestWriteBoolFalse(t *testing.T) {
	expected := []byte{0x0b, 0x00}
	buffer := bytes.Buffer{}
	_, err := writeBool(&buffer, false)

	assert.Nil(t, err)
	assert.Equal(t, expected, buffer.Bytes())
}

func TestReadBoolFalse(t *testing.T) {
	reader := bytes.NewReader([]byte{0x0b, 0x00})
	val, err := readBool(reader)

	assert.Nil(t, err)
	assert.Equal(t, false, val)
}

func TestWriteDouble(t *testing.T) {
	expected := []byte{0x09, 0x2a, 0x80, 0x6f, 0xfc, 0x8c, 0x78, 0xe2, 0x3f}
	buffer := bytes.Buffer{}
	_, err := writeFloat64(&buffer, 0.5772156649)

	assert.Nil(t, err)
	assert.Equal(t, expected, buffer.Bytes())
}

func TestReadDouble(t *testing.T) {
	reader := bytes.NewReader([]byte{0x2a, 0x80, 0x6f, 0xfc, 0x8c, 0x78, 0xe2, 0x3f})
	val, err := readFloat64(reader)

	assert.Nil(t, err)
	assert.Equal(t, float64(0.5772156649), val)
}

func TestWriteBinaryString(t *testing.T) {
	expected := []byte{0x0a, 0x2c, 0x73, 0x61, 0x79, 0x20, 0x6d, 0x79, 0x20, 0x6e, 0x61, 0x6d, 0x65}
	buffer := bytes.Buffer{}
	_, err := writeBinaryString(&buffer, []byte("say my name"))

	assert.Nil(t, err)
	assert.Equal(t, expected, buffer.Bytes())
}

func TestReadBinaryString(t *testing.T) {
	reader := bytes.NewReader([]byte{0x2c, 0x73, 0x61, 0x79, 0x20, 0x6d, 0x79, 0x20, 0x6e, 0x61, 0x6d, 0x65})
	val, err := readBinaryString(reader)

	assert.Nil(t, err)
	assert.Equal(t, []byte("say my name"), val)
}
