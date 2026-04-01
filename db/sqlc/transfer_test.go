package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"venusbank.sqlc.dev/app/util"
)

func CreateRandomTransfer(t *testing.T) Transfer {
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	CreateRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer1 := CreateRandomTransfer(t)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)
	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)

	for i := 0; i < 10; i++ {
		_, err := testQueries.CreateTransfer(context.Background(), CreateTransferParams{
			FromAccountID: account1.ID,
			ToAccountID:   account2.ID,
			Amount:        util.RandomMoney(),
		})
		require.NoError(t, err)
	}

	transfers, err := testQueries.ListTransfers(context.Background(), ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Limit:         5,
		Offset:        5,
	})
	require.NoError(t, err)
	require.Len(t, transfers, 5)
	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}

func TestTransferTx(t *testing.T) {
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)

	n := 3
	amount := int64(50)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := testStore.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)
		require.NotZero(t, result.Transfer.ID)
		require.NotZero(t, result.FromEntry.ID)
		require.NotZero(t, result.ToEntry.ID)
	}

	updatedAccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	require.Equal(t, account1.Balance-int64(n)*amount, updatedAccount1.Balance)
	require.Equal(t, account2.Balance+int64(n)*amount, updatedAccount2.Balance)
}
