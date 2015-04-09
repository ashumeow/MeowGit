package commits

import (

	"fmt"
	"MeowGit/mgo"
	"MeowGit/mgo/bson"
)

type Commit struct {

	Title string
	Author string
	Language string
	VersionID string
	Id string
}

// Creates an account and adds it to the Database
func CreateCommit(commit *Commit) bool {

	// Dial up a mongoDB session
	session, err := mgo.Dial("127.0.0.1:27017/")
	if err != nil {
		fmt.Println(err)
		return false
	}

	// Opens the "repository" databases, "commits" collection
	c := session.DB("repository").C("commits")
	result := Commit{}
	
	// Search for the CommitID, place in result.Id
	err = c.Find(bson.M{"id": commit.Id}).One(&result)
	
	if result.Id != "" {
		// return true because commit is present in the repository DB
		// and we can say, "it's been added" without causing errors
		return true
	}

	// Add commit if it is not already in the repository DB
	err = c.Insert(*commit)

	if err != nil {
		return false
	}
	
	session.Close()
	
	return true
}

// Finds commit already in database
func FindCommit(commitid string) (commit *Commit) {

	session, err := mgo.Dial("127.0.0.1:27017/")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	
	// Opens the "repository" databases, "commits" collection
	c := session.DB("repository").C("commits")
	
	// Finds the commitid and fills the commit struct with the data
	err = c.Find(bson.M{"id": commitid}).One(&commit)
	
	session.Close()

	return commit
}
