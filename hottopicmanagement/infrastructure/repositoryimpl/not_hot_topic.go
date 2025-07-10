package repositoryimpl

import (
	"go.mongodb.org/mongo-driver/bson"

	commonRepo "github.com/opensourceways/hot-topic-website-backend/common/domain/repository"
	"github.com/opensourceways/hot-topic-website-backend/hottopicmanagement/domain"
)

func NewNotHotTopic(dao map[string]Dao) *notHotTopic {
	return &notHotTopic{
		daoMap: dao,
	}
}

type notHotTopic struct {
	daoMap
}

func (impl *notHotTopic) add(community string, v *domain.NotHotTopic) error {
	do := tonotNotHotTopicDO(v)
	doc, err := do.toDoc()
	if err != nil {
		return err
	}

	docFilter := bson.M{fieldTitle: v.Title}

	dao, err := impl.dao(community)
	if err != nil {
		return err
	}

	_, err = dao.InsertDocIfNotExists(docFilter, doc)
	if err != nil && dao.IsDocExists(err) {
		err = commonRepo.NewErrorDuplicateCreating(err)
	}

	return err
}

func (impl *notHotTopic) Save(community string, items []domain.NotHotTopic) error {
	dao, err := impl.dao(community)
	if err != nil {
		return err
	}

	if err := dao.DeleteDocs(bson.M{}); err != nil {
		return err
	}

	for i := range items {
		item := &items[i]

		do := tonotNotHotTopicDO(item)
		doc, err := do.toDoc()
		if err != nil {
			return err
		}

		if _, err = dao.InsertDoc(doc); err != nil {
			return err
		}
	}

	return nil
}

func (impl *notHotTopic) FindAll(community string) ([]domain.NotHotTopic, error) {
	dao, err := impl.dao(community)
	if err != nil {
		return nil, err
	}

	var dos []notNotHotTopicDO

	if err := dao.GetDocs(nil, nil, nil, &dos); err != nil {
		return nil, err
	}

	v := make([]domain.NotHotTopic, len(dos))
	for i := range dos {
		v[i] = dos[i].toNotHotTopic()
	}

	return v, nil
}
