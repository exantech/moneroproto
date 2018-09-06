package moneroproto

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/cobinhood/moneroutil"
)

func hashFromString(str string) moneroutil.Hash {
	_, hash := NewHashFromHexStr([]byte(str))
	return *hash
}

var hash1 = hashFromString("112233445566778899aabbccddeeff00112233445566778899aabbccddeeff00")
var hash2 = hashFromString("00ffeeddccbbaa99887766554433221100ffeeddccbbaa998877665544332211")

func TestGetHashesFastRequestEncode(t *testing.T) {
	expected := []byte{0x01, 0x11, 0x01, 0x01, 0x01, 0x01, 0x02, 0x01, 0x01, 0x08, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
		0x5f, 0x69, 0x64, 0x73, 0x0a, 0x01, 0x01, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb,
		0xcc, 0xdd, 0xee, 0xff, 0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd,
		0xee, 0xff, 0x00, 0x00, 0xff, 0xee, 0xdd, 0xcc, 0xbb, 0xaa, 0x99, 0x88, 0x77, 0x66, 0x55, 0x44, 0x33, 0x22,
		0x11, 0x00, 0xff, 0xee, 0xdd, 0xcc, 0xbb, 0xaa, 0x99, 0x88, 0x77, 0x66, 0x55, 0x44, 0x33, 0x22, 0x11, 0x0c,
		0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x05, 0xbe, 0xba, 0xad, 0xde, 0xef,
		0xbe, 0xad, 0xde}

	obj := GetHashesFastRequest {
		StartHeight: uint64(0xdeadbeefdeadbabe),
	}
	obj.SetHashes([]moneroutil.Hash{hash1, hash2})

	buffer := bytes.Buffer{}
	err := Write(&buffer, obj)

	assert.Nil(t, err)
	assert.Equal(t, expected, buffer.Bytes())
}

func TestGetHashesFastRequestDecode(t *testing.T) {
	expected := GetHashesFastRequest{
		StartHeight: uint64(0xdeadbeefdeadbabe),
	}
	expected.SetHashes([]moneroutil.Hash{hash1, hash2})

	reader := bytes.NewReader([]byte{0x01, 0x11, 0x01, 0x01, 0x01, 0x01, 0x02, 0x01, 0x01, 0x08, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
		0x5f, 0x69, 0x64, 0x73, 0x0a, 0x01, 0x01, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb,
		0xcc, 0xdd, 0xee, 0xff, 0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd,
		0xee, 0xff, 0x00, 0x00, 0xff, 0xee, 0xdd, 0xcc, 0xbb, 0xaa, 0x99, 0x88, 0x77, 0x66, 0x55, 0x44, 0x33, 0x22,
		0x11, 0x00, 0xff, 0xee, 0xdd, 0xcc, 0xbb, 0xaa, 0x99, 0x88, 0x77, 0x66, 0x55, 0x44, 0x33, 0x22, 0x11, 0x0c,
		0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x05, 0xbe, 0xba, 0xad, 0xde, 0xef,
		0xbe, 0xad, 0xde})

	var obj GetHashesFastRequest
	err := Read(reader, &obj)

	assert.Nil(t, err)
	assert.Equal(t, expected, obj)
}

func TestGetHashesFastResponseDecode(t *testing.T) {
	expected := GetHashesFastResponse{
		StartHeight: uint64(0xdeadbeefdeadbabe),
		CurrentHeight: uint64(0xdeadbeefdeadbaff),
		Status: []byte("coolio"),
		Untrusted: true,
	}
	expected.SetHashes([]moneroutil.Hash{hash1, hash2})

	// obtained from monero
	reader := bytes.NewReader([]byte{0x01, 0x11, 0x01, 0x01, 0x01, 0x01, 0x02, 0x01, 0x01, 0x14, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65,
		0x6e, 0x74, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x05, 0xff, 0xba, 0xad, 0xde, 0xef, 0xbe, 0xad, 0xde, 0x0b,
		0x6d, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x73, 0x0a, 0x01, 0x01, 0x11, 0x22, 0x33, 0x44, 0x55,
		0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff, 0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88,
		0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff, 0x00, 0x00, 0xff, 0xee, 0xdd, 0xcc, 0xbb, 0xaa, 0x99, 0x88, 0x77, 0x66,
		0x55, 0x44, 0x33, 0x22, 0x11, 0x00, 0xff, 0xee, 0xdd, 0xcc, 0xbb, 0xaa, 0x99, 0x88, 0x77, 0x66, 0x55, 0x44, 0x33,
		0x22, 0x11, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x05, 0xbe, 0xba, 0xad,
		0xde, 0xef, 0xbe, 0xad, 0xde, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x0a, 0x18, 0x63, 0x6f, 0x6f, 0x6c, 0x69,
		0x6f, 0x09, 0x75, 0x6e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x0b, 0x01})

	var obj GetHashesFastResponse
	err := Read(reader, &obj)

	assert.Nil(t, err)
	assert.Equal(t, expected, obj)
}

func TestGetHashesFastResponseSerialization(t *testing.T) {
	// since this structure hash several fields there's no sense to compare obtained from monero bytes as is
	// because this serializer encodes fields in different order. we encode and decode the structure to make sure
	// we are able to serialize it properly

	primary := GetHashesFastResponse{
		StartHeight: uint64(0xdeadbeefdeadbabe),
		CurrentHeight: uint64(0xdeadbeefdeadbaff),
		Status: []byte("coolio"),
		Untrusted: true,
	}
	primary.SetHashes([]moneroutil.Hash{hash1, hash2})

	buffer := bytes.Buffer{}
	err := Write(&buffer, primary)
	assert.Nil(t, err)

	reader := bytes.NewReader(buffer.Bytes())

	var restored GetHashesFastResponse
	err = Read(reader, &restored)

	assert.Nil(t, err)
	assert.Equal(t, primary, restored)
}

func TestGetBlocksFastRequestDecode(t *testing.T) {
	expected := GetBlocksFastRequest {
		StartHeight: uint64(0xdeadbeefdeadbabe),
		Prune: true,
		NoMinerTx: false,
	}
	expected.SetHashes([]moneroutil.Hash{hash1, hash2})

	reader := bytes.NewReader([]byte{0x01, 0x11, 0x01, 0x01, 0x01, 0x01, 0x02, 0x01, 0x01, 0x10, 0x09, 0x62, 0x6c, 0x6f,
		0x63, 0x6b, 0x5f, 0x69, 0x64, 0x73, 0x0a, 0x01, 0x01, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa,
		0xbb, 0xcc, 0xdd, 0xee, 0xff, 0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd,
		0xee, 0xff, 0x00, 0x00, 0xff, 0xee, 0xdd, 0xcc, 0xbb, 0xaa, 0x99, 0x88, 0x77, 0x66, 0x55, 0x44, 0x33, 0x22, 0x11,
		0x00, 0xff, 0xee, 0xdd, 0xcc, 0xbb, 0xaa, 0x99, 0x88, 0x77, 0x66, 0x55, 0x44, 0x33, 0x22, 0x11, 0x0b, 0x6e, 0x6f,
		0x5f, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x74, 0x78, 0x0b, 0x00, 0x05, 0x70, 0x72, 0x75, 0x6e, 0x65, 0x0b, 0x01,
		0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x05, 0xbe, 0xba, 0xad, 0xde, 0xef,
		0xbe, 0xad, 0xde})

	var obj GetBlocksFastRequest
	err := Read(reader, &obj)

	assert.Nil(t, err)
	assert.Equal(t, expected, obj)
}

func TestGetBlocksFastRequestSerialize(t *testing.T) {
	expected := GetBlocksFastRequest {
		StartHeight: uint64(0xdeadbeefdeadbabe),
		Prune: true,
		NoMinerTx: false,
	}
	expected.SetHashes([]moneroutil.Hash{hash1, hash2})

	buf := bytes.Buffer{}
	err := Write(&buf, expected)
	assert.Nil(t, err)

	var obj GetBlocksFastRequest
	reader := bytes.NewReader(buf.Bytes())
	err = Read(reader, &obj)

	assert.Nil(t, err)
	assert.Equal(t, expected, obj)
}

var expectedGetBlocksFastResponse = GetBlocksFastResponse {
	Blocks: []BlockCompleteEntry {
		BlockCompleteEntry {
			Block: []byte("AAAblockAAA"),
			Txs: [][]byte{
				[]byte("tx1"), []byte("tx2"), []byte("tx3"),
			},
		},
		BlockCompleteEntry {
			Block: []byte("BBBblockBBB"),
			Txs: [][]byte{
				[]byte("Btx1"), []byte("Btx2"), []byte("Btx3"),
			},
		},
	},
	StartHeight: 112233,
	CurrentHeight: 445566,
	Status: []byte("hell!"),
	Untrusted: true,
	OutputIndices: []BlockOutputIndices {
		BlockOutputIndices {
			Indices: []TxOutputIndices {
				TxOutputIndices {
					Indices: []uint64{1, 2, 3, 4, 5},
				},
				TxOutputIndices {
					Indices: []uint64{6, 7, 8},
				},
			},
		},
		BlockOutputIndices {
			Indices: []TxOutputIndices {
				TxOutputIndices {
					Indices: []uint64{9, 10, 11, 12, 13},
				},
				TxOutputIndices {
					Indices: []uint64{14, 15, 16, 17, 18, 19},
				},
			},
		},
	},
}

func TestGetBlocksFastResponseDecode(t *testing.T) {
	reader := bytes.NewReader([]byte{0x01, 0x11, 0x01, 0x01, 0x01, 0x01, 0x02, 0x01, 0x01, 0x18, 0x06, 0x62, 0x6c, 0x6f,
		0x63, 0x6b, 0x73, 0x8c, 0x08, 0x08, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x0a, 0x2c, 0x41, 0x41, 0x41, 0x62, 0x6c,
		0x6f, 0x63, 0x6b, 0x41, 0x41, 0x41, 0x03, 0x74, 0x78, 0x73, 0x8a, 0x0c, 0x0c, 0x74, 0x78, 0x31, 0x0c, 0x74, 0x78,
		0x32, 0x0c, 0x74, 0x78, 0x33, 0x08, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x0a, 0x2c, 0x42, 0x42, 0x42, 0x62, 0x6c,
		0x6f, 0x63, 0x6b, 0x42, 0x42, 0x42, 0x03, 0x74, 0x78, 0x73, 0x8a, 0x0c, 0x10, 0x42, 0x74, 0x78, 0x31, 0x10, 0x42,
		0x74, 0x78, 0x32, 0x10, 0x42, 0x74, 0x78, 0x33, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x65,
		0x69, 0x67, 0x68, 0x74, 0x05, 0x7e, 0xcc, 0x06, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0e, 0x6f, 0x75, 0x74, 0x70, 0x75,
		0x74, 0x5f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x8c, 0x08, 0x04, 0x07, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65,
		0x73, 0x8c, 0x08, 0x04, 0x07, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x85, 0x14, 0x01, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0x07, 0x69,
		0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x85, 0x0c, 0x06, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x07, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0x07, 0x69, 0x6e, 0x64, 0x69,
		0x63, 0x65, 0x73, 0x8c, 0x08, 0x04, 0x07, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x85, 0x14, 0x09, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04,
		0x07, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x85, 0x18, 0x0e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0f,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x11, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x12, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x13, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x05, 0x69, 0xb6, 0x01, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x0a, 0x14, 0x68, 0x65, 0x6c, 0x6c, 0x21, 0x09,
		0x75, 0x6e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x0b, 0x01})

	var obj GetBlocksFastResponse
	err := Read(reader, &obj)

	assert.Nil(t, err)
	assert.Equal(t, expectedGetBlocksFastResponse, obj)
}

func TestGetBlocksFastResponseSerialize(t *testing.T) {
	buf := bytes.Buffer{}
	err := Write(&buf, &expectedGetBlocksFastResponse)
	assert.Nil(t, err)

	var obj GetBlocksFastResponse
	reader := bytes.NewReader(buf.Bytes())
	err = Read(reader, &obj)

	assert.Nil(t, err)
	assert.Equal(t, expectedGetBlocksFastResponse, obj)
}