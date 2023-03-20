package main

import "testing"

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		// wallet struct
		wallet := Wallet{}

		wallet.Deposite(10)

		got := wallet.Balance()
		// To make the int as a Bitcoin type
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(10)

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

		assertNoError(t, err)

	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(50)

		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	// We've introduced t.Fatal which will stop the test if it is called.
	// This is because we don't want to make any more assertions on the error returned if there isn't one around.
	// Without this the test would carry on to the next step and panic because of a nil pointer.
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
		// The fact that Go takes a copy of values is useful a lot of the time but
	}

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatal("got an error but didn't want one")
	}
}
