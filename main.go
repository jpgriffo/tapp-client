package main

import "fmt"
import "net/http"
import "io/ioutil"
import "encoding/xml"
import "encoding/json"
import "os"
import "log"
import "crypto/tls"
import "runtime"
import "bytes"


import "github.com/jpgriffo/tapp-client/firewall"
import "github.com/jpgriffo/tapp-client/firewall/data"
import "./script"


type TappConfig struct {
	XMLName   xml.Name `xml:"tapp"`
	ApiEndpoint string `xml:"server,attr"`
	LogFile string `xml:"log_file,attr"`
	LogLevel string `xml:"log_level,attr"`
	Certificate Cert `xml:"ssl"`
}

type Cert struct {
	Cert string `xml:"cert,attr"`
	Key string `xml:"key,attr"`
	Ca string `xml:"server_ca,attr"`
}

func openTappConfiguration(fileLocation string) (config TappConfig) {
    xmlFile, err := os.Open(fileLocation)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()
	b, _ := ioutil.ReadAll(xmlFile) 
	// var config TappConfig
	xml.Unmarshal(b, &config)
	return config
}

func createClient(config TappConfig) (client *http.Client) {
	/**
	 * Loads Clients Certificates and creates and 509KeyPair
	 */
	cert, err := tls.LoadX509KeyPair(config.Certificate.Cert, config.Certificate.Key)
    if err != nil {
            log.Fatalln(err)
    }

    /**
     * Creates a client with specific transport configurations
     */
 	transport := &http.Transport{
        TLSClientConfig: &tls.Config { Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true},
    }
    client = &http.Client{Transport: transport}	
    return client
}

func main() {

	configPath := "./tapp/client.xml"
	endPoint := "cloud/firewall_profile"



	var policy data.Policy

	// connection.Get(configPath, endPoint, policy)



	config := openTappConfiguration(configPath)
    client := createClient(config)



    response, err := client.Get(config.ApiEndpoint + endPoint)
    if err != nil {
        log.Fatalln(err)
    }
    defer response.Body.Close()

    body, err := ioutil.ReadAll(response.Body)
   	if err != nil {
		log.Fatalln(err)
	} 


	fmt.Printf("%s\n\n",body)


	
	json.Unmarshal(body, &policy)

	fmt.Println( runtime.GOOS)

	if len(policy.Rules) > 0 {
		firewall.Drop()
		firewall.Apply(policy)
	}

	response, err = client.Get(config.ApiEndpoint + "blueprint/script_characterizations?type=boot")
    if err != nil {
        log.Fatalln(err)
    }
    defer response.Body.Close()

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}


	fmt.Printf("%s\n\n",body)

	var script_characterizations []script.ScriptCharacterization

	json.Unmarshal(body, &script_characterizations)
	fmt.Printf("%v\n", script_characterizations)

	for _, script_characterization := range script_characterizations {
		var b []byte
		script_conclusion, err := script_characterization.Execute()
		if err != nil {
			fmt.Println("error:", err)
		}
		b, err = script_conclusion.ToJson()
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Printf("%s\n", b)

		post_body := "{\"script_conclusion\":" + string(b[:]) + "}"
		response, err = client.Post(config.ApiEndpoint + "blueprint/script_conclusions", "application/json", bytes.NewBufferString(post_body))
	    if err != nil {
	        log.Fatalln(err)
	    }
	    defer response.Body.Close()
	}


}