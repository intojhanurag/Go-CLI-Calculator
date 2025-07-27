package calc

import (
	"log"
	"os"
)

var Logger *log.Logger

func InitLogger(){
	file,_:=os.OpenFile("app.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	Logger=log.New(file,"CALC_LOG: ",log.Ldate|log.Ltime|log.Lshortfile)


	//O_CREATE= CREATE THE FILE IF IT DOESN'T EXIST
	//OS_WRONLY= OPEN THE FILE IN WRITE-ONLY MODE 
	//OS_APPEND=APPEND TO THE FILE - DON'T OVERWRITE EXISTING CONTENT
	// 0666 - it is a file permission mode in octal (base-8)
	// 0- special bit
	// 6- Owner (User)
	// 6- Group
	// 6- Others(World)

	// 6 = 4 + 2 → Read + Write
	

}