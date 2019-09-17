package main

import (
	"fmt"
	"regexp"
)

func normalize(s string) string {
	re := regexp.MustCompile("[^0-9]")
	ans := re.ReplaceAllString(s, "")
	return ans
}

func main() {
	fmt.Println("Phone normalizer")
	db := database{nil, "phone", "root", "k9eUnehq#KD72"}
	//defer db.closeConnection(db.db)
	err := db.resetDatabase()
	hanldeError(err)

	db.db, err = db.connectToDatabase(db.dbName)
	hanldeError(err)
	err = db.createPhoneTable()
	_, err = db.insertPhone("123 456 7891")
	hanldeError(err)

	_, err = db.insertPhone("1234567890")
	hanldeError(err)
	_, err = db.insertPhone("123 456 7891")
	hanldeError(err)
	_, err = db.insertPhone("(123) 456 7892")
	hanldeError(err)
	_, err = db.insertPhone("(123) 456-7893")
	hanldeError(err)
	_, err = db.insertPhone("123-456-7894")
	hanldeError(err)
	_, err = db.insertPhone("123-456-7890")
	hanldeError(err)
	_, err = db.insertPhone("1234567892")
	hanldeError(err)
	_, err = db.insertPhone("(123)456-7892")
	hanldeError(err)

	number, err := db.getPhone(1)
	fmt.Printf("phone is %s\n", number)

	phoneList, err := db.getAllPhones()
	for _, p := range phoneList {
		number := normalize(p.number)
		if number != p.number {
			fmt.Println("Updating or deleting...")
			existing, err := db.findPhone(number)
			hanldeError(err)
			if existing == nil {
				p.number = number
				hanldeError(db.updatePhone(p))
			} else {
				hanldeError(db.deletePhone(p))
			}
		}
	}
	//fmt.Println(phoneList)
	defer db.closeConnection()
}
