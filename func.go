import (
	"hash/crc32"
	"strconv"
)

func (d *Decider) withinPercentile(id uint64, val float64, feature string) bool {
	uid := d.crc(id, feature)
	percentage := uint32(val * 100)

	return uid%100 < percentage
}

func (d *Decider) crc(id uint64, feature string) uint32 {
	bytesToHash := []byte(feature + strconv.FormatInt(int64(id), 10))

	return crc32.ChecksumIEEE(bytesToHash)
}
