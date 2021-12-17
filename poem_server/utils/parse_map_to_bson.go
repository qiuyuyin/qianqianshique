package utils

import "go.mongodb.org/mongo-driver/bson"

func ParseMapToBson(m map[string]interface{}) (error, bson.M) {
	res := bson.M{}
	marshal, err := bson.Marshal(m)
	if err != nil {
		return err, nil
	}
	err = bson.Unmarshal(marshal, res)
	if err != nil {
		return err, nil
	}
	return err, res
}
