package repository

import (
	"microservice-assignment/shared"
	"microservice-assignment/src/movie/v1/model"
	"time"
)

func (r *MovieRepoMySQL) SaveLog(req *model.LogsRequest) <-chan shared.CommonResult {

	output := make(chan shared.CommonResult)
	go func() {
		defer close(output)

		sq := `INSERT INTO logs
				(
					uri, response
				)
			VALUES
				(
					?, ?
				)`

		tx, err := r.writeDB.Begin()
		if err != nil {
			output <- shared.CommonResult{Error: err}
			return
		}

		stmt, err := tx.Prepare(sq)
		if err != nil {
			tx.Rollback()
			output <- shared.CommonResult{Error: err}
			return
		}
		defer stmt.Close()

		_, err = stmt.Exec(req.Uri, req.Response)
		if err != nil {
			tx.Rollback()
			output <- shared.CommonResult{Error: err}
			return
		}

		tx.Commit()

		req.CreatedString = req.Created.Format(time.RFC3339)

		output <- shared.CommonResult{Result: req}

	}()

	return output
}