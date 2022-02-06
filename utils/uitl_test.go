package utils

import (
    "database/sql"
    "reflect"
    "testing"
)

func TestCreateTableWithUser(t *testing.T) {
    tests := []struct {
        name string
    }{
        // TODO: Add test cases.

    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {

        })
    }
}

func TestInitMysql(t *testing.T) {
    tests := []struct {
        name string
    }{
        // TODO: Add test cases.
        {name: "test1"},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            InitMysql()
        })
    }
}

func TestModifyDB(t *testing.T) {
    type args struct {
        sql  string
        args []interface{}
    }
    tests := []struct {
        name    string
        args    args
        want    int64
        wantErr bool
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := ModifyDB(tt.args.sql, tt.args.args...)
            if (err != nil) != tt.wantErr {
                t.Errorf("ModifyDB() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("ModifyDB() got = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestQueryRowDB(t *testing.T) {
    type args struct {
        sql string
    }
    tests := []struct {
        name string
        args args
        want *sql.Row
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := QueryRowDB(tt.args.sql); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("QueryRowDB() = %v, want %v", got, tt.want)
            }
        })
    }
}
