package reader

// Record represents a ROS message record
type Record struct {
	headerTotalLength uint32
	header            []Header
	dataTotalLength   uint32
	data              []byte
}

// Header represents the header of a record
type Header struct {
	length uint32
	key    string
	value  interface{}
}

// BagHeader represents the header of bag file and is always the first record in bagfile.
// The bag header record is padded out by filling data with ASCII
// space characters (0x20) so that additional information can be
// added after the bag file is recorded. Currently, this padding is
// such that the header is 4096 bytes long.
type BagHeader struct {
	length     uint32 // Length of bag header
	indexPos   uint64
	connCount  uint32
	chunkCount uint32
	op         int8

	keyVal map[string]string // Optional key value pairs
}
