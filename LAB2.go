collection := client.Database("mydb").Collection("mycollection")
filter := bson.D{{"name", "Dhoni"}}
var result bson.M
err := collection.FindOne(context.Background(), filter).Decode(&result)
if err != nil {
	log.Fatal(err)
}
fmt.Println(result)