package templatex

import (
	"fmt"
	"os"
	"path/filepath"
)

func StartTemplate() {

	os.MkdirAll("views/layout", os.ModePerm)

	indexFilePath := filepath.Join("views/layout", "index.html")
	if _, err := os.Stat(indexFilePath); err == nil {
		fmt.Println("index.html already exists, skipping creation.")
		return
	} else if os.IsNotExist(err) {
		err := os.WriteFile(indexFilePath, []byte(IndexHTML), 0644)
		if err != nil {
			fmt.Printf("Failed to create index.html: %v\n", err)
			return
		}
	} else {
		fmt.Printf("Error checking index.html: %v\n", err)
		return
	}

	navbarFilePath := filepath.Join("views/layout", "navbar.html")
	if _, err := os.Stat(navbarFilePath); err == nil {
		fmt.Println("navbar.html already exists, skipping creation.")
	} else if os.IsNotExist(err) {
		err = os.WriteFile(navbarFilePath, []byte(NavbarHTML), 0644)
		if err != nil {
			fmt.Printf("Failed to create navbar.html: %v\n", err)
			return
		}
	} else {
		fmt.Printf("Error checking navbar.html: %v\n", err)
		return
	}

	homeFilePath := filepath.Join("views", "home.html")
	if _, err := os.Stat(homeFilePath); os.IsNotExist(err) {
		err = os.WriteFile(homeFilePath, []byte(ContentHTML), 0644)
		if err != nil {
			fmt.Printf("Failed to create home.html: %v\n", err)
			return
		}
	}

	footerFilePath := filepath.Join("views/layout", "footer.html")
	if _, err := os.Stat(footerFilePath); os.IsNotExist(err) {
		err = os.WriteFile(footerFilePath, []byte(FooterHTML), 0644)
		if err != nil {
			fmt.Printf("Failed to create footer.html: %v\n", err)
			return
		}
	}

}
