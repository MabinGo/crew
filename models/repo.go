package models

import (
	"fmt"
	"time"

	"github.com/containerops/wrench/db"
)

type Repository struct {
	UUID          string    `json:"UUID"`          //
	Repository    string    `json:"repository"`    //
	Namespace     string    `json:"namespace"`     //
	NamespaceType bool      `json:"namespacetype"` //
	Organization  string    `json:"organization"`  //
	Tags          []string  `json:"tags"`          //
	Starts        []string  `json:"starts"`        //
	Comments      []string  `json:"comments"`      //
	Short         string    `json:"short"`         //
	Description   string    `json:"description"`   //
	JSON          string    `json:"json"`          //
	Dockerfile    string    `json:"dockerfile"`    //
	Agent         string    `json:"agent"`         //
	Links         string    `json:"links"`         //
	Size          int64     `json:"size"`          //
	Download      int64     `json:"download"`      //
	Uploaded      bool      `json:"uploaded"`      //
	Checksum      string    `json:"checksum"`      //
	Checksumed    bool      `json:"checksumed"`    //
	Icon          string    `json:"icon"`          //
	Sign          string    `json:"sign"`          //
	Privated      bool      `json:"privated"`      //
	Clear         string    `json:"clear"`         //
	Cleared       bool      `json:"cleared"`       //
	Encrypted     bool      `json:"encrypted"`     //
	Created       int64     `json:"created"`       //
	Updated       int64     `json:"updated"`       //
	Memo          []string  `json:"memo"`          //
	Version       int64     `json:"version"`       //
	Privilege     Privilege `json:"privilege"`     //
}

type Star struct {
	UUID       string   `json:"UUID"`       //
	User       string   `json:"user"`       //
	Repository string   `json:"repository"` //
	Time       int64    `json:"time"`       //
	Memo       []string `json:"memo"`       //
}

type Comment struct {
	UUID       string   `json:"UUID"`       //
	Comment    string   `json:"comment"`    //
	User       string   `json:"user"`       //
	Repository string   `json:"repository"` //
	Time       int64    `json:"time"`       //
	Memo       []string `json:"memo"`       //
}

type Privilege struct {
	UUID       string   `json:"UUID"`       //
	Privilege  bool     `json:"privilege"`  //
	Team       string   `json:"team"`       //
	Repository string   `json:"repository"` //
	Memo       []string `json:"memo"`       //
}

type Tag struct {
	UUID       string   `json:"uuid"`       //
	Name       string   `json:"name"`       //
	ImageId    string   `json:"imageid"`    //
	Namespace  string   `json:"namespace"`  //
	Repository string   `json:"repository"` //
	Sign       string   `json:"sign"`       //
	Manifest   string   `json:"manifest"`   //
	Memo       []string `json:"memo"`       //
}

func (r *Repository) Has(namespace, repository string) (bool, string, error) {

	UUID, err := GetUUID("repository", fmt.Sprintf("%s:%s", namespace, repository))

	if err != nil {
		return false, "", err
	}

	if len(UUID) <= 0 {
		return false, "", nil
	}
	err = Get(r, UUID)

	return true, UUID, err
}

func (r *Repository) Save() error {
	if err := Save(r, r.UUID); err != nil {
		return err
	}

	if _, err := db.Client.HSet(db.GLOBAL_REPOSITORY_INDEX, (fmt.Sprintf("%s:%s", r.Namespace, r.Repository)), r.UUID).Result(); err != nil {
		return err
	}

	return nil
}

func (r *Repository) Put(namespace, repository, json, agent string, version int64) error {
	if has, _, err := r.Has(namespace, repository); err != nil {
		return err
	} else if has == false {
		r.UUID = db.GeneralDBKey(fmt.Sprintf("%s:%s", namespace, repository))
		r.Created = time.Now().UnixNano() / int64(time.Millisecond)
	}

	r.Namespace, r.Repository, r.JSON, r.Agent, r.Version = namespace, repository, json, agent, version

	r.Updated = time.Now().UnixNano() / int64(time.Millisecond)
	r.Checksumed, r.Uploaded, r.Cleared, r.Encrypted = false, false, false, false
	r.Size, r.Download = 0, 0

	if err := r.Save(); err != nil {
		return err
	}

	return nil
}
