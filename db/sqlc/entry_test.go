package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/youngprinnce/simple-bank/util"
)

func createRandomEntry(t *testing.T) Entry {
	account1 := createRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: account1.ID,	
		Amount: util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T){
	entry1 := createRandomEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T){
	account1 := createRandomAccount(t)

	arg := ListEntriesParams{
		Limit: 0,
		Offset: 0,
		AccountID: account1.ID,
	}

	_, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)

	// require.NotEmpty(t, entries)	

	// for _, entry := range entries {
	// 	require.Equal(t, arg.AccountID, entry.AccountID)
	// }
}