package commands

import (
	"bufio"
	"cli/src/utils"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)



func CallFunction(txt string) error{
	url := strings.TrimSpace(txt)
	response, err := utils.GetCall(url)
	if err != nil {
		log.Println("Error: check for the URL entered, try something like https://www.google.com")
		return err
	}

	err = ValidateURLResponse(response)

	if err != nil {
		log.Println("Error: ", err)
		return err
	}
	//fmt.Println(response.Status)

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Println("Error: ", err)
		return err
	}
	_, err = utils.ParseXml(string(data))
	if err != nil {
		log.Println("Error: ", err)
		return err
	}
	return nil
}

var CallFunctionCmd = &cobra.Command{
	Use: "CallFunction",
	Short: "CallFunction",
	SuggestionsMinimumDistance: 1,
	RunE: func(cmd *cobra.Command, args []string) error {
		CallFunction(args[0])
		return nil
	},
}

// urlCmd represents the url command
var urlCmd = &cobra.Command{
	Use:                        "url",
	Short:                      "A brief description of your command",
	SuggestionsMinimumDistance: 1,
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("Write exit to terminate")
		reader := bufio.NewReader(os.Stdin)
		log.Println("Enter the URL: ")
		txt, _ := reader.ReadString('\n')
		for ; txt!="exit\n" ; {
			//Checks the url again and again, until you type exit
			CallFunction(txt)
			reader = bufio.NewReader(os.Stdin)
			log.Println("Enter the URL: ")
			txt, _ = reader.ReadString('\n')
		}

		return nil
	},
}

func ValidateURLResponse(response *http.Response) error {
	if !utils.IsStatusOk(response) {
		return errors.New("the given website corresponds to url is down")
	}
	if !utils.IsReturnTypeJson(response) {
		return errors.New("no json response found, try with different url")
	}
	return nil
}

func init() {
	rootCmd.AddCommand(urlCmd)
}


