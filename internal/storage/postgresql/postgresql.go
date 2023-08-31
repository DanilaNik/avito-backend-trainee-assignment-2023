package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/DanilaNik/avito-backend-trainee-assignment-2023/internal/config"
	"github.com/DanilaNik/avito-backend-trainee-assignment-2023/internal/storage"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New(cfg config.Storage) (*Storage, error) {
	const op = "storage.postgresql.New"

	dataSource := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s  sslmode=%s",
		cfg.Addr, cfg.Port, cfg.User, cfg.DB, cfg.Password, cfg.Sslmode,
	)
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveUser() (*storage.UserDTO, error) {
	const op = "storage.postgresql.SaveUser"

	stmt, err := s.db.Prepare("INSERT INTO users(id) VALUES(DEFAULT) RETURNING *")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var user storage.UserDTO
	err = stmt.QueryRow().Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to get last insert id %w", op, err)
	}
	return &user, nil
}

func (s *Storage) DeleteUser(userId int64) error {
	const op = "storage.postgresql.DeleteUser"

	stmt, err := s.db.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(userId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Storage) SaveSegment(name string) (*storage.SegmentDTO, error) {
	const op = "storage.postgresql.SaveSegment"

	stmt, err := s.db.Prepare("INSERT INTO segments(name) VALUES($1) RETURNING *")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var segment storage.SegmentDTO
	err = stmt.QueryRow(name).Scan(&segment.ID, &segment.Name)
	if err != nil {
		if err, ok := err.(*pq.Error); ok && err.Code == "23505" {
			return nil, storage.ErrSegmentExists
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &segment, nil
}

func (s *Storage) DeleteSegment(name string) (*storage.SegmentDTO, error) {
	const op = "storage.postgresql.DeleteSegment"

	stmt, err := s.db.Prepare("DELETE FROM segments WHERE name = $1 RETURNING *")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var segment storage.SegmentDTO
	err = stmt.QueryRow(name).Scan(&segment.ID, &segment.Name)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &segment, nil
}

func (s *Storage) AddUserToSegment(userId int64, segmentName string) (*storage.UserInSegmentDTO, error) {
	const op = "storage.postgresql.AddUserToSegment"

	stmt, err := s.db.Prepare("INSERT INTO user_segments() VALUES($1) RETURNING *")
	var userInSegment storage.UserInSegmentDTO

	return &userInSegment, nil
}
