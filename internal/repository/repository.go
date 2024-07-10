package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/user/article-construct-demo/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrUserNotFound = errors.New("construct not found")
)

type ConstructRepository interface {
	GetConstructForIan(ian string, ctx context.Context) (*model.Item, error)
	GetItemForIan(ian string, ctx context.Context) (*model.Item, error)
	GetCaseForIan(ian string, ctx context.Context) (*model.Case, error)
	GetVariantForIan(ian string, ctx context.Context) (*model.Variant, error)
}

type repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) ConstructRepository {
	return &repository{db: db}
}

func (r repository) GetItemForIan(ian string, ctx context.Context) (*model.Item, error) {
	var item *model.Item
	errs := r.db.Collection("item").FindOne(ctx, bson.D{bson.E{Key: "ian", Value: ian}}).Decode(&item)
	return item, errs
}

func (r repository) GetCaseForIan(ian string, ctx context.Context) (*model.Case, error) {
	var sk *model.Case
	errs := r.db.Collection("case").FindOne(ctx, bson.D{bson.E{Key: "ian", Value: ian}}).Decode(&sk)
	return sk, errs
}

func (r repository) GetVariantForIan(ian string, ctx context.Context) (*model.Variant, error) {
	var ea *model.Variant
	errs := r.db.Collection("case").FindOne(ctx, bson.D{bson.E{Key: "ian", Value: ian}}).Decode(&ea)
	return ea, errs
}

func (r repository) GetConstructForIan(ian string, ctx context.Context) (*model.Item, error) {

	matchStage := bson.D{
		bson.E{Key: "$match", Value: bson.D{
			bson.E{Key: "ian", Value: ian},
		}},
	}

	lookupStage := bson.D{
		bson.E{Key: "$lookup", Value: bson.D{
			bson.E{Key: "from", Value: "case"},
			bson.E{Key: "localField", Value: "itemID"},
			bson.E{Key: "foreignField", Value: "UniqueCaseIDs"},
			bson.E{Key: "as", Value: "cases"},
			bson.E{Key: "pipeline", Value: bson.A{
				bson.D{
					bson.E{Key: "$lookup", Value: bson.D{
						bson.E{Key: "from", Value: "variant"},
						bson.E{Key: "localField", Value: "uniqueVariantIds"},
						bson.E{Key: "foreignField", Value: "uniqueId"},
						bson.E{Key: "as", Value: "variants"},
						bson.E{Key: "pipeline", Value: bson.A{
							bson.D{
								bson.E{Key: "$project", Value: bson.D{
									bson.E{Key: "_id", Value: 0},
									bson.E{Key: "ian", Value: 1},
									bson.E{Key: "nat", Value: 1},
									bson.E{Key: "status", Value: 1},
								}},
							},
						}},
					}},
				},
				bson.D{
					bson.E{Key: "$project", Value: bson.D{
						bson.E{Key: "_id", Value: 0},
						bson.E{Key: "ian", Value: 1},
						bson.E{Key: "nat", Value: 1},
						bson.E{Key: "status", Value: 1},
						bson.E{Key: "variants", Value: 1},
					}},
				},
			}},
		}},
	}

	projectItemStage := bson.D{
		bson.E{Key: "$project", Value: bson.D{
			bson.E{Key: "_id", Value: 0},
			bson.E{Key: "ian", Value: 1},
			bson.E{Key: "nat", Value: 1},
			bson.E{Key: "status", Value: 1},
			bson.E{Key: "cases", Value: 1},
		}},
	}
	pipeline := mongo.Pipeline{matchStage, lookupStage, projectItemStage}

	aggregate, err := r.db.Collection("item").Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("Error occurred with query", pipeline, err)
	}

	var result []model.Item
	if err = aggregate.All(ctx, &result); err != nil {
		return nil, fmt.Errorf("Message deconding error", err)
	}

	if len(result) == 0 {
		return nil, ErrUserNotFound
	}

	cas := result[0]

	return &cas, nil
}
