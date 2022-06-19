package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pwhb/rf_telebot/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Student struct {
	Id               primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	TelegramId       int64              `bson:"telegram_id" json:"telegram_id,omitempty"`
	FirstName        string             `bson:"first_name" json:"first_name,omitempty"`
	LastName         string             `bson:"last_name" json:"last_name,omitempty"`
	TelegramUsername string             `bson:"telegram_username" json:"telegram_username,omitempty"`
	GithubUsername   string             `bson:"github_username" json:"github_username,omitempty"`
	Email            string             `bson:"email" json:"email,omitempty"`
	StudentId        int8               `bson:"student_id" json:"student_id,omitempty"`
	ReasonForJoining string             `bson:"reason_for_joining" json:"reason_for_joining,omitempty"`
	ClassName        string             `bson:"class_name" json:"class_name,omitempty"`
	BatchName        string             `bson:"batch_name" json:"batch_name,omitempty"`
	CreatedAt        time.Time          `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt        time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
}

var studentCol *mongo.Collection

var ctx = context.TODO()

func init() {
	c, err := config.GetMongoDBClient()
	if err != nil {
		log.Panicln(err)
	}
	db := c.Database("test")
	studentCol = db.Collection("b17_students")
	log.Println("database connected")
}

func GetAllStudents() ([]*Student, error) {
	var students []*Student
	cur, err := studentCol.Find(ctx, bson.D{})
	if err != nil {
		return students, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var s Student
		err := cur.Decode(&s)
		if err != nil {
			return students, err
		}
		students = append(students, &s)
	}

	return students, nil
}

func CreateNewStudent(s *Student) {
	s.Id = primitive.NewObjectID()
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
	num, err := studentCol.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Println(err)
	}
	s.StudentId = int8(num)
	result, err := studentCol.InsertOne(ctx, s)
	log.Printf("created: %v\n", result.InsertedID)
	if err != nil {
		log.Panicln(err)
	}
}

func GetOneStudent(telegram_id int64) (*Student, error) {
	var s *Student
	filter := bson.M{"telegram_id": telegram_id}
	err := studentCol.FindOne(ctx, filter).Decode(&s)
	return s, err

}

func GetInfoString(s *Student) string {
	return fmt.Sprintf("သင့် information\nFirst Name: %v\nLast Name: %v\nStudent ID: %v\nSecret ID: %v\nEmail: %v\nGithub Username: %v\n",
		s.FirstName,
		s.LastName,
		s.StudentId,
		s.Id.Hex(),
		s.Email,
		s.GithubUsername)
}
