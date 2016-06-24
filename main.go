// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	pnm "github.com/kkdai/pnm"
)

var mess = &Messenger{}
var petsDB *pnm.Pets

func main() {
	petsDB = pnm.NewPets()
	port := os.Getenv("PORT")
	log.Println("Server start in port:", port)
	mess.VerifyToken = os.Getenv("TOKEN")
	mess.AccessToken = os.Getenv("TOKEN")
	log.Println("Bot start in token:", mess.VerifyToken)
	mess.MessageReceived = MessageReceived
	http.HandleFunc("/webhook", mess.Handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

//MessageReceived :Callback to handle when message received.
func MessageReceived(event Event, opts MessageOpts, msg ReceivedMessage) {
	// log.Println("event:", event, " opt:", opts, " msg:", msg)
	profile, err := mess.GetProfile(opts.Sender.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	petObj := petsDB.GetNextPet()
	out := fmt.Sprintf("您好，目前的動物：名為%s, 所在地為:%s, 敘述: %s 電話為:%s", petObj.Name, petObj.Resettlement, petObj.Note, petObj.Phone)

	resp, err := mess.SendSimpleMessage(opts.Sender.ID, fmt.Sprintf("%s %s 你好，這裡有寵物等著你領樣: %s", profile.FirstName, profile.LastName, out))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", resp)
}
