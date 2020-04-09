package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Stringer) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("[FAILED]\ngot => %s\nwant => %s", got, want)
		}
	}

	assertError := func(t *testing.T, got, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("No error, but it should have been")
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	assertNoError := func(t *testing.T, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("got an error but didn't want one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {

		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		// fmt.Printf("address of balance in test is: %v \n", &wallet.balance)

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startinBalance := Bitcoin(20)
		wallet := Wallet{startinBalance}

		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startinBalance)
		assertError(t, err, ErrInsufficientFunds)

	})

}
