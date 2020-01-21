package main

import "testing"


func TestWallet(t *testing.T) {

    assertError := func(t *testing.T, got error, want error) {
        t.Helper()
        if got == nil {
            t.Fatal("didn't get an error but wanted one")
        }

        if got.Error() != want.Error() {
            t.Errorf("got %q, want %q", got, want)
        }
    }

    t.Run("Deposit", func(t *testing.T) {
        wallet := Wallet{}
        wallet.Deposit(Bitcoin(10))

        got := wallet.Balance()
        want := Bitcoin(10)

        if got != want {
            t.Errorf("got %s want %s", got, want)
        }
    })

    t.Run("Withdraw", func(t *testing.T) {
        wallet := Wallet{balance: Bitcoin(20)}
        wallet.Withdraw(Bitcoin(10))

        got := wallet.Balance()

        want := Bitcoin(10)

        if got != want {
            t.Errorf("got %s want %s", got, want)
        }
    })

    t.Run("Withdraw insufficient founds", func(t *testing.T) {
        wallet := Wallet{balance: Bitcoin(20)}
        err := wallet.Withdraw(Bitcoin(100))

        assertError(t, err, ErrInsufficientFunds)
    })
}
