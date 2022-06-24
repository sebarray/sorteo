package db

import "sorteo/model"

func CreateIp(user model.User) (Models.Planta, error) {

	var planta Models.Planta

	//query := " INSERT INTO usuarios (NAME, EMAIL) VALUES (\"" + user.IdUser + "\", \"" + user.TIPO + "\")"
	query := idplant
	create, err := ConDB.Prepare(query)
	if err != nil {
		return planta, err
	}
	create.Exec()
	return planta, nil
}
