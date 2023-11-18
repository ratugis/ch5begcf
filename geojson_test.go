package gcfbackendfa

import (
	"context"
	"fmt"
	"testing"

	pasproj "github.com/HRMonitorr/PasetoprojectBackend"
	"github.com/whatsauth/watoken"
)

//func TestGCHandlerFunc(t *testing.T) {
//	data := GCHandlerFunc("string", "GIS", "geogis")
//
//	fmt.Printf("%+v", data)
//}

func TestGeneratePaseto(t *testing.T) {
	privateKey, publicKey := watoken.GenerateKey()
	fmt.Println(privateKey)
	fmt.Println(publicKey)
	hasil, err := watoken.Encode("gis5", "PrivateKey")
	fmt.Println(hasil, err)
}

func TestUpdateData(t *testing.T) {
	data := LonLatProperties{
		Type:   "Polygon",
		Name:   "Martabak Manis",
		Volume: "1",
	}
	up := UpdateNameGeo("MONGOSTRING", "gis4", context.Background(), data)
	fmt.Println(up)
}

func TestDeleteDataGeo(t *testing.T) {
	data := LonLatProperties{
		Type:   "Polygon",
		Name:   "Martabak Manis",
		Volume: "1",
	}
	up := DeleteDataGeo("MONGOSTRING", "gis4", context.Background(), data)
	fmt.Println(up)
}

func TestInsertUser(t *testing.T) {
	conn := GetConnectionMongo("MONGOSTRING", "gis4")
	pass, _ := pasproj.HashPass("iyaslodons")
	data := RegisterStruct{
		Username: "iyas",
		Password: pass,
	}
	ins := InsertUserdata(conn, data.Username, data.Password)
	fmt.Println(ins)
}

func TestIsexist(t *testing.T) {
	data := IsExist("token", "PublicKey")
	fmt.Println(data)
}
