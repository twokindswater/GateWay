package web

var (
	web *Web
)

//func TestMain(m *testing.M) {
//	ctx := context.Background()
//	conf := config.GetConfig()
//
//	web, err := Init(ctx, conf.Web.Port)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	serial, err := serializer.Init(conf.Serializer.Type)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	db, err := db.Init(conf.DB.Type, conf.DB.Address, serial)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	err = db.Clear(ctx)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	repo, err := Init(web, db)
//	repo.AddHandler(ctx)
//
//	code := m.Run()
//	os.Exit(code)
//}
//
//func TestRepository_SetAccountHandler(t *testing.T) {
//
//	account := &model.AccountInfo{
//		Id:        "id",
//		Name:      "Lee",
//		Image:     "image",
//		SSID:      "127.0.0.1",
//		BSSID:     "127.0.0.2",
//		Street:    "Seoul",
//		InitDate:  20210101,
//		Latitude:  10.0,
//		Longitude: 20.0,
//	}
//	r, w := io.Pipe()
//
//	err := binary.Write(w, binary.LittleEndian, &account)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	req := httptest.NewRequest(setAccountPath, "", r)
//	fmt.Printf("%v", req)
//
//}
