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
		Z: int32(b.Z >> 16),
	}
}

type ChunkCoords struct {
	X, Z int32
}

// Radius returns an array containing the ChunkCoords of an rxr area around this chunk
func (orig ChunkCoords) Radius(r int) []ChunkCoords {
	region := make([]ChunkCoords, (r*2+1)*(r*2+1))
	i := 0
	a := int32(r)
	for x := orig.X - a; x <= orig.X+a; x++ {
		for z := orig.Z - a; z <= orig.Z+a; z++ {
			region[i] = ChunkCoords{x, z}
			i++
		}
	}
	return region
}
