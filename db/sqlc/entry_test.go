package db

import (
	"context"
	_ "database/sql"
	"testing"
	_ "time"

	"github.com/stretchr/testify/require"
	"venusbank.sqlc.dev/app/util"
)

func CreateRandomEntry(t *testing.T) Entry {
	account1 := CreateRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: account1.ID,
		Amount:    util.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entry.AccountID, account1.ID)
	require.Equal(t, entry.Amount, arg.Amount)

	require.NotZero(t, entry.CreatedAt)
	require.NotZero(t, entry.ID)

	return entry
}

func TestCreateEntry(t *testing.T) {
	CreateRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry1 := CreateRandomEntry(t)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.CreatedAt, entry2.CreatedAt)

}

func TestListEntries(t *testing.T) {
	account1 := CreateRandomAccount(t)
	for i := 0; i < 10; i++ {
		arg := CreateEntryParams{
			AccountID: account1.ID,
			Amount:    util.RandomMoney(),
		}
		entry, err := testQueries.CreateEntry(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, entry)
	}
	arg := ListEntriesParams{
		AccountID: account1.ID,
		Limit:     5,
		Offset:    5,
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entries)
	require.Len(t, entries, 5)

	for _, ent := range entries {
		require.NotEmpty(t, ent)
	}
}
