package dbclient 

import(
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
	"os"
)

var(
	bc IBoltClient = &BoltClient{}
)

func Test(t *testing.T){
	os.Remove("bolt.db")
}
func TestQueueInsertAndDelete(t *testing.T){
	var err error
	bc.OpenBoltDb()
	err=bc.QueueInsert("abc")
	assert.Nil(t,err)
	err=bc.QueueDelete("abc")
	assert.Nil(t,err)
	err=bc.QueueInsert("abc")
	assert.Nil(t,err)
	err=bc.QueueInsert("abc")
	assert.NotNil(t,err)
	err=bc.QueueDelete("abc")
	assert.Nil(t,err)
}

func TestQueueSelect(t *testing.T){
	var err error
	err=bc.QueueInsert("abc")
	assert.Nil(t,err)
	err=bc.QueueInsert("bdd")
	assert.Nil(t,err)
	fmt.Println(bc.QueueSelect())
}
