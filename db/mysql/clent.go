package mysql

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"

	"github.com/OscarTango1995/AuthServiceMicroservice/config"
	"github.com/OscarTango1995/AuthServiceMicroservice/db"
	"github.com/OscarTango1995/AuthServiceMicroservice/models"
)

const (
	sessionTableName = "session"
	userTableName    = "user"
)

func init() {
	db.Register("mysql", NewClient)
}

//The first implementation.
type client struct {
	db *sqlx.DB
}

func formatDSN() string {
	cfg := mysql.NewConfig()
	cfg.Net = "tcp"
	cfg.Addr = fmt.Sprintf("%s:%s", viper.GetString(config.DbHost), viper.GetString(config.DbPort))
	cfg.DBName = viper.GetString(config.DbName)
	cfg.ParseTime = true
	cfg.User = viper.GetString(config.DbUser)
	cfg.Passwd = viper.GetString(config.DbPass)
	return cfg.FormatDSN()
}

// NewClient initializes a mysql database connection
func NewClient(conf db.Option) (db.DataStore, error) {
	cli, err := sqlx.Connect("mysql", formatDSN())
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to db")
	}
	return &client{db: cli}, nil
}

func (c *client) CreateSession(session *models.Session) (string, error) {
	if session.ID != "" {
		return "", errors.New("id is not empty")
	}
	session.ID = uuid.NewV4().String()
	names := session.Names()
	if _, err := c.db.NamedExec(fmt.Sprintf(`INSERT INTO %s (%s) VALUES(%s)`, sessionTableName, strings.Join(names, ","), strings.Join(mkPlaceHolder(names, ":", func(name, prefix string) string {
		return prefix + name
	}), ",")), session); err != nil {

		return "", errors.Wrap(err, "failed to add user")
	}
	return "", nil
}

func (c *client) ListSession(id string) (*models.Session, error) {
	var sess models.Session
	if err := c.db.Get(&sess, fmt.Sprintf(`SELECT * FROM %s WHERE id = '%s'`, sessionTableName, id)); err != nil {
		return nil, err
	}
	return &sess, nil
}

func (c *client) UpdateSession(session *models.Session) error {
	names := session.Names()
	if _, err := c.db.NamedExec(fmt.Sprintf(`UPDATE %s SET %s WHERE id=:id`, sessionTableName, strings.Join(mkPlaceHolder(names[1:], "=:", func(name, prefix string) string {
		return name + prefix + name
	}), ",")), session); err != nil {
		return errors.Wrap(err, "failed to update user")
	}

	return nil

}

func (c *client) DeleteSession(id string) error {
	if _, err := c.db.Query(fmt.Sprintf(`DELETE FROM %s WHERE id= '%s'`, sessionTableName, id)); err != nil {
		return errors.Wrap(err, "failed to delete user")
	}

	return nil
}

func (c *client) CreateUser(user *models.User) (string, error) {
	if user.ID != "" {
		return "", errors.New("id is not empty")
	}
	user.ID = uuid.NewV4().String()
	names := user.Names()
	if _, err := c.db.NamedExec(fmt.Sprintf(`INSERT INTO %s (%s) VALUES(%s)`, userTableName, strings.Join(names, ","), strings.Join(mkPlaceHolder(names, ":", func(name, prefix string) string {
		return prefix + name
	}), ",")), user); err != nil {
		return "", errors.Wrap(err, "failed to create user")
	}
	return "", nil
}

func (c *client) ListUser(id string) (*models.User, error) {
	var tch models.User
	if err := c.db.Get(&tch, fmt.Sprintf(`SELECT * FROM %s WHERE id = '%s'`, userTableName, id)); err != nil {
		/*if err == sql.ErrNoRows {
			return nil, domainErr.NewAPIError(domainErr.NotFound, fmt.Sprintf("teacher: %s not found", id))
		}*/
		return nil, err
	}
	return &tch, nil
}

func (c client) UpdateUser(user *models.User) error {
	names := user.Names()
	if _, err := c.db.NamedExec(fmt.Sprintf(`UPDATE %s SET %s where id=:id`, userTableName, strings.Join(mkPlaceHolder(names[1:], "=:", func(name, prefix string) string {
		return name + prefix + name
	}), ",")), user); err != nil {
		return errors.Wrap(err, "failed to update employee")
	}
	return nil
}

func (c *client) DeleteUser(id string) error {
	if _, err := c.db.Query(fmt.Sprintf(`DELETE FROM %s WHERE id= '%s'`, userTableName, id)); err != nil {
		return errors.Wrap(err, "failed to delete user")
	}
	return nil
}

func mkPlaceHolder(names []string, prefix string, formatName func(name, prefix string) string) []string {
	ph := make([]string, len(names))
	for i, name := range names {
		ph[i] = formatName(name, prefix)
	}
	return ph
}
