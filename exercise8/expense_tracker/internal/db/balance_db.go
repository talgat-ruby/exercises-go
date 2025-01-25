package db

func (e *ExpencesDBSt) DBBalance(id int) (map[string]float64, error) {
	stmt := `
	SELECT amount 
	FROM budget
	WHERE id_user = $1;
	`
	balance := make(map[string]float64)
	row := e.NewDb.QueryRow(stmt, id)
	var sum float64
	err := row.Scan(
		&sum,
	)
	if err != nil {
		return nil, err
	}
	balance["balance"] = sum
	return balance, nil
}
