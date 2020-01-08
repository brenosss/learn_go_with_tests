package integers

import "testing"

func TestAdder(t *testing.T) {

    t.Run("Sum expected 4", func(t *testing.T) {
        sum := Add(2, 2)
        expected := 4
        if sum != expected {
            t.Errorf("expected '%d' but got '%d'", expected, sum)
        }
    })
    t.Run("Sum expected 6", func(t *testing.T) {
        sum := Add(3, 3)
        expected := 6
        if sum != expected {
            t.Errorf("expected '%d' but got '%d'", expected, sum)
        }
    })
}
