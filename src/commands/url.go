package commands

import (
	"bufio"
	"cli/src/utils"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// urlCmd represents the url command
var urlCmd = &cobra.Command{
	Use:                        "url",
	Short:                      "A brief description of your command",
	SuggestionsMinimumDistance: 1,
	RunE: func(cmd *cobra.Command, args []string) error {
		input:
			log.Println("Write exit to terminate")
			reader := bufio.NewReader(os.Stdin)
			log.Println("Enter the URL: ")
			txt, _ := reader.ReadString('\n')
			if txt == "exit\n" {
				return nil
			}
			url := strings.TrimSpace(txt)
			response, err := utils.GetCall(url)
			if err != nil {
				fmt.Println("Error: check for the URL entered, try something like https://www.google.com")
				goto input
			}
		
			fmt.Println(response.Status)
			err = ValidateURLResponse(response)
			if err != nil {
				fmt.Println("Error: ", err)
				goto input
			}
			defer response.Body.Close()
			data, err := ioutil.ReadAll(response.Body)

			if err != nil {
				fmt.Println("Error: ", err)
				goto input
			}
			_, err = utils.ParseXml(string(data))
			if err != nil {
				fmt.Println("Error: ", err)
				goto input
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

/*
TODO:
	- understand the code
	- Add unit tests and benchmarking
	- Think of more condition which we could use to validate URL response !!
	- Add a single line comment to every function which explains what the function is doing
	- Add a readme file here which should contains what this small project is doing,
		what technologies we have used here for unit tests, benchmarks and static code analysers,
		step to setup this and how could we run this project on our machine.
	- In case of error program should not be terminated.
	- termination only happens when we press enter exit or ctrl + c
	- find different way to run the program.
*/
