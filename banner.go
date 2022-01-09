package main

import "fmt"

const banner = `
   _____       __          
  / ___/__  __/ /___ ______
  \__ \/ / / / / __  / ___/
 ___/ / /_/ / / /_/ (__  ) 
/____/\__, /_/\__,_/____/  
     /____/  %s 
`

// showBanner is used to show the banner to the user
func showBanner() {
	fmt.Printf("%s\n", fmt.Sprintf(banner, Version))
	fmt.Print("\t sylas \n\n")
}
