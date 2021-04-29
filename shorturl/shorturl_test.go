package shorturl

// func TestGetShortUrl(t *testing.T) {
// 	in := httptest.NewRequest("GET", "/api/short-url?url=blah", nil)
// 	out := httptest.NewRecorder()
// 	logger := log.New(os.Stdout, "", log.LstdFlags)
// 	shorturlHandler := NewHandler(logger, bitly.NewClient(logger))
// 	shorturlHandler.shortUrlHandler(out, in)
// 	if out.Code != http.StatusOK {
// 		t.Logf("expected: %d\tgot: %d", http.StatusOK, out.Code)
// 	}
// 	if out.Body.String() != "blah" {
// 		t.Logf("expected: %d\tgot: %d", http.StatusOK, out.Code)
// 	}
// }
