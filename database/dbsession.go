package database

/*
// ServerSession ...
var ServerSession *session.Session

func init() {

	encoder := session.Base64Encode
	decoder := session.Base64Decode

	var provider session.Provider
	var err error

	//portln := os.Getenv("PORT")
	//if len(portln) != 0 {
	//	cfg := postgre.NewConfigWith("127.0.0.1", 5432, "postgres", "1234", "walid", "session")
	//	provider, err = postgre.New(cfg)

	//} else if !strings.HasPrefix(":", portln) {
	cfg := postgre.NewConfigWith("ec2-34-235-240-133.compute-1.amazonaws.com", 5432, "qrkpjazlukiadp", "129bc7576d446c3e85369edbd5563edd18d9be2f44521e89dca5956bd5ee0df0", "dfingkdn5kgjg", "session")
	provider, err = postgre.New(cfg)
	fmt.Println(err)
	//}

	//provider, err = memory.New(memory.Config{})
	cfg1 := session.NewDefaultConfig()
	cfg1.EncodeFunc = encoder
	cfg1.DecodeFunc = decoder
	ServerSession = session.New(cfg1)

	if err = ServerSession.SetProvider(provider); err != nil {
		log.Fatal(err)
	}

}
*/
