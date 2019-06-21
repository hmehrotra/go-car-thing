package reader

import (
	"encoding/binary"
	"fmt"
	"os"
	"strings"

	"github.com/go-car-thing/src/utils"
)

func printDebugString(str string, debug bool) {
	if debug {
		fmt.Println(str)
	}
}

func readBagFileHeader(f *os.File, h *BagHeader, debug bool) {

	h.length = utils.ReadInt32(f)
	bytesToRead := h.length

	printDebugString("Bag File Header", debug)
	printDebugString("----------------------------------------------------------", debug)

	for bytesToRead > 0 {
		length := utils.ReadInt32(f)

		keyVal := make([]byte, length)
		_, err := f.Read(keyVal)
		utils.Check(err)

		pairs := strings.Split(string(keyVal), "=")
		if pairs[0] == "index_pos" {
			h.indexPos = binary.LittleEndian.Uint64(keyVal[length-8:])
			printDebugString(fmt.Sprintf("Index Pos: %d", h.indexPos), debug)
		} else if pairs[0] == "conn_count" {
			h.connCount = binary.LittleEndian.Uint32(keyVal[length-4:])
			printDebugString(fmt.Sprintf("Conn Count: %d", h.connCount), debug)
		} else if pairs[0] == "chunk_count" {
			h.chunkCount = binary.LittleEndian.Uint32(keyVal[length-4:])
			printDebugString(fmt.Sprintf("Chunk Count: %d", h.chunkCount), debug)
		} else if pairs[0] == "op" {
			h.op = int8(keyVal[length-1])
			printDebugString(fmt.Sprintf("Op code: %d", h.op), debug)
		} else {
			// Generic key value pair in header
			if h.keyVal == nil {
				h.keyVal = make(map[string]string)
			}
			h.keyVal[pairs[0]] = pairs[1]
			fmt.Println(pairs[0])
		}
		bytesToRead = bytesToRead - length - 4
	}
}

// ReadRosBag reads the rosbag file
func ReadRosBag(f *os.File, debug bool) {

	// Read the title of ROSbag file
	bagFileHeader := make([]byte, 13)
	_, err := f.Read(bagFileHeader)
	utils.Check(err)

	// Read the bag header record
	bagHeader := &BagHeader{}
	readBagFileHeader(f, bagHeader, debug)
}
