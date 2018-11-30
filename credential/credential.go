package main

import(
	"time"
	"encoding/json"
	"crypto/sha256"
	"taurus.com/proxy/dbclient"
	"os"
	"log"
)

var DBClient dbclient.IBoltClient

func setupSecret()error{
	list:=DBClient.QueueSelect()
	exists:=false
	for _,v:=range list{
		if v==dbclient.SecretKey{
			exists=true
		}
	}
	if !exists{
		log.Printf("secret bucket does not exists, create it")
		err:=DBClient.QueueInsert(dbclient.SecretKey)
		if err!=nil{
			return err
		}
	}else{
		log.Printf("secret bucket exists")
	}
	return nil
}

func main(){

	DBClient=&dbclient.BoltClient{}
	DBClient.OpenBoltDb()
	err:=setupSecret()
	if err!=nil{
		log.Fatal(err)
	}
	
	args:=os.Args[1:]

	switch len(args){
	case 2:
		username:=args[0]
		log.Print(username)
		b:=DBClient.SelectRecord(dbclient.SecretKey, username)

		if len(b)!=0{
			log.Fatal("username exists")
		}
		h:=sha256.New()
		h.Write([]byte(args[1]))
		credential:=dbclient.Credential{Password:h.Sum(nil),Timestamp:time.Now().Unix()}
		jcred,err:=json.Marshal(credential)
		if err!=nil{
			log.Fatal("json error: %v",err)
		}
		
		err=DBClient.InsertCredential(dbclient.SecretKey,args[0],jcred)
		if err!=nil{
			log.Fatal(err)
		}
	case 1:
		err:=DBClient.DeleteRecord(dbclient.SecretKey,args[0])
		if err!=nil{
			log.Fatal(err)
		}
		log.Printf("credentials %s deleted\n",args[0])
	case 0:
		b:=DBClient.QueueSelectId(dbclient.SecretKey)
		log.Printf("%d credentials found\n",len(b))
		for i,v:=range b{
			a:=DBClient.SelectRecord(dbclient.SecretKey,v)
			log.Printf("%d %s:%s\n",i,v,string(a))
		}
	default:
		log.Fatal("Invalid number of arguments")
	}

}
