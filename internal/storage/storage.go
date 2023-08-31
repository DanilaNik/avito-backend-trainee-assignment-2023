package storage

import "errors"

var (
	ErrSegmentNotFound      = errors.New("Segment not found")
	ErrSegmentExists        = errors.New("Segment exists")
	ErrUserAlreadyInSegment = errors.New("User already in segment")
	ErrUserSegmentNotFound  = errors.New("User in segment not found")
)

type UserDTO struct {
	ID int64
}

type SegmentDTO struct {
	ID   int64
	Name string
}

type UserInSegmentDTO struct {
	ID      int64
	User    UserDTO
	Segment SegmentDTO
}

//type UsersSegmentsDTO struct {
//	User     UserDTO
//	Segments []SegmentDTO
//}
