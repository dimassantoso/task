package repository

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	sqlMock "gopkg.in/DATA-DOG/go-sqlmock.v1"
	"microservice-assignment/src/movie/v1/model"
	"strconv"
	"testing"
)

type testCase struct {
	name, query                                  string
	rows                                         []string
	isResultError, wantReturnError, wantArgError bool
}

func TestVendorRepoPostgres_Save(t *testing.T) {
	logRequest := model.LogsRequest{
		Uri:      "http://www.omdbapi.com?apikey=faf7e5bb&s=captain&p=1&y=2020&type=series",
		Response: "{\"Response\":\"True\",\"totalResults\":\"4\",\"Search\":[{\"imdbID\":\"tt12580610\",\"Title\":\"The Epic Tales of Captain Underpants in Space\",\"Year\":\"2020–\",\"Type\":\"series\",\"Poster\":\"https://m.media-amazon.com/images/M/MV5BY2VmMTMzNTUtZWU0Zi00Yzg4LTk1NGUtYjY2ZDk5ZDZiNDg2XkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_SX300.jpg\"},{\"imdbID\":\"tt12604360\",\"Title\":\"Captain Fantasy\",\"Year\":\"2020–\",\"Type\":\"series\",\"Poster\":\"https://m.media-amazon.com/images/M/MV5BYjMzY2FmNzYtYWQ1MC00YjdiLTlkNzAtNDAwZjY2NzMwZWJiXkEyXkFqcGdeQXVyMjgxMzc0MDQ@._V1_SX300.jpg\"},{\"imdbID\":\"tt12715090\",\"Title\":\"Captain Manicorn\",\"Year\":\"2020–\",\"Type\":\"series\",\"Poster\":\"https://m.media-amazon.com/images/M/MV5BY2NjYTI2ZGEtODg0Yi00ZWQzLWJhYjUtZTA2Zjc2MWI1YmEyXkEyXkFqcGdeQXVyMTIxMDQzOTA1._V1_SX300.jpg\"},{\"imdbID\":\"tt12186092\",\"Title\":\"Captain Carol's Cosmic Conquest\",\"Year\":\"2020–\",\"Type\":\"series\",\"Poster\":\"https://m.media-amazon.com/images/M/MV5BZTEyOGRiZjEtYTFkZS00Nzc4LWEzYmQtYzI5MGVlZGQxY2NjXkEyXkFqcGdeQXVyNjgxMTIwNjY@._V1_SX300.jpg\"}]}",
	}

	tests := []testCase{
		{
			name:            "TestCase #%s: Success",
			query:           `INSERT .+`,
			isResultError:   false,
			wantReturnError: false,
		},
		{
			name:            "TestCase #%s: Failed",
			query:           `UPDATE .+`,
			isResultError:   true,
			wantReturnError: false,
		},
		{
			name:            "TestCase #%s: Failed",
			query:           `INSERT .+`,
			isResultError:   true,
			wantReturnError: true,
		},
		{
			name:            "TestCase #%s: Failed",
			query:           `INSERT .+`,
			isResultError:   true,
			wantReturnError: false,
			wantArgError:    true,
		},
	}

	for id, tt := range tests {
		testCaseID := strconv.Itoa(id + 1)
		testCaseName := fmt.Sprintf(tt.name, testCaseID)
		t.Run(testCaseName, func(t *testing.T) {
			db, mock, _ := sqlMock.New()
			defer db.Close()

			if !tt.wantReturnError {
				mock.ExpectBegin()
			}

			if !tt.wantArgError {
				mock.ExpectPrepare(tt.query).ExpectExec().WithArgs(logRequest.Uri, logRequest.Response).WillReturnResult(sqlMock.NewResult(1, 1))
				mock.ExpectCommit()
			} else {
				mock.ExpectPrepare(tt.query).ExpectExec().WithArgs(logRequest.Uri).WillReturnResult(sqlMock.NewResult(1, 1))
				mock.ExpectCommit()
			}

			r := NewMovieRepositoryMySQL(db, db)
			result := <-r.SaveLog(&logRequest)

			isError := false
			if result.Error != nil {
				isError = true
			}
			assert.Equal(t, isError, tt.isResultError)
		})
	}
}
