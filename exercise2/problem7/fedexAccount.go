package problem7

import "fmt"

type FedexAccount struct {
	name     string
	packages []string
}

func (fedexAcc *FedexAccount) getName() string {
	return fedexAcc.name
}

func (fedexAcc *FedexAccount) sendPackage(recipient string) {
	message := fmt.Sprintf("%s send package to %s", fedexAcc.getName(), recipient)
	fedexAcc.packages = append(fedexAcc.packages, message)
}
