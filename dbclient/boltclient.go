package dbclient

import(
	"log"
	"github.com/boltdb/bolt"
	"github.com/satori/go.uuid"
)

type IBoltClient interface{
	OpenBoltDb()
	QueueSelect()[]string																						// select list of queueId
	QueueInsert(queueId string)error																// insert queueId
	QueueDelete(queueId string)error																// delete queueId
	QueueSelectId(queueId string)[]string														// select list of id from queue queueId	
	SelectRecord(queueId,id string)[]byte														// select value from queueId with key = id
	InsertRecord(queueId string,data []byte)(string,error)					// insert into queueId with value = data return key as string
	DeleteRecord(queueId,id string)error														// delete from queueId with key = id return key as string	
}

type BoltClient struct{
	boltDB	*bolt.DB
}

func(bc *BoltClient)DeleteRecord(queueId,id string)error{
	err:=bc.boltDB.Update(func(tx *bolt.Tx)error{
		b:=tx.Bucket([]byte(queueId))
		err:=b.Delete([]byte(id))
		return err
	})
	return err
}

func(bc *BoltClient)InsertRecord(queueId string,data []byte)(string,error){
	u1,err:=uuid.NewV4()
	if err!=nil{
		return "",err 
	}
	err=bc.boltDB.Update(func(tx *bolt.Tx)error{
		b:=tx.Bucket([]byte(queueId))
	  err:=b.Put(u1.Bytes(),data)
		return err
	})
	return u1.String(),err
}

func(bc *BoltClient)SelectRecord(queueId,id string)[]byte{
	var result []byte
	bc.boltDB.View(func(tx *bolt.Tx)error{
		b:=tx.Bucket([]byte(queueId))
		result=b.Get([]byte(id))
		return nil
	})
	return result
}

func(bc *BoltClient)QueueSelectId(queueId,id string)[]string{
	result:=[]string{}
	bc.boltDB.View(func(tx *bolt.Tx)error{
		b:=tx.Bucket([]byte(queueId))
		c:=b.Cursor()
		for k,_:=c.First();k!=nil;k,_=c.Next(){
			result=append(result,string(k))
		}
		return nil
	})
	return result
}

func(bc *BoltClient)QueueDelete(queueId string)error{
	err:=bc.boltDB.Update(func(tx *bolt.Tx)error{
		return tx.DeleteBucket([]byte(queueId))
	})
	return err
}

func(bc *BoltClient)QueueInsert(queueId string)error{
	err:=bc.boltDB.Update(func(tx *bolt.Tx)error{
		_,err:=tx.CreateBucket([]byte(queueId))
		return err
	})
	return err
}

func(bc *BoltClient)QueueSelect()[]string{
	result:=[]string{}
	bc.boltDB.View(func(tx *bolt.Tx)error{
		tx.ForEach(func(name []byte,b *bolt.Bucket)error{
			result=append(result,string(name))
			return nil
		})
		return nil 
	})
	return result
}

func(bc *BoltClient)OpenBoltDb(){
	var err error
	bc.boltDB,err=bolt.Open("bolt.db",0600,nil)
	if err!=nil{
		log.Fatal(err)
	}
}

