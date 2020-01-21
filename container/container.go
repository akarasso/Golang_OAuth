package container

import "go.mongodb.org/mongo-driver/mongo"

type dependencies struct {
	AccessToken       *mongo.Collection `db:"oauth" collection:"access_token"`
	AuthorizationCode *mongo.Collection `db:"oauth" collection:"authorization_code"`
	Client            *mongo.Collection `db:"oauth" collection:"client"`
	RefreshToken      *mongo.Collection `db:"oauth" collection:"refresh_token"`
}

// Dependencies :: Container services dependencies
var Dependencies dependencies
