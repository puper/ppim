package hash

import (
	"hash/crc32"
)

func HashBucketId(userId string, numBuckets uint32) uint32 {
	return crc32.ChecksumIEEE([]byte(userId)) % numBuckets
}

func HashServerId(bucketId uint32, numServers uint32) uint32 {
	var b int64 = -1
	var j int64
	key := uint64(bucketId)
	numBuckets := int64(numServers)
	for j < numBuckets {
		b = j
		key = key*2862933555777941757 + 1
		j = int64(float64(b+1) * (float64(int64(1)<<31) / float64((key>>33)+1)))
	}

	return uint32(b)
}
