package ut

import "testing"

func BenchmarkBasicTranslation(b *testing.B) {

	ut, _ := New("en", "en")
	loc := ut.FindTranslator("en")

	translations := []struct {
		key      interface{}
		trans    string
		expected error
	}{
		{
			key:      "welcome",
			trans:    "Welcome to the site",
			expected: nil,
		},
		{
			key:      "welcome-user",
			trans:    "Welcome to the site {0}",
			expected: nil,
		},
		{
			key:      "welcome-user2",
			trans:    "Welcome to the site {0}, your location is {1}",
			expected: nil,
		},
	}

	for _, tt := range translations {
		if err := loc.Add(tt.key, tt.trans); err != nil {
			b.Fatalf("adding translation '%s' failed with key '%s'", tt.trans, tt.key)
		}
	}

	b.ResetTimer()

	b.Run("", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			loc.T("welcome")
		}
	})

	b.Run("Parallel", func(b *testing.B) {

		b.RunParallel(func(pb *testing.PB) {

			for pb.Next() {
				loc.T("welcome")
			}
		})
	})

	b.Run("With1Param", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			loc.T("welcome-user", "Joeybloggs")
		}
	})

	b.Run("ParallelWith1Param", func(b *testing.B) {

		b.RunParallel(func(pb *testing.PB) {

			for pb.Next() {
				loc.T("welcome-user", "Joeybloggs")
			}
		})
	})

	b.Run("With2Param", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			loc.T("welcome-user2", "Joeybloggs", "/dev/tty0")
		}
	})

	b.Run("ParallelWith2Param", func(b *testing.B) {

		b.RunParallel(func(pb *testing.PB) {

			for pb.Next() {
				loc.T("welcome-user2", "Joeybloggs", "/dev/tty0")
			}
		})
	})
}
