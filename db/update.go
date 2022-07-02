package db

import "sorteo/model"

func UpdatePremio(preimio model.Premio) error {
	query := `UPDATE premio SET ip ='` + preimio.Ip + ` ' WHERE premio ='` + preimio.Premio + `'`
	update, err := ConDB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = update.Exec()
	if err != nil {
		return err
	}
	return nil
}
