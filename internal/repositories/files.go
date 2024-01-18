package repositories

import (
	"context"
	"encoding/hex"
	"github.com/WildEgor/g-cdn/internal/db"
	"github.com/WildEgor/g-cdn/internal/models"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type FileRepository struct {
	name string
	db   *db.MongoDBConnection
}

func NewFileRepository(
	db *db.MongoDBConnection,
) *FileRepository {

	return &FileRepository{
		db.DbName(),
		db,
	}
}

func (r *FileRepository) AddFile(filename string, checksum []byte) (string, error) {
	model := &models.FileModel{
		Name:      filename,
		CheckSum:  hex.EncodeToString(checksum),
		Status:    models.ActiveStatus,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	filter := bson.D{{"check_sum", bson.D{{"$eq", model.CheckSum}}}}

	existed := r.db.Instance().Database(r.name).Collection(models.CollectionFiles).FindOne(context.Background(), filter) // TODO: ctx

	if existed != nil {
		err := existed.Decode(&model)
		if err == nil {
			return model.Name, nil
		}
	}

	_, err := r.db.Instance().Database(r.name).Collection(models.CollectionFiles).InsertOne(context.Background(), model) // TODO: ctx
	if err != nil {
		return "", errors.New(`[AddFile] err insert`) // TODO
	}

	return model.Name, nil
}

func (r *FileRepository) DeleteFile(filename string) (string, error) {
	var model *models.FileModel
	filter := bson.D{{Key: "file_name", Value: filename}}
	err := r.db.Instance().Database(r.name).Collection(models.CollectionFiles).FindOne(context.Background(), filter).Decode(model) // TODO: ctx
	if err != nil {
		return "", errors.New(`Mongo error`) // TODO
	}

	if model != nil {
		_, err := r.db.Instance().Database(r.name).Collection(models.CollectionFiles).DeleteOne(context.Background(), bson.D{{Key: "_id", Value: model.Id}}) // TODO: ctx
		if err != nil {
			return "", errors.New(`Mongo error`) // TODO
		}
	}

	return filename, nil
}

func (r *FileRepository) RenameFile(oldname, newname string) error {
	update := bson.D{
		{"$set",
			bson.D{
				{"file_name", newname},
				{"updated_at", time.Now().UTC()},
			},
		},
	}

	_, err := r.db.Instance().Database(r.name).Collection(models.CollectionFiles).UpdateOne(context.Background(), bson.D{{Key: "file_name", Value: oldname}}, update) // TODO: ctx
	if err != nil {
		return errors.New(`Mongo error`) // TODO
	}

	return nil
}
