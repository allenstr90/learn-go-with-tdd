package pointererrors

import "testing"

func TestWallet(t *testing.T) {

	t.Run("Depopsit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(100))

		want := Bitcoin(100)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(5))
		want := Bitcoin(15)

		assertNoError(t, err)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		initBalance := Bitcoin(10)
		wallet := Wallet{balance: initBalance}
		err := wallet.Withdraw(Bitcoin(20))

		assertError(t, err, InsufficientFundsException)
		assertBalance(t, wallet, initBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s but want %s", got, want)
	}
}

func assertError(t testing.TB, err error, want error) {
	t.Helper()
	if err == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if err != want {
		t.Errorf("got %s but want %s", err, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("Got an error but didn't want one")
	}
}
