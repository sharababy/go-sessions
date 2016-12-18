package customSessions

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"log"
)


type Session struct{

	Username string `bson:"Username" json:"Username" `
	UniqueIP string `bson:"UniqueIP" json:"UniqueIP" `
}


func GetSession( DB_Url string, DB_Name string , Collection_Name string ,UniqueIP string) ([]Session,error){

	var lookfor []Session

	session := CreateSession(DB_Url)

	f := session.DB(DB_Name).C(Collection_Name)

	FindWith := bson.M{"UniqueIP" : UniqueIP}

    err := f.Find(FindWith).All(&lookfor)

    session.Close()

    if(err!=nil){
        	
        return lookfor , err 
    } else{
        return lookfor , nil
    }


}


func PutSession(DB_Url string, DB_Name string , Collection_Name string ,This Session) error{

	session := CreateSession(DB_Url)

	c :=session.DB(DB_Name).C(Collection_Name)

	err := c.Insert(This)

	if(err!=nil){
		log.Println(err)
		return err
	}
	//fmt.Println("Inserted doc..")
	session.Close()

	return nil

}


func DeleteSession(DB_Url string, DB_Name string, Collection_Name string,find_type string ,find_with string) (error) {

		session := CreateSession(DB_Url)

       	collection := session.DB(DB_Name).C(Collection_Name)

		Delete_this := bson.M{ find_type : find_with }

        err := collection.Remove(Delete_this)

        session.Close()

        if(err!=nil){
        	return err 
        } else{
        	return nil
        }
}


func CreateSession(DB_Url string) *mgo.Session{

	session, err := mgo.Dial(DB_Url)

	if(err!=nil){
		log.Fatal(err)
	}

	session.SetSafe(&mgo.Safe{})

	return session
}