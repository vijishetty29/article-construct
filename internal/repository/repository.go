package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/user/article-construct-demo/internal/dto"
	"github.com/user/article-construct-demo/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrUserNotFound = errors.New("construct not found")
)

type ConstructRepository interface {
	GetConstructForIan(ian string, country string, ctx context.Context) (*model.Item, error)
	GetItemForIan(ian string, ctx context.Context) (*dto.ItemDto, error)
	GetCaseForIan(ian string, ctx context.Context) (*dto.CaseDto, error)
	GetVariantForIan(ian string, ctx context.Context) (*dto.VariantDto, error)
}

type repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) ConstructRepository {
	return &repository{db: db}
}

func (r repository) GetItemForIan(ian string, ctx context.Context) (*dto.ItemDto, error) {
	//var item *model.Item
	//errs := r.db.Collection("item").FindOne(ctx, bson.D{bson.E{Key: "ian", Value: ian}}).Decode(&item)

	matchStage := bson.D{
		bson.E{Key: "$match", Value: bson.D{
			bson.E{Key: "ian", Value: ian},
		}},
	}

	lookupStage := bson.D{
		bson.E{Key: "$lookup", Value: bson.D{
			bson.E{Key: "from", Value: "case"},
			bson.E{Key: "localField", Value: "uniqueCaseIds"},
			bson.E{Key: "foreignField", Value: "uniqueId"},
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
									bson.E{Key: "itemStatus", Value: 1},
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
						bson.E{Key: "itemStatus", Value: 1},
						bson.E{Key: "variants", Value: 1},
					}},
				},
			}},
		}},
	}

	pipeline := mongo.Pipeline{matchStage, lookupStage}
	aggregate, err := r.db.Collection("item").Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("Error occurred with query", pipeline, err)
	}

	var result []dto.ItemDto
	if err = aggregate.All(ctx, &result); err != nil {
		return nil, fmt.Errorf("Message deconding error", err)
	}

	if len(result) == 0 {
		return nil, ErrUserNotFound
	}

	item := result[0]

	return &item, nil
}

func (r repository) GetCaseForIan(ian string, ctx context.Context) (*dto.CaseDto, error) {
	matchStage := bson.D{
		bson.E{Key: "$match", Value: bson.D{
			bson.E{Key: "ian", Value: ian},
		}},
	}

	lookupStage := bson.D{
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
						bson.E{Key: "itemStatus", Value: 1},
					}},
				},
			}},
		}},
	}

	pipeline := mongo.Pipeline{matchStage, lookupStage}
	aggregate, err := r.db.Collection("case").Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("Error occurred with query", pipeline, err)
	}

	var result []dto.CaseDto
	if err = aggregate.All(ctx, &result); err != nil {
		return nil, fmt.Errorf("Message deconding error", err)
	}

	if len(result) == 0 {
		return nil, ErrUserNotFound
	}

	sk := result[0]

	return &sk, nil
}

func (r repository) GetVariantForIan(ian string, ctx context.Context) (*dto.VariantDto, error) {
	matchStage := bson.D{
		bson.E{Key: "$match", Value: bson.D{
			bson.E{Key: "ian", Value: ian},
		}},
	}

	lookupStage := bson.D{
		bson.E{Key: "$lookup", Value: bson.D{
			bson.E{Key: "from", Value: "case"},
			bson.E{Key: "localField", Value: "caseIds"},
			bson.E{Key: "foreignField", Value: "uniqueId"},
			bson.E{Key: "as", Value: "cases"},
			bson.E{Key: "pipeline", Value: bson.A{
				bson.D{
					bson.E{Key: "$project", Value: bson.D{
						bson.E{Key: "_id", Value: 0},
						bson.E{Key: "ian", Value: 1},
						bson.E{Key: "nat", Value: 1},
						bson.E{Key: "itemStatus", Value: 1},
					}},
				},
			}},
		}},
	}

	pipeline := mongo.Pipeline{matchStage, lookupStage}
	aggregate, err := r.db.Collection("variant").Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("Error occurred with query", pipeline, err)
	}

	var result []dto.VariantDto
	if err = aggregate.All(ctx, &result); err != nil {
		return nil, fmt.Errorf("Message deconding error", err)
	}

	if len(result) == 0 {
		return nil, ErrUserNotFound
	}

	ea := result[0]

	return &ea, nil
}

func (r repository) GetConstructForIan(ian string, country string, ctx context.Context) (*model.Item, error) {

	matchStage := bson.D{
		bson.E{Key: "$match", Value: bson.D{
			bson.E{Key: "ian", Value: ian},
			bson.E{Key: "country", Value: country},
		}},
	}

	lookupStage := bson.D{
		bson.E{Key: "$lookup", Value: bson.D{
			bson.E{Key: "from", Value: "case"},
			bson.E{Key: "localField", Value: "uniqueCaseIds"},
			bson.E{Key: "foreignField", Value: "uniqueId"},
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
									bson.E{Key: "itemStatus", Value: 1},
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
			bson.E{Key: "itemStatus", Value: 1},
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
