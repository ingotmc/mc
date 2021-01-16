package mc

import "math"

// Coords are a triplet of (x,y,z) float values.
type Coords vector3

// BlockCoords returns the floor of the (x,y,z) Coords values as int.
func (c Coords) BlockCoords() BlockCoords {
	return BlockCoords{
		X: int(math.Floor(c.X)),
		Y: int(math.Floor(c.Y)),
		Z: int(math.Floor(c.Z)),
	}
}

func (c Coords) ChunkCoords() ChunkCoords {
	return c.BlockCoords().ChunkCoords()
}

type BlockCoords struct {
	X, Y, Z int
}

func (b BlockCoords) ChunkCoords() ChunkCoords {
	return ChunkCoords{
		X: int32(b.X >> 4),
		Z: int32(b.Z >> 4),
	}
}

type ChunkCoords struct {
	X, Z int32
}

func (o ChunkCoords) Ring(radius int) []ChunkCoords {
	start := ChunkCoords{o.X + int32(-radius), o.Z + int32(radius)}
	length := 8 * radius
	segLen := 2 * radius
	coeff := ChunkCoords{1, 0}
	res := make([]ChunkCoords, length)
	for j := 0; j < 4; j++ {
		for i := 0; i < segLen; i++ {
			res[j*segLen+i] = ChunkCoords{
				start.X + coeff.X*int32(i),
				start.Z + coeff.Z*int32(i),
			}
		}
		last := res[segLen*(j+1)-1]
		start = ChunkCoords{last.X + coeff.X, last.Z + coeff.Z}
		switch j {
		case 0:
			coeff = ChunkCoords{0, -1}
		case 1:
			coeff = ChunkCoords{-1, 0}
		case 2:
			coeff = ChunkCoords{0, 1}
		}
	}
	return res
}

// Radius returns an array containing the ChunkCoords of an rxr area around this chunk
func (o ChunkCoords) Radius(r int) []ChunkCoords {
	region := make([]ChunkCoords, (r*2+1)*(r*2+1))
	region[0] = o
	chunkIdx := 1
	for i := 1; i <= r; i++ {
		ring := o.Ring(i)
		for _, c := range ring {
			region[chunkIdx] = c
			chunkIdx++
		}
	}
	return region
}
