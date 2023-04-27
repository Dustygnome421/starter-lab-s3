package replication

import (
	"errors"
	"fmt"
	"os"
	"time"
)

/*
TODO: implement the methods at the bottom of this file.
*/

/*
InitializeNodes creates folders for each "node" computer in the mini-S3 system.

File directory will be /src/replication/nodes/NODE_NUMBERS/BUCKET_NAME/FILENAME
*/
func InitializeNodes(nodes int) {
	setNumberNodes(nodes)
	for i := 0; i < nodes; i++ {
		nodeDir := fmt.Sprintf("nodes/%d", i)
		err := os.MkdirAll(nodeDir, os.ModePerm)
		checkError(err)
	}
}

/*
ResetNodes removes /nodes/ and all subdirectories and all files within those directories
 */
func ResetNodes() {
	err := os.RemoveAll("nodes")
	checkError(err)
}

/*
BucketExists determines if the specified bucket exists
*/
func BucketExists(bucketName string) bool {
	numberNodes := getNumberNodes()
	for i := 0; i < numberNodes; i++ {
		bucket := fmt.Sprintf("nodes/%d/%s", i, bucketName)
		if _, err := os.Stat(bucket); errors.Is(err, os.ErrNotExist) {
			// bucket does not exist at node i
			return false
		}
	}
	return true
}

/*
CreateBucket should create a fake S3 bucket in the form of a directory. A bucket must be created before a file can be written
to the bucket. The bucket must be created in each node.
*/
func CreateBucket(
	bucketName string,
) {
	// TODO: implement this method
	if !BucketExists(bucketName) {
		// create bucket in each node
		// nodes/i/bucketName
		numberNodes := getNumberNodes()
		for i := 0; i < numberNodes; i++ {
			// create bucket directory
			bucket := fmt.Sprintf("nodes/%d/%s", i, bucketName)
			// create bucket directory
			err := os.Mkdir(bucket, os.ModePerm)
			checkError(err)
		}
	}
}


/*
WriteNodeFile should write a byte array (contents) to the specified bucket to the specified node (nodeIndex) with the
specified file name.

Should write the file version to a paired file (perhaps something like fileName.version?). version should be the current
time.

Returns the number of bytes written
 */
func WriteNodeFile(
	nodeIndex int,
	bucketName string,
	fileName string,
	contents []byte,
	version time.Time,
) int {
	// TODO: implement this method.

	// create bucket if it doesn't exist
	bucket := fmt.Sprintf("nodes/%d/%s", nodeIndex, bucketName)
	
	// create file
	filepath := fmt.Sprintf("%s/%s", bucket, fileName)
	versionFilePath := fmt.Sprint(filepath, ".", version.Format("2006-01-02"))

	// create file
	file, err := os.Create(filepath)
	checkError(err)

	defer file.Close()

	// write contents to file
	bytesWritten, err := file.Write(contents)
	checkError(err)

	// write version to version file
	versionFile, err := os.Create(versionFilePath)
	checkError(err) // stuck here

	defer versionFile.Close()

	_, err = fmt.Fprintln(versionFile, version)
	checkError(err)

	// return number of bytes written
	return bytesWritten
}

/*
ReadNodeFile should read the specified file from the specified node from the specified bucket

Returns file contents, file version
 */
func ReadNodeFile(
	nodeIndex int,
	bucketName string,
	fileName string,
) ([]byte, time.Time) {
	// TODO: implement this method.

	// check if bucket exists
	if !BucketExists(bucketName) {
		return nil, time.Now()
	}

	// read file
	// nodes/i/bucketName/fileName
	bucket := fmt.Sprintf("nodes/%d/%s", nodeIndex, bucketName)
	filepath := fmt.Sprintf("%s/%s", bucket, fileName)
	
	// open file
	file, err := os.Open(filepath)
	if err != nil {
		return nil, time.Now()
	}	

	defer file.Close()

	// read file contents stats
	fileInfo, err := file.Stat()
	checkError(err)

	// read file contents
	content := make([]byte, fileInfo.Size())
	_, err = file.Read(content)
	checkError(err)

	// return file contents and version
	return content, fileInfo.ModTime().Round(time.Second)
}

