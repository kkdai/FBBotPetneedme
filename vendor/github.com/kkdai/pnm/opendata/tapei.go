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

package opendata

import (
	"encoding/json"
	"log"
)

//OpenDataURL :

const (
	TaipeiURL string = "http://data.taipei/opendata/datalist/apiAccess?scope=resourceAquire&rid=f4a75ba9-7721-4363-884d-c3820b0b917c"
)

//TaipeiPets :Get from  http://data.taipei/opendata/datalist/apiAccess?scope=resourceAquire&rid=f4a75ba9-7721-4363-884d-c3820b0b917c
type TaipeiPets struct {
	Result struct {
		Offset  int         `json:"offset"`
		Limit   int         `json:"limit"`
		Count   int         `json:"count"`
		Sort    string      `json:"sort"`
		Results []TaipeiPet `json:"results"`
	} `json:"result"`
}

//TaipeiPet :
type TaipeiPet struct {
	ID              string `json:"_id"`
	Name            string `json:"Name"`
	Sex             string `json:"Sex"`
	Type            string `json:"Type"`
	Build           string `json:"Build"`
	Age             string `json:"Age"`
	Variety         string `json:"Variety"`
	Reason          string `json:"Reason"`
	AcceptNum       string `json:"AcceptNum"`
	ChipNum         string `json:"ChipNum"`
	IsSterilization string `json:"IsSterilization"`
	HairType        string `json:"HairType"`
	Note            string `json:"Note"`
	Resettlement    string `json:"Resettlement"`
	Phone           string `json:"Phone"`
	Email           string `json:"Email"`
	ChildreAnlong   string `json:"ChildreAnlong"`
	AnimalAnlong    string `json:"AnimalAnlong"`
	Bodyweight      string `json:"Bodyweight"`
	ImageName       string `json:"ImageName"`
}

//GetURL :
func (t *TaipeiPets) GetURL() string {
	return TaipeiURL
}

//GetPets :
func (t *TaipeiPets) GetPets() []Pet {
	c := NewClient(t.GetURL())
	body, err := c.GetHttpRes()
	if err != nil {
		return nil
	}

	// log.Println("ret:", string(body))
	var result TaipeiPets
	err = json.Unmarshal(body, &result)

	if err != nil {
		//error
		log.Fatal(err)
	}
	// log.Println("First Data:", result.Result.Results[0])

	var pets []Pet
	for _, v := range result.Result.Results {
		var pet Pet
		pet.Name = v.Name
		pet.Resettlement = v.Resettlement
		pet.Note = v.Note
		pet.Phone = v.Phone
		pet.ImageName = v.ImageName
		pet.Type = v.Type
		pets = append(pets, pet)
	}

	return pets
}
