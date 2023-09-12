go install github.com/gin-gonic/gin@latest
go mod init url-shortener
go get github.com/lib/pq
go get -u github.com/skip2/go-qrcode


const (
	pgHost     = "dpg-chcpv9e7avjcvo46f710-a.singapore-postgres.render.com"
	pgPort     = 5432
	pgUser     = "wxli3388"
	pgPassword = "wRiD7OgH7RgGUctB9YvWL1757uQQVAjV"
	pgDbname   = "demo_9dtw"
)

os.Getenv("pgHost")
os.Getenv("pgPort")
os.Getenv("pgUser")
os.Getenv("pgPassword")
os.Getenv("pgDbname")

os.Setenv("pgHost","dpg-chcpv9e7avjcvo46f710-a.singapore-postgres.render.com")
os.Setenv("pgPort","5432")
os.Setenv("pgUser","wxli3388")
os.Setenv("pgPassword","wRiD7OgH7RgGUctB9YvWL1757uQQVAjV")
os.Setenv("pgDbname","demo_9dtw")
os.Setenv("key","val")

pgHost="dpg-chcpv9e7avjcvo46f710-a.singapore-postgres.render.com" pgPort="5432" pgUser="wxli3388" pgPassword="wRiD7OgH7RgGUctB9YvWL1757uQQVAjV" pgDbname="demo_9dtw"