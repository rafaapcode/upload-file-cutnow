package database_pkg

import (
	"context"
	"fmt"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client *mongo.Client
	HexId  string
}

func Connect() *mongo.Client {
	opts := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic("Erro ao se conectar ao banco")
	}

	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client
}

func (db Database) handleAndUpdateData(pathToImage, collectioName, fieldToUpdate string) (*mongo.UpdateResult, error) {
	bannerUrl := fmt.Sprintf("https://deb5gzd2n1wd.cloudfront.net/%s", pathToImage)
	coll := db.Client.Database("cutnow").Collection(collectioName)
	id, err := primitive.ObjectIDFromHex(db.HexId)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{fieldToUpdate, bannerUrl}}}}

	result, err := coll.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (db Database) Disconnect() {
	if err := db.Client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func (db Database) UpdateBarberBanner(pathToBanner string) (*mongo.UpdateResult, error) {
	// bannerUrl := fmt.Sprintf("https://deb5gzd2n1wd.cloudfront.net/%s", pathToBanner)
	// coll := db.Client.Database("cutnow").Collection("Barbeiro")
	// id, err := primitive.ObjectIDFromHex(db.HexId)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return nil, err
	// }

	// filter := bson.D{{"_id", id}}
	// update := bson.D{{"$set", bson.D{{"informacoes.banner", bannerUrl}}}}

	// result, err := coll.UpdateOne(context.TODO(), filter, update)

	result, err := db.handleAndUpdateData(pathToBanner, "Barbeiro", "informacoes.banner")

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return result, nil
}

func (db Database) UpdateBarbershopBanner(pathToBanner string) (*mongo.UpdateResult, error) {
	// bannerUrl := fmt.Sprintf("https://deb5gzd2n1wd.cloudfront.net/%s", pathToBanner)

	// coll := db.Client.Database("cutnow").Collection("Barbearia")
	// id, err := primitive.ObjectIDFromHex(db.HexId)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return nil, err
	// }

	// filter := bson.D{{"_id", id}}
	// update := bson.D{{"$set", bson.D{{"informacoes.fotoBanner", bannerUrl}}}}
	// result, err := coll.UpdateOne(context.TODO(), filter, update)

	result, err := db.handleAndUpdateData(pathToBanner, "Barbearia", "informacoes.fotoBanner")

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return result, nil
}

func (db Database) UpdateBarberFoto(pathToFoto string) (*mongo.UpdateResult, error) {
	// fotoUrl := fmt.Sprintf("https://deb5gzd2n1wd.cloudfront.net/%s", pathToFoto)
	// coll := db.Client.Database("cutnow").Collection("Barbeiro")

	// id, err := primitive.ObjectIDFromHex(db.HexId)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return nil, err
	// }

	// filter := bson.D{{"_id", id}}
	// update := bson.D{{"$set", bson.D{{"informacoes.foto", fotoUrl}}}}

	// result, err := coll.UpdateOne(context.TODO(), filter, update)

	result, err := db.handleAndUpdateData(pathToFoto, "Barbeiro", "informacoes.foto")

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return result, nil
}

func (db Database) UpdateBarbershopLogo(pathToLogo string) (*mongo.UpdateResult, error) {
	// logoUrl := fmt.Sprintf("https://deb5gzd2n1wd.cloudfront.net/%s", pathToLogo)
	// coll := db.Client.Database("cutnow").Collection("Barbearia")
	// id, err := primitive.ObjectIDFromHex(db.HexId)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return nil, err
	// }

	// filter := bson.D{{"_id", id}}
	// update := bson.D{{"$set", bson.D{{"informacoes.logo", logoUrl}}}}

	// result, err := coll.UpdateOne(context.TODO(), filter, update)

	result, err := db.handleAndUpdateData(pathToLogo, "Barbearia", "informacoes.logo")

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return result, nil
}

func (db Database) UpdateBarberPotfolio(pathToPortfolio []string) (*mongo.UpdateResult, error) {
	var portfolioUrls []string

	for _, pathPotfolio := range pathToPortfolio {
		portfolioUrls = append(portfolioUrls, fmt.Sprintf("https://deb5gzd2n1wd.cloudfront.net/%s", pathPotfolio))
	}

	coll := db.Client.Database("cutnow").Collection("Barbeiro")
	id, err := primitive.ObjectIDFromHex(db.HexId)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"informacoes.portfolio", portfolioUrls}}}}

	result, err := coll.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return result, nil
}

func (db Database) UpdateBarbershopStructure(pathToStructure []string) (*mongo.UpdateResult, error) {
	var structureUrls []string

	for _, pathStrucure := range pathToStructure {
		structureUrls = append(structureUrls, fmt.Sprintf("https://deb5gzd2n1wd.cloudfront.net/%s", pathStrucure))
	}

	coll := db.Client.Database("cutnow").Collection("Barbearia")
	id, err := primitive.ObjectIDFromHex(db.HexId)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"informacoes.fotosEstruturaBarbearia", structureUrls}}}}

	result, err := coll.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return result, nil
}

func (db Database) DeleteStructureImages(index int) (string, error) {
	coll := db.Client.Database("cutnow").Collection("Barbearia")
	id, err := primitive.ObjectIDFromHex(db.HexId)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	filter := bson.D{{"_id", id}}
	var results bson.M
	err = coll.FindOne(context.TODO(), filter).Decode(&results)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	informacoes, ok := results["informacoes"].(bson.M)
	fotosEstruturaBarbearia, ok := informacoes["fotosEstruturaBarbearia"].(bson.A)

	if !ok {
		return "", fmt.Errorf("Erro ao acessar as imagens")
	}

	if index > len(fotosEstruturaBarbearia) {
		return "", fmt.Errorf("Indice não existe")
	}

	var newImages []string
	var deletedImages string

	for key, val := range fotosEstruturaBarbearia {
		urlImg := val.(string)
		if key != index {
			newImages = append(newImages, urlImg)
		} else {
			pathToS3Object := strings.SplitAfterN(urlImg, "/", 4)
			deletedImages = pathToS3Object[3]
		}
	}
	update := bson.D{{"$set", bson.D{{"informacoes.fotosEstruturaBarbearia", newImages}}}}

	_, err = coll.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return deletedImages, nil
}
