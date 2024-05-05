package lilymetrics

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/efs"
	_ "github.com/lilypad-tech/lilypad/pkg/data"

	// "github.com/aws/aws-sdk-go"
	// "github.com/aws/aws-sdk-go/aws/session"
	// // "github.com/aws/aws-sdk-go-v2/aws"
	// // "github.com/aws/aws-sdk-go/aws/session"
	"net/http"
)

// Deal
func LogDeal(source string, deal, status string) {
	log.Print(deal)
	url := "http://" + os.Getenv("METRICS_HOST") + ":8000/metrics-dashboard/deal"
	json := fmt.Sprintf(`{"Type":"%s","Details":"%s"}`, deal, status)
	fmt.Println(json)
	data := []byte(json)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))
	return
}
func LogJob(source string, module_id string, job_id string, status string) {
	log.Print(module_id)

	url := "http://" + os.Getenv("METRICS_HOST") + ":8000/metrics-dashboard/job"
	json := fmt.Sprintf(`{"Type":"%s","Details":"%s"}`, module_id, status)
	fmt.Println(json)
	data := []byte(json)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))
	return
}

func LogMetric(module_id string, detail string) {
	log.Print(module_id)
	url := "http://" + os.Getenv("METRICS_HOST") + ":8000/metrics-dashboard/log"
	json := fmt.Sprintf(`{"Type":"%s","Details":"%s"}`, module_id, detail)
	fmt.Println(json)
	data := []byte(json)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))
	return
	// Create a new AWS session
	// sess, err := session.NewSession(&aws.Config{
	// 	Region: aws.String("your-region"),
	// })
	// if err != nil {
	// 	log.Fatal("Failed to create session: ", err)
	// }

	// Create an Amazon EFS service client
	// svc := efs.New(sess)

	// // Specify the file system ID of your EFS file system
	// fileSystemID := "your-efs-file-system-id"

	// // Specify the path of the log file on the EFS file system
	// logFilePath := "/path/to/logfile.log"

	// Log an event
	// logEvent := "This is a test log event."

	// Write the log event to the log file on EFS
	// err = writeToEFS(svc, fileSystemID, logFilePath, logEvent)
	// if err != nil {
	// 	log.Fatal("Failed to write log event to EFS: ", err)
	// }

	// log.Println("Log event has been successfully written to EFS.")
}

func writeToEFS(svc *efs.EFS, fileSystemID, filePath, content string) error {

	// Construct input parameters for the DescribeFileSystems API call
	input := &efs.DescribeFileSystemsInput{
		FileSystemId: aws.String(fileSystemID),
	}

	// Call DescribeFileSystems to get information about the specified EFS file system
	result, err := svc.DescribeFileSystems(input)
	if err != nil {
		return err
	}

	// Check if the file system exists
	if len(result.FileSystems) == 0 {
		return err
	}

	// Get the DNS name of the file system
	// result.FileSystems[0].fileSystemDNSName
	// fileSystemDNSName := result.FileSystems[0].DNSName
	var fileSystemDNSName = ""
	// Construct the path for the EFS endpoint
	efsEndpoint := "https://" + fileSystemDNSName

	// Create a new session for EFS with custom endpoint
	sessEFS, err := session.NewSession(&aws.Config{
		Endpoint: aws.String(efsEndpoint),
	})
	if err != nil {
		return err
	}

	// Create a new EFS service client with custom endpoint
	svcEFS := efs.New(sessEFS)

	// Specify the content to be written to the log file
	// logContent := []byte(content)

	// Construct input parameters for the CreateAccessPoint API call

	inputAccessPoint := &efs.CreateAccessPointInput{
		ClientToken:  aws.String("unique-client-token"), //todo: get token from dopler
		FileSystemId: aws.String(fileSystemID),
		PosixUser: &efs.PosixUser{
			Gid: aws.Int64(1001),
			Uid: aws.Int64(1001),
		},
	}

	// Call CreateAccessPoint to create an access point for the EFS file system
	accessPointOutput, err := svcEFS.CreateAccessPoint(inputAccessPoint)
	if err != nil {
		return err
	}

	// Get the ID of the access point
	accessPointID := accessPointOutput.AccessPointId

	// Construct input parameters for the PutObject API call
	// inputPutObject := &efs.PutObjectInput{
	// 	FileSystemId: aws.String(fileSystemID),
	// 	Path:         aws.String(filePath),
	// 	Body:         bytes.NewReader(logContent),
	// }
	// x := &efs.CreateTagsInput{
	// 	FileSystemId: aws.String(fileSystemID),
	// 	Path:         aws.String(filePath),
	// 	Body:         bytes.NewReader(logContent),
	// }

	// Call PutObject to write the content to the log file on EFS
	// _, err = svcEFS.put.PutObject(inputPutObject)
	// if err != nil {
	// 	return err
	// }

	// Clean up: delete the access point
	_, err = svcEFS.DeleteAccessPoint(&efs.DeleteAccessPointInput{
		AccessPointId: accessPointID,
	})
	if err != nil {
		return err
	}

	return nil
}
