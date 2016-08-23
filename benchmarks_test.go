package ut

import "testing"

func BenchmarkBasicTranslation(b *testing.B) {

	ut, _ := New("en", "en")
	loc := ut.FindTranslator("en")
	loc.Add("welcome", "Welcome to the site")

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
}
