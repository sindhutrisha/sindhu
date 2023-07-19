package daos

import (
	"database/sql"
	"errors"
	"github.com/sindhutrisha/sindhu/teams/pkg/rest/server/daos/clients/sqls"
	"github.com/sindhutrisha/sindhu/teams/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
)

type TrishaDao struct {
	sqlClient *sqls.SQLiteClient
}

func migrateTrishas(r *sqls.SQLiteClient) error {
	query := `
	CREATE TABLE IF NOT EXISTS trishas(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
        
		Fields TEXT NOT NULL,
		Verified INTEGER NOT NULL,
        CONSTRAINT id_unique_key UNIQUE (Id)
	)
	`
	_, err1 := r.DB.Exec(query)
	return err1
}

func NewTrishaDao() (*TrishaDao, error) {
	sqlClient, err := sqls.InitSqliteDB()
	if err != nil {
		return nil, err
	}
	err = migrateTrishas(sqlClient)
	if err != nil {
		return nil, err
	}
	return &TrishaDao{
		sqlClient,
	}, nil
}

func (trishaDao *TrishaDao) CreateTrisha(m *models.Trisha) (*models.Trisha, error) {
	insertQuery := "INSERT INTO trishas(Fields, Verified)values(?, ?)"
	res, err := trishaDao.sqlClient.DB.Exec(insertQuery, m.Fields, m.Verified)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	m.Id = id

	log.Debugf("trisha created")
	return m, nil
}

func (trishaDao *TrishaDao) UpdateTrisha(id int64, m *models.Trisha) (*models.Trisha, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	trisha, err := trishaDao.GetTrisha(id)
	if err != nil {
		return nil, err
	}
	if trisha == nil {
		return nil, sql.ErrNoRows
	}

	updateQuery := "UPDATE trishas SET Fields = ?, Verified = ? WHERE Id = ?"
	res, err := trishaDao.sqlClient.DB.Exec(updateQuery, m.Fields, m.Verified, id)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sqls.ErrUpdateFailed
	}

	log.Debugf("trisha updated")
	return m, nil
}

func (trishaDao *TrishaDao) DeleteTrisha(id int64) error {
	deleteQuery := "DELETE FROM trishas WHERE Id = ?"
	res, err := trishaDao.sqlClient.DB.Exec(deleteQuery, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sqls.ErrDeleteFailed
	}

	log.Debugf("trisha deleted")
	return nil
}

func (trishaDao *TrishaDao) ListTrishas() ([]*models.Trisha, error) {
	selectQuery := "SELECT * FROM trishas"
	rows, err := trishaDao.sqlClient.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	var trishas []*models.Trisha
	for rows.Next() {
		m := models.Trisha{}
		if err = rows.Scan(&m.Id, &m.Fields, &m.Verified); err != nil {
			return nil, err
		}
		trishas = append(trishas, &m)
	}
	if trishas == nil {
		trishas = []*models.Trisha{}
	}

	log.Debugf("trisha listed")
	return trishas, nil
}

func (trishaDao *TrishaDao) GetTrisha(id int64) (*models.Trisha, error) {
	selectQuery := "SELECT * FROM trishas WHERE Id = ?"
	row := trishaDao.sqlClient.DB.QueryRow(selectQuery, id)
	m := models.Trisha{}
	if err := row.Scan(&m.Id, &m.Fields, &m.Verified); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}

	log.Debugf("trisha retrieved")
	return &m, nil
}
