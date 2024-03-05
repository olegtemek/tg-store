package folder

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/olegtemek/tg-store/internal/utils"
	pb "github.com/olegtemek/tg-store/proto"
)

type Repository struct {
	log *slog.Logger
	db  *pgxpool.Pool
}

func NewRepository(log *slog.Logger, db *pgxpool.Pool) *Repository {
	return &Repository{
		log: log,
		db:  db,
	}
}

func (r *Repository) Create(title string, userId int) (*pb.FolderModel, *utils.WrappError) {
	folder := &pb.FolderModel{}
	q := `INSERT INTO Folders (title, userId) VALUES ($1, $2) RETURNING id, title, userId`
	err := r.db.QueryRow(context.Background(), q, title, userId).Scan(&folder.Id, &folder.Title, &folder.UserId)
	if err != nil {
		r.log.Error("ERROR",
			slog.String("SQL", q),
			slog.Any("ERROR", err),
		)
		return nil, &utils.WrappError{Err: err}
	}
	return folder, nil
}

func (r *Repository) Update(id int, title string, userId int) (*pb.FolderModel, *utils.WrappError) {
	folder := &pb.FolderModel{}
	q := `UPDATE Folders SET title = $3 WHERE id = $1 AND userId = $2 RETURNING id, title, userId`
	err := r.db.QueryRow(context.Background(), q, id, userId, title).Scan(&folder.Id, &folder.Title, &folder.UserId)
	if err != nil {
		r.log.Error("ERROR",
			slog.String("SQL", q),
			slog.Any("ERROR", err),
		)
		return nil, &utils.WrappError{Err: err}
	}
	return folder, nil
}

func (r *Repository) GetAll(userId int) ([]*pb.FolderModel, *utils.WrappError) {
	folders := []*pb.FolderModel{}
	q := `
		SELECT fol.id, fol.title, fol.userId, COUNT(fi)
		FROM Folders AS fol 
		LEFT JOIN Files as fi ON fol.id = fi.folderId 
		WHERE fol.userId = $1 
		GROUP BY fol.id, fol.title, fol.userId`

	rows, err := r.db.Query(context.Background(), q, userId)
	if err != nil {
		r.log.Error("ERROR",
			slog.String("SQL", q),
			slog.Any("ERROR", err),
		)
		return nil, &utils.WrappError{Err: err}
	}
	defer rows.Close()

	for rows.Next() {
		folder := &pb.FolderModel{}
		err := rows.Scan(&folder.Id, &folder.Title, &folder.UserId, &folder.Files)

		if err != nil {
			r.log.Error("ERROR",
				slog.String("SQL", q),
				slog.Any("ERROR", err),
			)
			return nil, &utils.WrappError{Err: err}
		}

		folders = append(folders, folder)
	}

	return folders, nil
}

func (r *Repository) Delete(id int, userId int) (*pb.FolderModel, *utils.WrappError) {
	folder := &pb.FolderModel{}
	q := `DELETE FROM Folders WHERE id = $1 AND userId = $2 RETURNING id, title, userId`
	err := r.db.QueryRow(context.Background(), q, id, userId).Scan(&folder.Id, &folder.Title, &folder.UserId)
	if err != nil {
		r.log.Error("ERROR",
			slog.String("SQL", q),
			slog.Any("ERROR", err),
		)
		return nil, &utils.WrappError{Err: err}
	}
	return folder, nil
}
