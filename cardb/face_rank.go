package cardb

const MaxFaceRank = 12
const MinFaceRank = 0

var facesAsc = []Face{Deuce, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

var faceRanks = make(map[Face]int, len(facesAsc))
var rankFaces = make(map[int]Face, len(facesAsc))

func init() {
	for i, face := range facesAsc {
		faceRanks[face] = i
	}
	for i, face := range facesAsc {
		rankFaces[i]=face
	}
}

func FaceRank(f Face) int {
	return faceRanks[f]
}

func FaceFromRank(rank int) Face {
	if rank > MaxFaceRank {
		rank = MaxFaceRank
	}
	if rank < MinFaceRank {
		rank = MinFaceRank
	}
	return rankFaces[rank]
}

