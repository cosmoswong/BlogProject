package models

import (
    "BlogProject/utils"
    "github.com/stretchr/testify/assert" // 这里引入了 testify
    "testing"
    "time"
)

func TestInsertUser(t *testing.T) {
    utils.InitMysql()

    type args struct {
        user User
    }
    tests := []struct {
        name    string
        args    args
        want    int64
        wantErr bool
    }{
        // TODO: Add test cases.
       {   name:"insert success",
           args: args{ User{1,"zhangsan","zhangsan",0,time.Now().Unix()}},
           want: 1,
           wantErr: true,
       },
    }
    for _, tt := range tests {
        got, err := InsertUser(tt.args.user)
        assert.Equal(t, tt.want, got)
        assert.Equal(t, tt.wantErr, err != nil)
    }
}
