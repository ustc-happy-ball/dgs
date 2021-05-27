package db

import (
	"context"
	"dgs/db/databaseGrpc"
	"dgs/model"
	"time"
)

var accountService databaseGrpc.AccountServiceClient

func getAccountService() (databaseGrpc.AccountServiceClient, error) {
	if accountService == nil {
		service, err := GetAccountService()
		if err != nil {
			return nil, err
		}
		accountService = service
		return service, nil
	}
	return accountService, nil
}

func AccountGetAccountInfoByPlayerId(playerId int32) (*model.Account, error) {
	service, err := getAccountService()
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	requestMsg := &databaseGrpc.AccountGetAccountInfoByPlayerIdRequest{PlayerId: playerId}
	res, err := service.AccountGetAccountByPlayerId(ctx, requestMsg)
	if err != nil {
		return nil, err
	}
	accountInfoPb := res.AccountInfo
	if accountInfoPb == nil {
		return nil, nil
	}
	accountInfo := &model.Account{
		LoginName:     accountInfoPb.Phone,
	}
	return accountInfo, nil
}