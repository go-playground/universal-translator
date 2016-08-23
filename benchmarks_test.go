package ut

import "testing"

func BenchmarkBasicTranslation(b *testing.B) {

	ut, _ := New("en", "en")
	loc := ut.FindTranslator("en")
	loc.Add("welcome", "Welcome to the site")
	loc.Add("welcome-user", "Welcome to the site {0}")
	loc.Add("welcome-user2", "Welcome to the site {0}, your location is {1}")

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
