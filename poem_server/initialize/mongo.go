package initialize

import (
	"context"
	"github.com/qiniu/qmgo"
	"log"
	"poem_server/global"
)

// Mongo 初始化mongo
func Mongo() *qmgo.Database {
	m := global.POEM_CONFIG.Mongo
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{
		Uri:         m.Uri(),
		MaxPoolSize: &m.MaxPoolSize,
		MinPoolSize: &m.MinPoolSize,
	})
	if err != nil {
		log.Println(err)
	}
	db := client.Database(m.Dbname)
	return db
}
