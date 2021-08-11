package third_party

func CalcCrc8(in string) int {
	var (
		crc byte
		buf = []byte(in)
	)
	for i := len(buf) - 1; i >= 0; i-- {
		crc ^= buf[i]
		for j := 8; j > 0; j-- {
			if crc&0x80 > 0 {
				crc = (crc << 1) ^ 0x31
			} else {
				crc = crc << 1
			}
		}
	}

	return int(crc)
}