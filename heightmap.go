package mc

type Heightmap [256]uint16

func (h Heightmap) HeightAt(x, z int) uint16 {
	x = x & 15
	z = z & 15
	return h[16 * z + x]
}

func (h *Heightmap) setHeightAt(x, z uint, y uint16) {
	if x > 16 || z > 16 {
		return
	}
	h[16 * z + x] = y
}

func HeightmapFromChunk(chunk Chunk) (res Heightmap) {
	max := 0
	for i := 15; i >= 0; i-- {
		if chunk.Sections[i] == nil {
			continue
		}
		if i > max {
			max = i
		}
	}
	s := chunk.Sections[max]
	for z := 0; z < 16; z ++ {
		for x := 0; x < 16; x ++ {
			hMax := 0
			for y := 15; y > 0; y-- {
				b := s.blockAt(BlockCoords{x, y, z})
				if b.ID != 0 {
					hMax = y
					break
				}
			}
			res.setHeightAt(uint(x), uint(z), uint16(16 * max + hMax))
		}
	}
	return
}