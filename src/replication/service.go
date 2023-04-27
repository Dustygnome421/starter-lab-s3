package replication

import (
	"sync"
	"time"
)

/*
service.go serves as an emulator for the Amazon Web Services Simple Storage Service (S3). Implement the below functions
to create a complete miniature implementation of MiniS3

TODO: implement all methods in this file
*/

var fileTableMutex sync.Mutex

/*
InitS3 initializes the node computers and unlocks the file table

TODO: you are encouraged to edit this method, if there is something you want to do at the start of each test.
 */
func InitS3(node int) {
	ResetNodes()
	InitializeNodes(node)
}

/*
RequestWriteFile puts the specified file to the specified bucket.

This function must consider multiple clients using S3 at the same time. If two clients want to write to the same file at
the same time, then the client that requested to write first gets to write first. Perhaps implement a kind of scheduler?
*/
func RequestWriteFile(bucketName string, fileName string, fileContents []byte) {
	// TODO: implement this method

	// lock file table
	fileTableMutex.Lock()
	// unlock file table upon function exit
	defer fileTableMutex.Unlock()

	// create bucket if it does not exist
	if !BucketExists(bucketName) {
		CreateBucket(bucketName)
	}

	// write file to each node
	numberNodes := getNumberNodes()
	version := time.Now()
	// write to quorum
	for i := 0; i < numberNodes; i++ {
		// write file to node i
		WriteNodeFile(i%getWriteQuorum(), bucketName, fileName, fileContents, version)
	}

}

/*
RequestReadFile gets the contents of a file from the specified bucket

RequestReadFile must retrieve the local file from each node and reach a quorum before it returns the correct file

Additionally, this function must consider multiple clients using S3 at the same time. If one client wants to write to a
file while another client wants to read the same file at the same time, then the client that requested first gets to do
its action first. Perhaps implement a kind of scheduler?
*/
func RequestReadFile(bucketName string, fileName string) []byte {
	// TODO: implement this method

	// lock file table
	fileTableMutex.Lock()
	// unlock file table upon function exit
	defer fileTableMutex.Unlock()

	// check if bucket exists
	if !BucketExists(bucketName) {
		return []byte("Error")
	}

	// read file from each node
	numberNodes := getWriteQuorum()
	
	// sync.WaitGroup is used to wait for the program to finish all goroutines
	responses := sync.WaitGroup{}
	responses.Add(numberNodes)

	// fileMap maps node index to file contents
	fileMap := make(map[int][]byte)
	// fileTime maps node index to file version
	fileTime := make(map[int]time.Time)

	// read from quorum
	for i := 0; i < numberNodes; i++ {
		go func(node int) {	
			defer responses.Done()
			// read file from node i
			fileContents, version := ReadNodeFile(node, bucketName, fileName)
			// add file contents to map
			if (fileContents != nil) {
				fileMap[node] = fileContents
				fileTime[node] = version
			}
		}(i)
	}
	// wait for all goroutines to finish
	responses.Wait()

	// check if all files are the same
	var allResponses [][]byte
	// add all responses to allResponses
	for _, v := range fileMap {
		allResponses = append(allResponses, v)
	}

	// check if all files are the same
	// if all files are the same, return the file
	return allResponses[0][:]

}
