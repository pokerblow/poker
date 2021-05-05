package cardb

type Face rune
const (
	UnknownFace Face = 0
	Ace         Face = 'A'
	King        Face = 'K'
	Queen       Face = 'Q'
	Jack        Face = 'J'
	Ten         Face = 'T'
	Nine        Face = '9'
	Eight       Face = '8'
	Seven       Face = '7'
	Six         Face = '6'
	Five        Face = '5'
	Four  Face = '4'
	Three Face = '3'
	Deuce Face = '2'
)

func (f Face) Inc() Face {
	res, _ := f.Incremented()
	return res
}

func (f Face) Incremented() (Face, bool) {
	rank := FaceRank(f)
	if rank == MaxFaceRank {
		return f, false
	}
	return FaceFromRank(rank +1), true
}

func (f Face) Dec() Face {
	res, _ := f.Decremented()
	return res
}

func (f Face) Decremented() (Face, bool) {
	rank := FaceRank(f)
	if rank == MinFaceRank {
		return f, false
	}
	return FaceFromRank(rank -1), true
}

func (f Face) DecrementedStraight() (Face, bool) {
	rank := FaceRank(f)
	if rank == MinFaceRank {
		return Ace, true
	}
	return FaceFromRank(rank -1), true
}
