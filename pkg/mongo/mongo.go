package database_pkg

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

func Disconnect(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func UpdateBarberBanner(client *mongo.Client, id string, pathToBanner string) *mongo.UpdateResult {
	bannerUrl := fmt.Sprintf("https://deb5gzd2n1wd.cloudfront.net/%s", pathToBanner)

	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"informacoes.banner", bannerUrl}}}}

	result, err := client.Database("cutnow").Collection("Barbeiro").UpdateByID(context.TODO(), filter, update)

	if err != nil {
		fmt.Println(err.Error())
		panic("Erro ao atualizar o banner no Banco de dados")
	}

	return result
}

func UpdateBarbershopBanner(client *mongo.Client, hexId string, pathToBanner string) (*mongo.UpdateResult, error) {
	bannerUrl := fmt.Sprintf("https://deb5gzd2n1wd.cloudfront.net/%s", pathToBanner)

	coll := client.Database("cutnow").Collection("Barbearia")
	id, err := primitive.ObjectIDFromHex(hexId)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"informacoes.fotoBanner", bannerUrl}}}}
	result, err := coll.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return result, nil
}

func UpdateBarberFoto(client *mongo.Client, id string, pathToFoto string) *mongo.UpdateResult {
	fotoUrl := fmt.Sprintf("https://deb5gzd2n1wd.cloudfront.net/%s", pathToFoto)

	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"informacoes.foto", fotoUrl}}}}

	result, err := client.Database("cutnow").Collection("Barbeiro").UpdateByID(context.TODO(), filter, update)

	if err != nil {
		fmt.Println(err.Error())
		panic("Erro ao atualizar o banner no Banco de dados")
	}

	return result
}

func UpdateBarbershopLogo(client *mongo.Client, id string, pathToLogo string) *mongo.UpdateResult {
	logoUrl := fmt.Sprintf("https://deb5gzd2n1wd.cloudfront.net/%s", pathToLogo)

	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"informacoes.logo", logoUrl}}}}

	result, err := client.Database("cutnow").Collection("Barbearia").UpdateByID(context.TODO(), filter, update)

	if err != nil {
		fmt.Println(err.Error())
		panic("Erro ao atualizar o banner no Banco de dados")
	}

	return result
}

func UpdateBarberPotfolio(client *mongo.Client, id string, pathToPortfolio []string) *mongo.UpdateResult {
	var portfolioUrls []string

	for _, pathPotfolio := range pathToPortfolio {
		portfolioUrls = append(portfolioUrls, fmt.Sprintf("https://deb5gzd2n1wd.cloudfront.net/%s", pathPotfolio))
	}

	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"informacoes.portfolio", portfolioUrls}}}}

	result, err := client.Database("cutnow").Collection("Barbeiro").UpdateByID(context.TODO(), filter, update)

	if err != nil {
		fmt.Println(err.Error())
		panic("Erro ao atualizar o banner no Banco de dados")
	}

	return result
}

func UpdateBarbershopStructure(client *mongo.Client, id string, pathToStructure []string) *mongo.UpdateResult {
	var structureUrls []string

	for _, pathStrucure := range pathToStructure {
		structureUrls = append(structureUrls, fmt.Sprintf("https://deb5gzd2n1wd.cloudfront.net/%s", pathStrucure))
	}

	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"informacoes.fotosEstruturaBarbearia", structureUrls}}}}

	result, err := client.Database("cutnow").Collection("Barbearia").UpdateByID(context.TODO(), filter, update)

	if err != nil {
		fmt.Println(err.Error())
		panic("Erro ao atualizar o banner no Banco de dados")
	}

	return result
}
