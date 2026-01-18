package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yuttana76/simbplebank/util"
)

func createRandomEntrie(t *testing.T) Entry {
	account := createRandomAccount(t)

	arg := CreateEntrieParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entrie, err := testQueries.CreateEntrie(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entrie)

	require.Equal(t, arg.AccountID, entrie.AccountID)
	require.Equal(t, arg.Amount, entrie.Amount)

	require.NotZero(t, entrie.ID)
	require.NotZero(t, entrie.CreatedAt)

	return entrie
}

func TestCreateEntrie(t *testing.T) {
	createRandomEntrie(t)
}

func TestGetEntrie(t *testing.T) {
	entrie1 := createRandomEntrie(t)
	entrie2, err := testQueries.GetEntrie(context.Background(), entrie1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entrie2)

	require.Equal(t, entrie1.ID, entrie2.ID)
	require.Equal(t, entrie1.AccountID, entrie2.AccountID)
	require.Equal(t, entrie1.Amount, entrie2.Amount)
	require.WithinDuration(t, entrie1.CreatedAt, entrie2.CreatedAt, 0)
}
func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntrie(t)
	}

	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entrie := range entries {
		require.NotEmpty(t, entrie)
	}
}
