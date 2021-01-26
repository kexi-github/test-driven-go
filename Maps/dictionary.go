package Maps

import "errors"

type Dictionary map[string]string

func (d Dictionary)Search(key string) (string,error){

	value,ok := d[key]
	if !ok {
		return "",errors.New("could not find the word you were looking for")
	}
	return value,nil
}

func (d Dictionary)Add(key,value string) error {

	_,err := d.Search(key)
	if err == nil{
		return errors.New("key is exists")
	} else{
		d[key] = value
		return nil
	}
}

func (d Dictionary)Update(key,value string) error {

		d[key] = value
		return nil
}

func (d Dictionary)Delete(key string) error {

	delete(d,key)
	return nil
}
