package model

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Usuario struct {
	Codigo            int    `json:"codigo" bson:"codigo"`
	Nome              string `json:"nome" bson:"nome"`
	Login             string `json:"login" bson:"login"`
	CodPersonificador int    `json:"cod_personificador" bson:"cod_personificador"`
}

type Access struct {
	Id             primitive.ObjectID `json:"id" bson:"_id"`
	App            string             `json:"app" bson:"app"`
	Usuario        Usuario            `json:"usuario" bson:"usuario"`
	RemoteAddr     string             `json:"remote_addr" bson:"remote_addr"`
	UserAgent      string             `json:"user_agent" bson:"user_agent"`
	RequestMethod  string             `json:"request_method" bson:"request_method"`
	Host           string             `json:"host" bson:"host"`
	ScriptName     string             `json:"script_name" bson:"script_name"`
	QueryString    string             `json:"query_string" bson:"query_string"`
	ScriptFilename string             `json:"script_filename" bson:"script_filename"`
	SessionId      string             `json:"session_id" bson:"session_id"`
	DataHora       time.Time          `json:"data_hora" bson:"data_hora"`
	Tempo          float64            `json:"tempo" bson:"tempo"`
	Memoria        int                `json:"memoria" bson:"memoria"`
}

func List(app string, ini time.Time, fim time.Time) ([]Access, error) {
	d := time.Duration(fim.UnixNano() - ini.UnixNano())
	if d > time.Hour*24 {
		return nil, fmt.Errorf("list: ini e fim não podem ter diferença de mais de 1 dia")
	}
	fmt.Println(d)

	fmt.Println()
	if app == "" {
		return nil, fmt.Errorf("list: app obrigatório`")
	}
	filter := bson.D{
		{"app", app},
		{"data_hora", bson.D{
			{"$gte", ini},
			{"$lte", fim},
		}},
	}
	cur, err := db.Collection("access").Find(nil, filter)
	if err != nil {
		return nil, fmt.Errorf("list: %w", err)
	}
	lista := make([]Access, 0)
	for cur.Next(nil) {
		var a Access
		if err := cur.Decode(&a); err != nil {
			return nil, fmt.Errorf("list: %w", err)
		}
		a.DataHora = a.DataHora.Local()
		lista = append(lista, a)
	}
	return lista, nil
}
