package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/youngprinnce/simple-bank/util"
)


func createRandomTransfer(t *testing.T) Transfer {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		Amount: util.RandomMoney(),
		ToAccountID: account2.ID,
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)	
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T){
	transfer1 := createRandomTransfer(t)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfer(t *testing.T){
	transfer1 := createRandomTransfer(t)
	// transfer2 := createRandomTransfer(t)

	arg := ListTransfersParams{
		FromAccountID: transfer1.FromAccountID,
		ToAccountID: transfer1.ToAccountID,
		Limit: 0,
		Offset: 0,
	}

	_, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	// require.NotEmpty(t, transfers)

	// require.Contains(t, transfers, transfer1)
	// require.Contains(t, transfers, transfer2)
}